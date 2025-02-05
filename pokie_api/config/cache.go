package config

import (
	"pokie_api/models"
	"time"
)

var CacheSettings = models.Settings{
	CustomExpire: 0,
	UseCache: true,
}

var DefaultCacheSettings = models.DefaultSettings{
	MaxExpire: 10 * time.Minute,
	MinExpire: 5 * time.Minute,
}

func SetCache(endpoint string, body []byte) {
	if CacheSettings.CustomExpire != 0 {
		c.Set(endpoint, body, CacheSettings.CustomExpire * time.Minute)
	} else {
		c.SetDefault(endpoint, body)
	}
}

func ClearCache() {
	c.Flush()
}