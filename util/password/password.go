package password

import (
	"dashboard/config"
	"encoding/base64"
	"log"
)

var aesKey = []byte(config.Key)

/* 对密码进行加密 */
func EncryptPassword(password string) (encryptPassword string, err error) {
	pass := []byte(password)
	xPass, err := AesEncrypt(pass, aesKey)
	if err != nil {
		log.Println("password encrypt error", err.Error())
		return encryptPassword, err
	}
	encryptPassword = base64.StdEncoding.EncodeToString(xPass)
	return encryptPassword, nil
}
