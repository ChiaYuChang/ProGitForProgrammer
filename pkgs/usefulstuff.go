package pkgs

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Greeting(name string) {
	fmt.Printf("Hello, %s. How are you.\n", name)
}

func GetTime() string {
	return time.Now().Format(time.RFC822)
}

func GetUserName() string {
	switch runtime.GOOS {
	case "windows":
		return GetWinUserName()
	case "linux":
		return GetLinuxUserName()
	default:
		return "unknown os"
	}
}

func GetWinUserName() string {
	return os.Getenv("USERNAME")
}

func GetLinuxUserName() string {
	return os.Getenv("USER")
}
