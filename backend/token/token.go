package token

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"strings"
)

func GenerateSha512(timestamp int64, fields ...string) string {

	hasher := sha512.New()
	hasher.Write([]byte(strings.Join(fields, "") + strconv.FormatInt(timestamp, 10)))

	return hex.EncodeToString(hasher.Sum(nil))
}
