package main

import (
	"context"
	"flag"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"temp/internal/xds"
	"temp/internal/xds/config"
	"temp/pkg/logger"
)

var (
	port       uint
	nodeID     string
	isDebug    bool
	configPath string
)

const (
	debug = "debug"
	warn  = "warn"
)

func init() {

	flag.BoolVar(&isDebug, "debug", false, "Enable xDS server debug logging")

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Path to config file
	flag.StringVar(&configPath, "cfgPath", "config/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	var err error

	var l *logger.Logger

	if isDebug {
		l, err = logger.New(debug)
	} else {
		l, err = logger.New(warn)
	}
	if err != nil {
		l.Fatalln(err)
	}

	cfg, err := config.New(configPath)
	if err != nil {
		l.Fatalln(err)
	}

	snapshotCache := xds.GetSnapshotCache(l, nodeID, *cfg)
	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: isDebug}
	srv := xds.NewServer(ctx, snapshotCache, cb)
	srv.Run(port, l)
}
