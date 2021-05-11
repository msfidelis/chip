package memory_cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var instance *cache.Cache

func GetInstance() *cache.Cache {
	// Memory Cache Singleton - Used to store key / values into container memory
	if instance == nil {
		instance = cache.New(5*time.Minute, 10*time.Minute)
	}
	return instance
}
