package main

import (
	"fmt"
	"os/exec"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	cmd := exec.Command("D:/software/DingDing" + "/DingtalkLauncher.exe")
	data, err := cmd.Output()
	checkError(err)
	fmt.Println(string(data))
}
