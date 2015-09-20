package main

import (
	"log"
	"os/exec"
)

func main() {
	var cmd *exec.Cmd

	cmd = exec.Command("go", "get")
	if err := cmd.Run(); err != nil {
		log.Fatalf(err.Error())
	}
	//	cmd = exec.Command("go", "fmt")
	//	if err := cmd.Run(); err != nil {
	//		log.Fatalf(err)
	//	}
	//	cmd = exec.Command("go", "vet")
	//	if err := cmd.Run(); err != nil {
	//		log.Fatalf(err)
	//	}
	cmd = exec.Command("go", "build")
	if err := cmd.Run(); err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("build successfully")
}
