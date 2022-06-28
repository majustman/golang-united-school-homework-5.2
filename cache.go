package cache

import "time"

type Cache struct {
	data    map[string]string
	expData map[string]time.Time
}

func NewCache() Cache {
	data := make(map[string]string)
	expData := make(map[string]time.Time)
	return Cache{data: data, expData: expData}
}

func (c *Cache) Get(key string) (string, bool) {
	t, ok := c.expData[key]
	if ok && t.Unix() <= time.Now().Unix() {
		return "", false
	}
	v, ok := c.data[key]
	if !ok {
		return "", false
	}
	return v, true
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
	delete(c.expData, key)
}

func (c *Cache) Keys() []string {
	tn := time.Now().Unix()
	var res []string
	for key := range c.data {
		t, ok := c.expData[key]
		if ok && t.Unix() <= tn {
			res = append(res, key)
		}
		if !ok {
			res = append(res, key)
		}
	}
	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.expData[key] = deadline
}
