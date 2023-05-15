/*
Package socket5 实现了一个基本的 SOCKS5 代理服务器，可以处理未加密的 TCP 流量。
该服务器支持多种认证方式，包括无认证、账号密码认证等，同时支持自定义认证接口。
该服务器还支持日志记录和流量监控，用户可以选择使用默认的实现或自定义实现。
*/
package socket5

import (
	"net"
	"strconv"
)

// Server 结构体定义了一个 socket5 服务器。
type Server struct {
	config        Config        // 服务器配置
	logger        Logger        // 日志记录器
	monitor       Monitor       // 流量监控器
	listener      net.Listener  // 监听
	authenticator Authenticator // 认证接口
}

// NewServer 函数创建一个 Server 的实例。
func NewServer(config Config, authenticator Authenticator, logger Logger, monitor Monitor) Server {
	if authenticator == nil {
		authenticator = &NoAuth{}
	}
	if logger == nil {
		logger = &DefaultLogger{}
	}
	if monitor == nil {
		monitor = &DefaultMonitor{}
	}
	return Server{config: config, authenticator: authenticator, logger: logger, monitor: monitor}
}

// ListenAndServe 函数启动 socket5 服务器。
func (s *Server) ListenAndServe() error {
	// 在指定的地址和端口上监听 TCP 连接
	listener, err := net.Listen("tcp", s.config.Address+":"+strconv.Itoa(s.config.Port))
	if err != nil {
		// 如果监听失败，则记录错误日志并返回错误信息
		s.logger.Error("Failed to listen on "+s.config.Address+":"+strconv.Itoa(s.config.Port), err)
		return err
	}

	// 记录服务器已经启动的信息
	s.logger.Infof("Listening on %s:%d", s.config.Address, s.config.Port)

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			// 如果连接失败，则记录错误日志并继续等待下一个连接
			s.logger.Error("Failed to accept connection", err)
			continue
		}

		// 创建一个新的连接对象，并在新的协程中处理客户端请求
		go func() {
			c := NewConnection(conn, s.config, s.authenticator, s.logger, s.monitor)
			c.Handle()
		}()
	}
}

// 关闭
func (s *Server) Close() {
	s.listener.Close()
}
