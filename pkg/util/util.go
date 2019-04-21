package util

import (
	"math/rand"
)

var randStrDict = make(map[string]string)

func init() {
	randStrDict["normal"] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randStrDict["strong"] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&’()*+,-./:;<=>?@[]^_`{|}~"
}

// RandStr generate a random string
func RandStr(_len int, _type string) string {
	if _, ok := randStrDict[_type]; ok == false || _type == "" {
		_type = "normal"
	}

	if _len < 3 {
		_len = 3
	}

	value, _ := randStrDict[_type]

	length := len(value)

	randByte := []byte{}

	for i := 0; i < _len; i++ {
		index := rand.Intn(length)
		randByte = append(randByte, value[index])
	}

	return string(randByte)
}
