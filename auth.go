package socket5

import (
	"net"
)

// Authenticator 接口定义了认证器的方法。
type Authenticator interface {
	GetCode() uint8
	Authenticate(conn net.Conn) error
}

// NoAuth 实现了无认证的接口。
type NoAuth struct{}

// GetCode
func (a *NoAuth) GetCode() uint8 {
	return uint8(0)
}

// Authenticate 方法实现了无认证方式下的认证逻辑。
func (a *NoAuth) Authenticate(conn net.Conn) error {
	return nil
}

// // PasswordAuth 实现了基于用户名和密码的认证接口。
// type PasswordAuth struct {
// 	Username string // 用户名
// 	Password string // 密码
// }

// // Authenticate 方法实现了基于用户名和密码的认证逻辑。
// func (a *PasswordAuth) Authenticate(conn net.Conn) error {
// 	// 发送认证方法响应消息，表示支持账号密码认证
// 	if err := sendAuthenticationMethodResponse(conn, AuthenticationMethodUsernamePassword); err != nil {
// 		return err
// 	}

// 	// 接收客户端发送过来的账号密码认证信息
// 	username, password, err := receiveUsernamePasswordAuthRequest(conn)
// 	if err != nil {
// 		return err
// 	}

// 	// 校验用户名和密码是否正确
// 	if username != a.Username || password != a.Password {
// 		return errors.New("invalid username or password")
// 	}

// 	// 发送认证成功的响应消息
// 	if err := sendUsernamePasswordAuthResponse(conn); err != nil {
// 		return err
// 	}

// 	return nil
// }
