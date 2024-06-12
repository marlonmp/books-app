package valobj

import "golang.org/x/crypto/bcrypt"

type Password struct {
	hash string
}

func PasswordFromHash(hash string) Password {
	return Password{hash}
}

func NewPassword(pwd string) (Password, error) {
	pwdB := []byte(pwd)

	hashB, err := bcrypt.GenerateFromPassword(pwdB, bcrypt.DefaultCost)

	if err != nil {
		return Password{}, err
	}

	hash := string(hashB)

	return Password{hash}, nil
}

func (p Password) IsEqual(pwd string) bool {
	hashB, pwdB := []byte(p.hash), []byte(pwd)

	err := bcrypt.CompareHashAndPassword(hashB, pwdB)

	return err == nil
}
