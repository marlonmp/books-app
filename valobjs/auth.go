package valobjs

import (
	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash  string
	isNil bool
}

func PasswordFromHash(hash string) Password {
	return Password{hash, false}
}

func NewPassword(pwd string) (Password, error) {
	pwdB := []byte(pwd)

	hashB, err := bcrypt.GenerateFromPassword(pwdB, bcrypt.DefaultCost)

	if err != nil {
		return Password{}, err
	}

	hash := string(hashB)

	return Password{hash, false}, nil
}

func (p Password) IsEqual(pwd string) bool {
	hashB, pwdB := []byte(p.hash), []byte(pwd)

	err := bcrypt.CompareHashAndPassword(hashB, pwdB)

	return err == nil
}

func (p *Password) Scan(src any) error {
	if src == nil {
		*p = Password{"", true}
		return nil
	}

	var hashedPassword []byte

	switch val := src.(type) {
	case string:
		hashedPassword = []byte(val)
	case []byte:
		hashedPassword = val
	}

	_, err := bcrypt.Cost(hashedPassword)

	if err != nil {
		return err
	}

	*p = Password{string([]byte(hashedPassword)), false}

	return nil
}

func (p Password) Value() (driver.Value, error) {
	if p.isNil {
		return nil, nil
	}

	return p.hash, nil
}
