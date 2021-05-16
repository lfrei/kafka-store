package store

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var products *cache.Cache

func init() {
	products = cache.New(5*time.Minute, 10*time.Minute)
}

func AddProduct(id, product string) {
	products.Set(id, product, cache.NoExpiration)
}

func GetProduct(id string) string {
	product, found := products.Get(id)
	if found {
		return product.(string)
	}
	return ""
}
