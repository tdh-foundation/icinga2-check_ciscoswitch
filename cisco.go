package main

import (
	"fmt"
	"regexp"
	"strings"
)

//noinspection ALL
const (
	Connected    = "connected"
	NotConnected = "notconnect"
	Disabled     = "disabled"
	ErrDisabled  = "err-disabled"
)

type CiscoSwitch struct {
	name   string
	status []InterfaceStatus
}

func NewCiscoSwitch(name string) *CiscoSwitch {
	return &CiscoSwitch{name, []InterfaceStatus{}}
}

// ParseInterfaceStatus from response received from Cisco Switch Request
func (sw CiscoSwitch) ParseInterfaceStatus(response string) error {

	// Clearing/resetting respStatus slice
	sw.status = sw.status[:0]

	// Converting multi-line string to slice of string
	respStatus := strings.Split(response, "\n")

	// Removing Blank and separator lines
	var tmp []string
	for _, s := range respStatus {
		if s != "" && (len(s) >= 3 && s[0:3] != "---") {
			tmp = append(tmp, s)
		}
	}

	// If slice length is less than 2 rows response is not comprehensive
	respStatus = tmp
	if len(respStatus) < 2 {
		return fmt.Errorf("error parsing Cisco interface respStatus response empty response")
	}

	// Interfaces response is a fixed size column array => finding position end size of each columns based of Header
	re := regexp.MustCompile(`(?i)(?P<Port>Port\s+)(?P<Name>Name\s+)(?P<Status>Status\s+)(?P<Vlan>Vlan\s+)(?P<Duplex>Duplex\s+)(?P<Speed>Speed\s+)(?P<Type>Type\s?)`)
	borders := re.FindStringSubmatchIndex(respStatus[0])

	if borders == nil || len(borders) != 16 {
		return fmt.Errorf("error parsing Cisco interface header not found in response")
	}

	// converting string to structured data
	for _, s := range respStatus[1:] {
		item := InterfaceStatus{}
		item.Port = strings.Trim(s[borders[2]:borders[3]], " \r")
		item.Name = strings.Trim(s[borders[4]:borders[5]], " \r")
		item.Status = strings.Trim(s[borders[6]:borders[7]], " \r")
		item.Vlan = strings.Trim(s[borders[8]:borders[9]], " \r")
		item.Duplex = strings.Trim(s[borders[10]:borders[11]], " \r")
		item.Speed = strings.Trim(s[borders[12]:borders[13]], " \r")
		item.Type = strings.Trim(s[borders[14]:], " \r")
		sw.status = append(sw.status, item)
	}

	return nil
}
