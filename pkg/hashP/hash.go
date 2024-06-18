package hashP

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Md5(s any) string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	h := md5.New()
	h.Write(jsonBytes)
	return hex.EncodeToString(h.Sum(nil))
}
