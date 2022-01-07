package path

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetRoot() (string, error) {
	bin, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(bin)
	if err != nil {
		return "", err
	}

	return filepath.Dir(filepath.Dir(abs)), nil
}
