package toolsbox

import (
	"crypto/sha256"
)

// sha256 encrypt
func Sha256(content []byte) []byte {
	hasher := sha256.New()
	hasher.Reset()
	res := []byte{}
	lang, err := hasher.Write(content)
	if lang != len(content) || err != nil {
		return nil
	}
	res = hasher.Sum(res)
	return res
}

// md5 encrypt
func Md5(content []byte) []byte {
	hasher := sha256.New()
	hasher.Reset()
	res := []byte{}
	lang, err := hasher.Write(content)
	if lang != len(content) || err != nil {
		return nil
	}
	res = hasher.Sum(res)
	return res
}
