package dbx

type Configs struct {
	*DatabaseConfig
	Replicas map[string]*DatabaseConfig `toml:"replicas"`
}

type DatabaseConfig struct {
	Driver          string `toml:"driver"`
	DSN             string `toml:"dsn"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	ConnMaxLifeTime int    `toml:"conn_max_lifetime"`
	ConnMaxIdleTime int    `toml:"conn_max_idletime"`
	Debug           bool   `toml:"debug"`
	Tracing         bool   `toml:"tracing"`
}
