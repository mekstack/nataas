package xds

import (
	"context"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"os"
	"temp/internal/xds/config"
	"temp/pkg/logger"
)

func GetSnapshotCache(l *logger.Logger, nodeID string, cfg config.Envoy) cache.SnapshotCache {
	// Create a snapshotCache
	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, l)

	// Create the sp that we'll serve to Envoy
	sp := GenerateSnapshot(cfg)
	if err := sp.Consistent(); err != nil {
		l.Errorf("sp inconsistency: %+v\n%+v", sp, err)
		os.Exit(1)
	}
	l.Debugf("will serve sp %+v", sp)

	// Add the sp to the snapshotCache
	if err := snapshotCache.SetSnapshot(context.Background(), nodeID, sp); err != nil {
		l.Errorf("sp error %q for %+v", err, sp)
		os.Exit(1)
	}

	return snapshotCache
}
