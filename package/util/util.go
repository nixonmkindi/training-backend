package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/beevik/ntp"
)

//TimestampSize const
const TimestampSize = 8

//ReverseBytes func
func ReverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

//AssertReaderEOF function
func AssertReaderEOF(reader *bytes.Reader) error {
	if reader.Len() != 0 {
		return fmt.Errorf("bad data length: %d unexpected bytes", reader.Len())
	}
	return nil
}

//DialTCP function
func DialTCP(addr *net.TCPAddr, timeout time.Duration) (*net.TCPConn, error) {
	// see also: go needs generics
	dialer := net.Dialer{Timeout: timeout}
	conn, err := dialer.Dial("tcp", addr.String())
	if err != nil {
		return nil, err
	}
	return conn.(*net.TCPConn), nil
}

//MustDecodeHex function for decoding string to hex
func MustDecodeHex(s string) []byte {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes
}

//MustDecodeHex32 function for decoding string to 32 bytes
func MustDecodeHex32(s string) [32]byte {
	var res [32]byte
	bytes := MustDecodeHex(s)
	copy(res[:], bytes)
	return res
}

//MustDecodeHex64 function for decoding hex to string
func MustDecodeHex64(s string) [64]byte {
	var res [64]byte
	bytes := MustDecodeHex(s)
	copy(res[:], bytes)
	return res
}

//Timestamp function generated timestamp
func Timestamp() uint64 {
	tstamp, _ := ntp.Time("0.beevik-ntp.pool.ntp.org") // TODO: Handle errors
	loc, _ := time.LoadLocation("Africa/Nairobi")      // TODO: handle errors
	tnano := tstamp.In(loc).UnixNano()
	//res := make([]byte, TimestampSize)
	//binary.BigEndian.PutUint64(res, uint64(tnano))
	return uint64(tnano)
}
