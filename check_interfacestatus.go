package main

import (
	"fmt"
	ict "github.com/tdh-foundation/icinga2-go-checktools"
	"regexp"
	"strings"
)

// Check_InterfaceStatus
func Check_InterfaceStatus(host string, username string, password string, identity string, port int) (ict.Icinga, error) {

	var ssh *ict.SSHTools
	var interfaces []CiscoInterfaceStatus

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
	// Converting multi-line string to slice of string
	status := strings.Split(ssh.Stdout, "\n")

	// Removing Blank and separator lines
	tmp := []string{}
	for _, s := range status {
		if s != "" && (len(s) >= 3 && s[0:3] != "---") {
			tmp = append(tmp, s)
		}
	}
	status = tmp

	// Interfaces status are fixed size column - finding position end size of each columns based of Header
	re := regexp.MustCompile(`(?i)(?P<Port>Port\s+)(?P<Name>Name\s+)(?P<Status>Status\s+)(?P<Vlan>Vlan\s+)(?P<Duplex>Duplex\s+)(?P<Speed>Speed\s+)(?P<Type>Type\s?)`)
	borders := re.FindStringSubmatchIndex(status[0])

	// converting string to structured data
	for i, s := range status[1:] {
		item := CiscoInterfaceStatus{}
		item.Port = strings.Trim(s[borders[2]:borders[3]], " \r")
		item.Name = strings.Trim(s[borders[4]:borders[5]], " \r")
		item.Status = strings.Trim(s[borders[6]:borders[7]], " \r")
		item.Vlan = strings.Trim(s[borders[8]:borders[9]], " \r")
		item.Duplex = strings.Trim(s[borders[10]:borders[11]], " \r")
		item.Speed = strings.Trim(s[borders[12]:borders[13]], " \r")
		item.Type = strings.Trim(s[borders[14]:], " \r")
		interfaces = append(interfaces, item)
		fmt.Printf("%4d: %s\t%s\t%s\n", i+1, item.Port, item.Status, item.Duplex)
	}

	return ict.Icinga{}, nil
}
