package protocol

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Packet struct {
	len   int64
	isCmd bool
	md5   string
	data  []byte
}

// 生成新的 Packet
func New(cmd bool, b ...byte) *Packet {

	length := len(b)

	return nil

}

// 将 Packet 类型转换成 []byte
// len 字段长度为 8 字节
func (p *Packet) Bytes() []byte {

	return nil
}

// 从连接里读取数据，并生成 Packet
func Pack(conn net.Conn) (*Packet, error) {

	var length int64
	header := make([]byte, 8)

	err := ReadLength(conn, header)

	if err != nil {
		return nil, err
	}

	err = binary.Read(bytes.NewBuffer(header), binary.BigEndian, &length)

	if err != nil {
		return nil, err
	}

	data := make([]byte, length)

	err = ReadLength(conn, data)

	if err != nil {
		return nil, err
	}

	// todo 将 header 和 data 转成 Packet

	return nil, err
}

// 从连接中读取长度为 len([]byte) 的字节流
func ReadLength(conn net.Conn, data []byte) error {

	var (
		n, sum int
		err    error
	)

	for {
		n, err = conn.Read(data[n:])
		if err != nil {
			return err
		}

		sum += n

		if n >= len(data) {
			break
		}
	}
	return nil
}
