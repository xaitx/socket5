package socket5

import (
	"fmt"
	"io"
	"net"
)

// 常量
const (
	CONNECT uint8 = 0x01
	BIND    uint8 = 0x02
	UDP     uint8 = 0x03

	IPv4       uint8 = 0x01
	DOMAINNAME uint8 = 0x03
	IPv6       uint8 = 0x04
)

// 收到的请求头
type Request struct {
	// 版本号
	VER uint8
	// CMD
	CMD uint8
	// RSV
	RSV uint8
	// ATYPE
	ATYPE uint8
	// addr
	ADDR []byte
	// PORT
	PORT []byte
}

func NewRequest() *Request {
	return &Request{}
}

// 响应头
func (r *Request) receiveRequest(conn net.Conn) error {
	// 读取并解析请求头的前 4 个字节
	header := make([]byte, 4)
	if _, err := io.ReadFull(conn, header); err != nil {
		return err
	}
	r.VER = header[0]
	r.CMD = header[1]
	r.RSV = header[2]
	r.ATYPE = header[3]

	// 根据 ATYPE 字段的不同解析 ADDR 字段
	var addr []byte
	switch r.ATYPE {
	case 1: // IPv4 地址
		addr = make([]byte, 4)
	case 3: // 域名
		// 读取一个字节表示域名的长度，然后读取该长度的字节作为域名
		lenBytes := make([]byte, 1)
		if _, err := io.ReadFull(conn, lenBytes); err != nil {
			return err
		}
		addrLen := int(lenBytes[0])
		addr = make([]byte, addrLen)
	case 4: // IPv6 地址
		addr = make([]byte, 16)
	default:
		return fmt.Errorf("unsupported address type: %d", r.ATYPE)
	}
	if _, err := io.ReadFull(conn, addr); err != nil {
		return err
	}
	r.ADDR = addr

	// 读取 PORT 字段
	port := make([]byte, 2)
	if _, err := io.ReadFull(conn, port); err != nil {
		return err
	}
	r.PORT = port
	return nil
}

// 值转为byte
func (r *Request) Marshal() []byte {
	b := make([]byte, 4+len(r.ADDR)+2)
	b[0] = r.VER
	b[1] = r.CMD
	b[2] = r.RSV
	b[3] = r.ATYPE
	copy(b[4:], r.ADDR)
	copy(b[4+len(r.ADDR):], r.PORT)
	return b
}
