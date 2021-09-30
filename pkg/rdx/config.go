package rdx

import "github.com/golang/protobuf/ptypes/duration"

type Config struct {
	Network      string
	Addr         string
	Password     string
	Db           int32
	DialTimeout  *duration.Duration
	ReadTimeout  *duration.Duration
	WriteTimeout *duration.Duration
}
