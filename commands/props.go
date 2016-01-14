package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readProps(path string) (props map[string]string, err error) {
	// read file
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	props = map[string]string{}
	for _, line := range lines {
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			// Invalid format; skip
			continue
		}
		props[kv[0]] = kv[1]
	}

	return

}

func writeProps(path string, props map[string]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for key, value := range props {
		line := fmt.Sprintf("%s=%s", strings.ToUpper(key), value)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
