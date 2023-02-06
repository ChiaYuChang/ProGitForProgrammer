package internal

import (
	"fmt"
	"runtime"
)

func Hello() {
	switch runtime.GOOS {
	case "windows":
		HelloFromWinBranch()
	case "linux":
		HelloFromLinuxBranch()
	default:
		fmt.Println("Hello from unknown os.")
	}
}

func HelloFromWinBranch() {
	fmt.Println("Hello, I'm from win branch!!")
}

func HelloFromLinuxBranch() {
	fmt.Println("Hello, I'm from linux branch!!")
}
