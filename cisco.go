// This file content implementation of methods to check CISCO 2960/Nexus switch
// interfaces status
package main

import (
	"fmt"
	ict "github.com/tdh-foundation/icinga2-go-checktools"
	"regexp"
	"strings"
)

//noinspection ALL
const (
	Up        = 1
	Down      = 0
	Exception = -1
)

var (
	Connected    = regexp.MustCompile(`(?i)^connected$`)
	NotConnected = regexp.MustCompile(`(?i)^notconnect?$`)
	Disabled     = regexp.MustCompile(`(?i)^disabled$`)
	ErrDisabled  = regexp.MustCompile(`(?i)^err-dis[a-zA-Z]+$`)
	XcrvrAbsen   = regexp.MustCompile(`(?i)^xcvrabsen$`)
	NoOperMem    = regexp.MustCompile(`(?i)^noOpermem$`)
	DownStatus   = regexp.MustCompile(`(?i)^down$`)

	// Status condtions
	OkCondition  = []*regexp.Regexp{Connected, NotConnected, Disabled, DownStatus, XcrvrAbsen}
	CriCondition = []*regexp.Regexp{ErrDisabled}
	WarCondition = []*regexp.Regexp{NoOperMem}

	// Metric conditions
	UpCondition   = []*regexp.Regexp{Connected}
	DownCondition = []*regexp.Regexp{NotConnected, Disabled, DownStatus, XcrvrAbsen}
)

// CiscoSwitchStatus implement SwitchStatus Interface
type CiscoSwitchStatus ict.SwitchStatus

// Instantiate a new CiscoSwitchStatus
func NewCiscoSwitch(name string) *CiscoSwitchStatus {
	cs := new(CiscoSwitchStatus)
	cs.Name = name
	return cs
}

// getStatusCondition is a private method who return Icinga Exit condition value
func getStatusCondition(status ict.SwitchInterfaceStatus) int {
	for i := range OkCondition {
		if OkCondition[i].MatchString(status.Status) {
			return ict.OkExit
		}
	}
	for i := range CriCondition {
		if CriCondition[i].MatchString(status.Status) {
			return ict.CriExit
		}
	}
	for i := range WarCondition {
		if WarCondition[i].MatchString(status.Status) {
			return ict.WarExit
		}
	}
	return ict.UnkExit
}

// getMetricCondition is a private method who return information about usage of interface for Icinga Metric
func getMetricCondition(status ict.SwitchInterfaceStatus) int {
	for i := range UpCondition {
		if UpCondition[i].MatchString(status.Status) {
			return Up
		}
	}
	for i := range DownCondition {
		if DownCondition[i].MatchString(status.Status) {
			return Down
		}
	}
	return Exception
}

// CheckInterfaceStatus
func (cSwitchStatus *CiscoSwitchStatus) CheckInterfaceStatus(host string, username string, password string, identity string, port int) (ict.Icinga, error) {

	var ssh *ict.SSHTools

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

	// Parsing Stdout data to a structured slice
	err = cSwitchStatus.ParseInterfaceStatus(ssh.Stdout)
	if err != nil {
		err = fmt.Errorf("error ParseInterfaceStatus: %v", err)
		return ict.Icinga{}, err
	}

	// Generate Icinga result "{status}:[Message][| Metric]"
	return cSwitchStatus.ReturnIcingaResult(), nil
}

// ParseInterfaceStatus from response received from Cisco Switch Request
//noinspection GoNilness
func (cSwitchStatus *CiscoSwitchStatus) ParseInterfaceStatus(response string) error {

	// Clearing/resetting respStatus slice
	cSwitchStatus.SwStatus = cSwitchStatus.SwStatus[:0]

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
	re := regexp.MustCompile(`(?i)(?P<Port>Port\s+)(?P<Name>Name\s+)(?P<status>status\s+)(?P<Vlan>Vlan\s+)(?P<Duplex>Duplex\s+)(?P<Speed>Speed\s+)(?P<Type>Type\s?)`)
	reSplit := regexp.MustCompile(`\s+`)
	borders := re.FindStringSubmatchIndex(respStatus[0])

	if borders == nil || len(borders) != 16 {
		return fmt.Errorf("error parsing Cisco interface header not found in response")
	}

	// converting string to structured data
	for _, s := range respStatus[1:] {
		item := ict.SwitchInterfaceStatus{}
		item.Port = strings.Trim(s[borders[2]:borders[3]], " \r")
		item.Name = strings.Trim(s[borders[4]:borders[5]], " \r")
		item.Status = strings.Trim(s[borders[6]:borders[7]], " \r")

		// For Cisco 2960X output, Duplex and speed are right justified
		vds := reSplit.Split(strings.Trim(s[borders[8]:borders[13]], " \r"), -1) //Vlan-Duplex-Speed
		if len(vds) != 3 {
			return fmt.Errorf("error parsing Vlan/Duplex/Speed in Cisco interface respStatus")
		}
		item.Vlan = strings.Trim(vds[0], " \r")
		item.Duplex = strings.Trim(vds[1], " \r")
		item.Speed = strings.Trim(vds[2], " \r")

		item.Type = strings.Trim(s[borders[14]:], " \r")
		cSwitchStatus.SwStatus = append(cSwitchStatus.SwStatus, item)
	}

	return nil
}

// ReturnIcingaResult convert
func (cSwitchStatus *CiscoSwitchStatus) ReturnIcingaResult() ict.Icinga {
	// Counter for metric
	var up, down, exception, all int

	// Initialize default Icinga status
	icinga := ict.Icinga{Message: ict.UnkMsg, Exit: ict.UnkExit, Metric: ""}

	// Analyze each interface status - worst status give global status
	var tmpMsg string
	for _, item := range cSwitchStatus.SwStatus {
		tmpMsg = ""
		switch getStatusCondition(item) {
		case ict.OkExit:
			if icinga.Exit == ict.UnkExit {
				icinga.Exit = ict.OkExit
			}
			if Disabled.MatchString(item.Status) {
				tmpMsg = item.Port + " Info[" + item.Status + "]"
			}
		case ict.WarExit:
			if icinga.Exit != ict.CriExit {
				icinga.Exit = ict.WarExit
			}
			tmpMsg = item.Port + " Warning[" + item.Status + "]"
		case ict.CriExit:
			icinga.Exit = ict.CriExit
			tmpMsg = item.Port + " Critical[" + item.Status + "]"
		default:
			if icinga.Exit != ict.CriExit && icinga.Exit != ict.WarExit {
				icinga.Exit = ict.UnkExit
			}
			tmpMsg = item.Port + " Unknown[" + item.Status + "]"
		}
		if icinga.Message == ict.UnkMsg || icinga.Message == "" {
			icinga.Message = tmpMsg
		} else {
			if tmpMsg != "" {
				icinga.Message += " / " + tmpMsg
			}
		}
		switch getMetricCondition(item) {
		case Up:
			up++
		case Down:
			down++
		default:
			exception++
		}
	}

	all = up + down + exception
	icinga.Metric = fmt.Sprintf("'Exception'=%d;;;0;%d 'Up'=%d;;;0;%d 'Down'=%d;;;0;%d", exception, all, up, all, down, all)

	// Everything is ok and no messages is set just say that
	if icinga.Message == "" {
		icinga.Message = "All is Ok..."
	}

	return icinga
}

// Status return SwitchInterfaceStatusArray
func (cSwitchStatus *CiscoSwitchStatus) Status() []ict.SwitchInterfaceStatus {
	return cSwitchStatus.SwStatus
}
