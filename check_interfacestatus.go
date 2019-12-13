package main

import (
	ict "github.com/tdh-foundation/icinga2-go-checktools"
)

type InterfaceStatus struct {
	Port   string
	Name   string
	Status string
	Vlan   string
	Duplex string
	Speed  string
	Type   string
}

type SwitchStatus interface {
	ParseStatus() ([]InterfaceStatus, error)
}

// CheckInterfaceStatus
func CheckInterfaceStatus(host string, username string, password string, identity string, port int) (ict.Icinga, error) {

	var ssh *ict.SSHTools
	var sw *CiscoSwitch

	sw = NewCiscoSwitch(host)

	// Opening a ssh session to the switch
	ssh, err = ict.NewSSHTools(host, username, password, identity, port)
	if err != nil {
		return ict.Icinga{}, err
	}

	// Sending command to the switch and getting returned data
	err = ssh.SendSSH("show interface status")
	if err != nil {
		return ict.Icinga{}, err
	}

	//==== Parsing Stdout data to a structured slice ====//
	err = sw.ParseInterfaceStatus(ssh.Stdout)

	return ict.Icinga{}, nil
}
