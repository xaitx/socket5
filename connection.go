package socket5

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	socks5Version = uint8(5)
)

// Connection 结构体定义了一个 SOCKS5 代理服务器连接。
type Connection struct {
	conn          net.Conn      // 原始连接
	config        Config        // 服务器配置
	authenticator Authenticator // 认证接口
	logger        Logger        // 日志记录器
	monitor       Monitor       // 流量监控器
}

// NewConnection 函数创建一个 Connection 的实例。
func NewConnection(conn net.Conn, config Config, authenticator Authenticator, logger Logger, monitor Monitor) *Connection {
	return &Connection{conn: conn, config: config, authenticator: authenticator, logger: logger, monitor: monitor}
}

// Handle 方法处理 SOCKS5 代理服务器连接。
func (c *Connection) Handle() {
	// 进行认证
	if err := c.authenticate(); err != nil {
		c.logger.Error("Failed to authenticate connection", err)
		return
	}

	// 接收客户端请求
	request, err := c.receiveRequest()
	if err != nil {
		c.logger.Error("Failed to receive request", err)
		return
	}

	// 根据请求类型进行处理
	switch request.CMD {
	case CONNECT:
		c.handleConnect(*request)
	default:
		c.logger.Error("Unsupported command")
		return
	}
}

// authenticate 方法进行认证。
func (c *Connection) authenticate() error {

	// 接收客户端发送过来的认证方法
	buf := make([]byte, 2)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}

	// 解析客户端发送过来的协议报文
	if buf[0] != socks5Version {
		return errors.New("invalid protocol version")
	}

	// 获取客户端支持的认证方法列表
	authMethods := make([]byte, buf[1])
	if _, err := io.ReadFull(c.conn, authMethods); err != nil {
		return err
	}

	// 检查客户端发送过来的认证方法是否被支持
	if !bytes.Contains(authMethods, []byte{c.authenticator.GetCode()}) {
		c.conn.Write([]byte{socks5Version, 0xff})
		return errors.New("no acceptable authentication methods")
	}

	// 发送认证方法响应消息，表示支持客户端发送过来的认证方法
	if _, err := c.conn.Write([]byte{socks5Version, c.authenticator.GetCode()}); err != nil {
		return err
	}

	// 使用认证接口进行认证
	if err := c.authenticator.Authenticate(c.conn); err != nil {
		return err
	}

	return nil
}

// receiveRequest 方法接收客户端请求。
func (c *Connection) receiveRequest() (*Request, error) {
	// 创建
	r := NewRequest()
	// 接收客户端请求
	if err := r.receiveRequest(c.conn); err != nil {
		return nil, err
	}

	// 记录客户端请求的信息
	c.logger.Infof("Received request: %v", r)

	return r, nil
}

// handleConnect 处理 CONNECT 类型的请求
func (c *Connection) handleConnect(r Request) {

	// 转为字符串
	addr := string(r.ADDR)
	port := int(binary.BigEndian.Uint16(r.PORT))

	// 根据请求的地址和端口，建立与目标服务器的连接
	dstConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		c.config.Logger.Error("failed to connect to target: ", err)
		// 发送响应消息 由于请求头和响应头结构一样，直接重复使用
		r.CMD = uint8(4)
		c.conn.Write(r.Marshal())
		return
	}

	// 发送 CONNECT 响应消息 (为了省事，值为空)
	r.CMD = uint8(0)
	r.PORT = make([]byte, len(r.PORT))
	r.ADDR = make([]byte, len(r.ADDR))
	if _, err := c.conn.Write(r.Marshal()); err != nil {
		c.config.Logger.Error("failed to send CONNECT response: ", err)
		return
	}

	// 在客户端与目标服务器之间复制数据 过滤还未实现
	ich := make(chan error, 2)
	go porxy(dstConn, c.conn, ich)
	go porxy(c.conn, dstConn, ich)

	for i := 0; i < 2; i++ {
		<-ich
	}
	c.config.Logger.Info("close->", addr, port)

}
func porxy(dst net.Conn, src net.Conn, ich chan error) {
	defer func() {
		src.Close()
		dst.Close()
	}()
	_, err := io.Copy(dst, src)
	if err != nil {
		ich <- err
	} else {
		ich <- nil
	}

}
