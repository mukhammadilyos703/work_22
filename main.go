package main

import (
	"log"
	"os"

	git "github.com/mukhammadilyos703/work_22.git/git"
)

func main() {
	name := git.GetUserName()
	email := git.GetUserEmail()
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(name); err != nil {
		log.Fatal(err)
		return
	}

	if _, err := file.WriteString(email); err != nil {
		log.Fatal(err)
		return
	}
}
