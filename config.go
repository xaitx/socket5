package socket5

// Config 服务器配置
type Config struct {
	// Address 服务器监听地址
	Address string

	// Port 服务器监听端口
	Port       int
	EnableAuth bool

	// Auth 认证方法
	Auth Authenticator

	// Logger 日志方法
	Logger Logger

	// Monitor 流量监控方法
	Monitor Monitor
}

// NewConfig 创建一个默认的服务器配置
func NewConfig() *Config {
	return &Config{
		Address: "0.0.0.0",
		Port:    8888,
		Auth:    &NoAuth{},
		Logger:  &DefaultLogger{},
		Monitor: &DefaultMonitor{},
	}
}
