package memory_test

import (
	"testing"
	"time"

	"github.com/nbarbey/limiter"
	"github.com/nbarbey/limiter/drivers/store/common"
	"github.com/nbarbey/limiter/drivers/store/memory"
)

func TestMemoryStoreSequentialAccess(t *testing.T) {
	common.TestStoreSequentialAccess(t, memory.NewStoreWithOptions(limiter.StoreOptions{
		Prefix:          "limiter:memory:sequential",
		CleanUpInterval: 30 * time.Second,
	}))
}

func TestMemoryStoreConcurrentAccess(t *testing.T) {
	common.TestStoreConcurrentAccess(t, memory.NewStoreWithOptions(limiter.StoreOptions{
		Prefix:          "limiter:memory:concurrent",
		CleanUpInterval: 1 * time.Nanosecond,
	}))
}
