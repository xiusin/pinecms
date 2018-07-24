package cache

// Cacher 缓存 Interface
type Cacher interface {
	Add() bool
	Save() bool
	Get() interface{}
	GetString() string
	GetInt() int
	GetInt64() int64
}
