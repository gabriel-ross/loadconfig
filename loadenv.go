package loadconfig

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv loads environment variables defined in the file at path. LoadEnv
// expects variables in the format {key}={value} and can handle values that
// contain equal signs.
func LoadEnv(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split((bufio.ScanLines))

	for fileScanner.Scan() {
		key, val, ok := strings.Cut(fileScanner.Text(), "=")
		if ok {
			os.Setenv(key, val)
		}
	}
	return nil
}
