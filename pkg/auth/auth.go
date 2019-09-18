package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(originPasswd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(originPasswd), bcrypt.DefaultCost)
	return string(bytes), err
}

func Compare(hashedPasswd, originPasswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(originPasswd))
}
