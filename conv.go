package convert

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// 16进制转byte
func Int2Bytes(i int) ([]byte, error) {
	hex, err := hex.DecodeString(fmt.Sprintf("%08X", i))
	return []byte(hex), err
}

// byte转10进制
func Bytes2Dec(b_hex []byte) (int, error) {
	str_hex := hex.EncodeToString(b_hex)
	base, err := strconv.ParseInt(str_hex, 16, 10)
	return int(base), err
}
