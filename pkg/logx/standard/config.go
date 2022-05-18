package standard

// Config 日志配置
type Config struct {
	// 日志
	Out string `toml:"out"`
	// # 日志类别: debug, warn, info，error
	Level string `toml:"level"`
}
