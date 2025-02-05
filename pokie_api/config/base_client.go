package config

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const apiUrl = "https://pokeapi.co/api/v2/"

var c *cache.Cache

func init() {
	c = cache.New(DefaultCacheSettings.MinExpire, DefaultCacheSettings.MaxExpire)
}

func call(endpoint string, obj interface{}) error {
	cached, found := c.Get(endpoint)
	if found && CacheSettings.UseCache {
		return json.Unmarshal(cached.([]byte), &obj)
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl+endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	SetCache(endpoint, body)

	return json.Unmarshal(body, &obj)
}