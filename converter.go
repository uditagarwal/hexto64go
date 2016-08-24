package hexto64

import (
	"encoding/base64"
	"encoding/hex"
)

func hexto64(hexcode string) string {
	hex2str, err := hex.DecodeString(hexcode)
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(hex2str))
	return string(encoded)
}
