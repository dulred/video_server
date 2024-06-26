package utils

import (
	"crypto/rand"
	"io"
	"fmt"
)

func NewUUID()(string,error){
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// UUID版本4的格式，设置版本号为0100，将第13位的值置为4
	uuid[6] = (uuid[6] & 0x0f) | 0x40 
	uuid[8] = (uuid[8] & 0x3f) | 0x80 

	// 将UUID字节转换为标准的字符串格式
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}