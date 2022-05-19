package cache

import "time"

type valueTime struct {
	value    string
	deadline time.Time
}

type Cache struct {
	data map[string]*valueTime
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]*valueTime),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	v := ""
	exist := false
	for _, k := range c.Keys() {
		if k == key {
			exist = true
			d := c.data[key]
			if d != nil {
				v = d.value
			}
			break
		}
	}
	return v, exist
}

func (c *Cache) Put(key, value string) {
	c.data[key] = &valueTime{
		value:    value,
		deadline: time.Unix(0, 0),
	}
}

func (c *Cache) Keys() []string {
	var s []string
	for k, vTime := range c.data {
		if vTime.deadline.UnixMicro() < time.Now().UnixMicro() && vTime.deadline.UnixMicro() > 0 {
			delete(c.data, k)
		} else {
			s = append(s, k)
		}
	}
	return s
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = &valueTime{
		value:    value,
		deadline: deadline,
	}
}
