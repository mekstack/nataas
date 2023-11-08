package xds

import (
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"log"
	"temp/internal/xds/config"
	"temp/internal/xds/models"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

const (
	routeName              = "local_route"
	httpv1                 = "http/1.1"
	httpv2                 = "h2"
	virtualHostName        = "local_service"
	managerStatPrefix      = "ingress_http"
	tlsTransportSocketName = "envoy.transport_sockets.tls"
	xdsClusterName         = "xds_cluster"
)

func createDownstreamTlsContextAny(tls models.TLS) (*anypb.Any, error) {

	upstreamTlsContext := &tlsv3.DownstreamTlsContext{
		CommonTlsContext: &tlsv3.CommonTlsContext{
			TlsCertificates: []*tlsv3.TlsCertificate{getTLSCertificate(tls.Cert, tls.Key)},
			AlpnProtocols:   []string{httpv1, httpv2},
		},
	}

	a, err := anypb.New(upstreamTlsContext)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func makeCluster(cl models.Cluster) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 cl.Name,
		ConnectTimeout:       durationpb.New(5 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_STRICT_DNS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(cl.Name, cl.Endpoints),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
	}
}

func makeEndpoint(clusterName string, eps []models.Endpoint) *endpoint.ClusterLoadAssignment {
	lbEndpoints := make([]*endpoint.LbEndpoint, 0)
	for _, ep := range eps {
		lbEndpoints = append(lbEndpoints, &endpoint.LbEndpoint{
			HostIdentifier: &endpoint.LbEndpoint_Endpoint{
				Endpoint: &endpoint.Endpoint{
					Address: &core.Address{
						Address: &core.Address_SocketAddress{
							SocketAddress: &core.SocketAddress{
								Protocol: core.SocketAddress_TCP,
								Address:  ep.Address,
								PortSpecifier: &core.SocketAddress_PortValue{
									PortValue: ep.Port,
								},
							},
						},
					},
				},
			},
		})
	}
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: lbEndpoints,
		}},
	}
}

func makeRoute(routes []models.Route) *route.RouteConfiguration {
	var rts []*route.Route
	for _, r := range routes {
		rts = append(rts, &route.Route{
			Match: &route.RouteMatch{
				PathSpecifier: &route.RouteMatch_Prefix{
					Prefix: r.Prefix,
				},
			},
			Action: &route.Route_Route{
				Route: &route.RouteAction{
					ClusterSpecifier: &route.RouteAction_Cluster{
						Cluster: r.ClusterName,
					},
				},
			},
		})
	}
	return &route.RouteConfiguration{
		Name: routeName,
		VirtualHosts: []*route.VirtualHost{{
			Name:    virtualHostName,
			Domains: []string{"*"},
			Routes:  rts,
		}},
	}
}

func makeHTTPListener(l models.Listener, tls models.TLS) (*listener.Listener, error) {
	a, err := createDownstreamTlsContextAny(tls)
	if err != nil {
		return nil, err
	}

	routerConfig, _ := anypb.New(&router.Router{})
	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: managerStatPrefix,
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    makeConfigSource(),
				RouteConfigName: routeName,
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name:       wellknown.Router,
			ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: routerConfig},
		}},
	}
	pbst, err := anypb.New(manager)
	if err != nil {
		return nil, err
	}

	return &listener.Listener{
		Name: l.Name,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  l.Address,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: l.Port,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
			TransportSocket: &core.TransportSocket{
				Name: tlsTransportSocketName,
				ConfigType: &core.TransportSocket_TypedConfig{
					TypedConfig: a,
				},
			},
		}},
	}, nil
}

func makeConfigSource() *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
		ApiConfigSource: &core.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   core.ApiConfigSource_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*core.GrpcService{{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: xdsClusterName},
				},
			}},
		},
	}
	return source
}

func getTLSCertificate(cert, key string) *tlsv3.TlsCertificate {
	return &tlsv3.TlsCertificate{
		CertificateChain: &core.DataSource{
			Specifier: &core.DataSource_Filename{Filename: cert},
		},
		PrivateKey: &core.DataSource{
			Specifier: &core.DataSource_Filename{Filename: key},
		},
	}
}

func GenerateSnapshot(cfg config.Envoy) *cache.Snapshot {
	clusters := make([]types.Resource, 0)
	listeners := make([]types.Resource, 0)
	for _, cl := range cfg.Clusters {
		clusters = append(clusters, makeCluster(cl))
	}

	for _, l := range cfg.Listeners {
		httpListener, err := makeHTTPListener(l, cfg.TLS)
		if err != nil {
			log.Fatalln(err)
		}
		listeners = append(listeners, httpListener)
	}

	snap, _ := cache.NewSnapshot("1",
		map[resource.Type][]types.Resource{
			resource.ClusterType:  clusters,
			resource.RouteType:    {makeRoute(cfg.Routes)},
			resource.ListenerType: listeners,
		},
	)
	return snap
}
