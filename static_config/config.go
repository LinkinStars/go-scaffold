package static_config

// StaticConfig config interface
type StaticConfig interface {
	// LoadAndSet load config from anywhere and set to conf struct
	LoadAndSet(conf interface{}) error
}
