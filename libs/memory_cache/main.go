package memory_cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var instance *cache.Cache

func GetInstance() *cache.Cache {
	if instance == nil {
		fmt.Println("Generating new instance for memory cache")
		instance = cache.New(5*time.Minute, 10*time.Minute)
	} else {
		fmt.Println("Memory Cache already instanced")
	}
	return instance
}
