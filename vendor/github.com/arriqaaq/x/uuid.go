package x

import (
	"github.com/nu7hatch/gouuid"
)

func GetUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), err
}
