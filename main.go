package main

import (
	"fmt"

	"github.com/ChiaYuChang/ProGitForProgrammer/internal"
	"github.com/ChiaYuChang/ProGitForProgrammer/pkgs"
)

func main() {
	pkgs.Greeting("Jane")
	fmt.Println("Current User:", pkgs.GetUserName())
	internal.HelloFromWinBranch()
}
