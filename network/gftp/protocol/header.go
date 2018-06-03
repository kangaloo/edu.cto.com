package protocol

type Header struct {
	length   int64
	checkSum string
	isCmd    bool
}

func NewHeader() *Header {

	return nil
}

// 将 Header 类型转换成 []byte 类型
func (header Header) Bytes() []byte {

	return nil
}
