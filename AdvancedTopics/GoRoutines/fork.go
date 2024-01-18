package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	//cmd := exec.Command("echo", "Hello, Windows!")
	//cmd := exec.Command("cmd", "/c", "dir")
	//cmd := exec.Command("systeminfo")
	//cmd := exec.Command("ipconfig", "/all")
	//cmd := exec.Command("powershell", "-Command", "Get-Process | Select-Object Name, Id")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("output : ", string(output))
}
