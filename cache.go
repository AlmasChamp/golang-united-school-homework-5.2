package cache

import "time"

type Cache struct {
	category map[string]string
	//expire
}

func NewCache() Cache {
	return Cache{
		category: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	if c.category[key] != "" {
		return c.category[key], true
	}
	return c.category[key], false
}

func (c *Cache) Put(key, value string) {
	//if c.category == nil {
	//	c.category = make(map[string]string)
	//}
	c.category[key] = value

}

func (c *Cache) Keys() []string {
	outArr := []string{}
	for k, _ := range c.category {
		outArr = append(outArr, k)
	}
	return outArr
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {

}
