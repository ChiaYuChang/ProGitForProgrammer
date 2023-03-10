package main

import (
	"fmt"

	"github.com/ChiaYuChang/ProGitForProgrammer/internal"
	"github.com/ChiaYuChang/ProGitForProgrammer/pkgs"
)

func main() {
	pkgs.Greeting("John")
	fmt.Println("Current time:", pkgs.GetTime())
	fmt.Println("Current user:", pkgs.GetUserName())
	internal.Hello()
}
