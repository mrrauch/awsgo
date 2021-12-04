package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	instanceID := os.Args[2]

}



}

func execute(){
	out, err := exec.Command("echo HelloWorld").Output()

	if err != nil {
		fmt.Printf(%s, err)
	}

	fmt.Println("Command Output Success")
	output := string(out[:])
	fmt.Println(output)

}

