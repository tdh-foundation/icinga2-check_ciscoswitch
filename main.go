package main

import (
	"fmt"
	ct "github.com/tdh-foundation/icinga2-go-checktools"
	"os"
)

func main() {
	var ict *ct.SSHTools
	var err error

	ict, err = ct.NewSSHTools("10.10.100.31", "icinga", "", "~/.ssh/id_rsa", 22)
	if err != nil {
		fmt.Printf("Error establishing SSH connection to 10.10.100.31 => %s", err)
	}
	err = ict.SendSSH("show interface status")
	if err != nil {
		fmt.Printf("Error sending command: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", ict.Stdout)

}
