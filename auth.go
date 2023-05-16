package socket5

import (
	"errors"
	"io"
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

type User struct {
}

func (u *User) authenticate(username []byte, password []byte) error {
	return nil
}

// PasswordAuth 实现了基于用户名和密码的认证接口。
type PasswordAuth struct {
	Username string
	Password string
	// 账号密码认证
	User *User
}

// GetCode
func (a *PasswordAuth) GetCode() uint8 {
	return uint8(2)
}

// Authenticate 方法实现了基于用户名和密码的认证逻辑。
func (a *PasswordAuth) Authenticate(conn net.Conn) error {
	// 验证版本 （不是socket5版本是认证协议的版本）
	ver := make([]byte, 1)
	io.ReadFull(conn, ver)
	if ver[0] != 1 {
		return errors.New("error Version")
	}

	// 读取账号长度
	userLen := make([]byte, 1)
	io.ReadFull(conn, userLen)

	// 读取账户
	user := make([]byte, int(userLen[0]))
	io.ReadFull(conn, user)

	// 读取密码长度
	passLen := make([]byte, 1)
	io.ReadFull(conn, passLen)

	// 读取密码
	pass := make([]byte, int(passLen[0]))
	io.ReadFull(conn, pass)

	// 校验用户名和密码是否正确
	if a.User == nil {
		if string(user) != a.Username || string(pass) != a.Password {
			return errors.New("invalid username or password")
		}
	} else {
		if err := a.User.authenticate(user, pass); err != nil {
			return err
		}

	}

	// 发送认证成功的响应消息
	conn.Write([]byte{socks5Version, 0x00})

	return nil
}
