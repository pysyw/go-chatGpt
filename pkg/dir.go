package pkg

import (
	"os"
)

func MkdirIfNotExists(name string) (string, error) {
	f, err := os.Stat(name)
	if err != nil {
		err = os.Mkdir(name, 0777)
		if err != nil {
			return "", err
		}
		return name, nil
	}
	return f.Name(), nil
}
