package password

import "golang.org/x/crypto/bcrypt"

func MakePassword(password string) (hashedPassword string, err error) {
	r, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
