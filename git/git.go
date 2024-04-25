package git

import (
	"log"
	"os/exec"
)

func GetUserName() string {
	cmd := exec.Command("git", "config", "user.name")
	command, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	name := string(command)
	return name
}