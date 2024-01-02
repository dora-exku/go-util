package password

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestMakePassword(t *testing.T) {
	password := "password"
	hashedPassword, err := MakePassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// 验证超长密码 大于 72 位
	password = "passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"
	hashedPassword, err = MakePassword(password)
	require.EqualError(t, err, bcrypt.ErrPasswordTooLong.Error())
	require.Zero(t, len(hashedPassword))
}

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := MakePassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	err = CheckPassword(hashedPassword, "wrong password")
	require.Error(t, err)
}
