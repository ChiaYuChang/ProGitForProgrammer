package pkgs

import (
	"fmt"
	"time"
)

func Greeting(name string) {
	fmt.Printf("Hello, %s. How are you.\n", name)
}

func GetTime() string {
	return time.Now().Format(time.RFC822)
}
