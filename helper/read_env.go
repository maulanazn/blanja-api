package helper

import "os"

func ReadEnv(source string) string {
	file, _ := os.ReadFile(source)

	return string(file)
}
