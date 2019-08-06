package common

type IZookeeper interface {
	Read(path string) ([]byte, error)
	Write(path string, data []byte) error
	Create(path string, data []byte) error
	Delete(path string) error
}

type ICounter interface {
	MustInit() error
	GetOrderNumber() (int, error)
}

type IInternalCacheEngine interface {
	Set(key, value string) error
	Get(key string) (CacheValueType, error)
}
