package hello

import "fmt"

func Hello(name string) string {
	res := fmt.Sprintf("Hello, %v!", name)

	return res
}
