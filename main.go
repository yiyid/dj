package main

import (
	"runtime"
)

func main() {

	switch runtime.GOOS {
	case "windows":
		WindowsDeploy()

	case "linux":
		LinuxDeploy()

	case "darwin":
		MacDeploy()
	}

}
