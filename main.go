package main

import (
	"fmt"

	"github.com/ChiaYuChang/ProGitForProgrammer/internal"
	"github.com/ChiaYuChang/ProGitForProgrammer/pkgs"
)

func main() {
<<<<<<< HEAD
	pkgs.Greeting("John")
	fmt.Println("Current time:", pkgs.GetTime())
	fmt.Println("Current user:", pkgs.GetLinuxUser())
	internal.HelloFromLinuxBranch()
=======
	pkgs.Greeting("Jane")
	fmt.Println("Current User:", pkgs.GetUserName())
	internal.HelloFromWinBranch()
>>>>>>> win_branch
}
