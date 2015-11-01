package bench5

import "io/ioutil"

var (
	digits = []byte("0123456789abcdef")
	data   = []byte("DATA data+0x00000000(SB)/1, $0x00\n")
)

func formatByte(buf []byte, n byte) {
	buf[0] = digits[(n>>4)&0x0f]
	buf[1] = digits[n&0xf]
}

func formatInt(buf []byte, n uint32) {
	buf[0] = digits[(n>>28)&0x0f]
	buf[1] = digits[(n>>24)&0x0f]
	buf[2] = digits[(n>>20)&0x0f]
	buf[3] = digits[(n>>16)&0x0f]
	buf[4] = digits[(n>>12)&0x0f]
	buf[5] = digits[(n>>8)&0x0f]
	buf[6] = digits[(n>>4)&0x0f]
	buf[7] = digits[n&0x0f]
}

func formatIntHigh(buf []byte, n uint32) {
	buf[0] = digits[(n>>28)&0x0f]
	buf[1] = digits[(n>>24)&0x0f]
	buf[2] = digits[(n>>20)&0x0f]
	buf[3] = digits[(n>>16)&0x0f]
}

func formatIntLow(buf []byte, n uint32) {
	buf[4] = digits[(n>>12)&0x0f]
	buf[5] = digits[(n>>8)&0x0f]
	buf[6] = digits[(n>>4)&0x0f]
	buf[7] = digits[n&0x0f]
}

func formatIntNew(buf []byte, n uint32) {
	formatIntHigh(buf, n)
	formatIntLow(buf, n)
}

func Format() {
	const length = 512 * 1024 * 1024

	for i := 0; i < length; i++ {
		formatInt(data[12:], uint32(i))
		formatByte(data[31:], byte(i))

		ioutil.Discard.Write(data)
	}
}

func FormatNew() {
	const length = 512 * 1024 * 1024

	for i := 0; i < length; i++ {
		formatIntNew(data[12:], uint32(i))
		formatByte(data[31:], byte(i))

		ioutil.Discard.Write(data)
	}
}
