package standard

// Log 数据库日志配置
type Log struct {
	Mode int `toml:"mode"` // 数据库日志：0-无日志, 1-写日志, 2-读写日志
	Cat  int `toml:"cat"`  // 日志类别: 1-无，2-Error, 3-Warn, 4-Info
}

// Config 数据库配置
type Config struct {
	// MySQL/PostgresSQL/SQLite
	Driver          string `toml:"driver"`             // 数据库驱动
	Dsn             string `toml:"dsn"`                // 数据库源
	ConnMaxLifeTime int    `toml:"conn_max_life_time"` // 数据库最大连接时长
	MaxIdleConns    int    `toml:"max_idle_conns"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	Slave           []struct {
		Dsn string `toml:"dsn"`
	} `toml:"slave"` // 数据库备用源
	Log *Log `toml:"log"` // 数据库日志
}

func (c *Config) GetLog() *Log {
	if c.Log == nil {
		c.Log = new(Log)
	}
	return c.Log
}
