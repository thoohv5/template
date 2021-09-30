package hpx

type Config struct {
	// 地址：端口
	LocalAddr string `toml:"local_addr"`
	// 是否启用
	SwaggerEnabled bool `toml:"swagger_enabled"`
}
