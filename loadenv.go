package loadenv

import (
	"bufio"
	"os"
	"strings"
)

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
