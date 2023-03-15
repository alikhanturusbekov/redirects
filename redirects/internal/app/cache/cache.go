package cache

import (
	"log"
)

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type CacheRedirects struct {
	items []CacheItem
}

type CacheItem struct {
	key   string
	value string
}

var Cr CacheRedirects

func (cr *CacheRedirects) Add(key, value string) {
	if cr.Len() > 1000 {
		log.Println("Cache is fully filled")
	}

	ci := CacheItem{key: key, value: value}
	cr.items = append(cr.items, ci)
}

func (cr *CacheRedirects) Get(key string) (value string, ok bool) {
	for _, v := range cr.items {
		if v.key == key {
			return v.value, true
		}
	}

	return "", false
}

func (cr *CacheRedirects) Len() int {
	return len(cr.items)
}

func ConnectCache() {
	Cr = CacheRedirects{}
}
