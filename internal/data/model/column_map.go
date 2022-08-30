package model

type ColumnMap struct {
	values map[string]interface{}
}

func NewColumnMap() *ColumnMap {
	return &ColumnMap{
		values: make(map[string]interface{}, 0),
	}
}

func (c *ColumnMap) ToMap() map[string]interface{} {
	var mm = make(map[string]interface{})
	for k, v := range c.values {
		mm[k] = v
	}
	return mm
}

func (c *ColumnMap) GetInt64Column(key string) int64 {
	v, ok := c.values[key]
	if ok {
		return v.(int64)
	}
	return 0
}

func (c *ColumnMap) GetStringColumn(key string) string {
	v, ok := c.values[key]
	if ok {
		return v.(string)
	}
	return ""
}

func (c *ColumnMap) AddInt64Column(key string, val int64) {
	c.values[key] = val
}

func (c *ColumnMap) AddStringColumn(key string, val string) {
	c.values[key] = val
}

func (c *ColumnMap) AddBytesColumn(key string, val []byte) {
	c.values[key] = val
}

func (c *ColumnMap) AddFloat64Column(key string, val float64) {
	c.values[key] = val
}

func (c *ColumnMap) AddBoolColumn(key string, val bool) {
	c.values[key] = val
}

func (c *ColumnMap) AddColumn(key string, val interface{}) {
	switch v := val.(type) {
	case int:
		c.AddInt64Column(key, int64(v))
	case int8:
		c.AddInt64Column(key, int64(v))
	case int16:
		c.AddInt64Column(key, int64(v))
	case int32:
		c.AddInt64Column(key, int64(v))
	case int64:
		c.AddInt64Column(key, v)
	case uint:
		c.AddInt64Column(key, int64(v))
	case uint8:
		c.AddInt64Column(key, int64(v))
	case uint16:
		c.AddInt64Column(key, int64(v))
	case uint32:
		c.AddInt64Column(key, int64(v))
	case uint64:
		c.AddInt64Column(key, int64(v))
	case string:
		c.AddStringColumn(key, v)
	case float64:
		c.AddFloat64Column(key, v)
	case float32:
		c.AddFloat64Column(key, float64(v))
	case bool:
		c.AddBoolColumn(key, v)
	case []byte:
		c.AddBytesColumn(key, v)
	default:
		panic("invalid value type")
	}
}

func FromMap(mm map[string]interface{}) *ColumnMap {
	c := NewColumnMap()
	for key, val := range mm {
		c.AddColumn(key, val)
	}
	return c
}
