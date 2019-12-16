// check_ciscoswitch is a icinga/nagios check plugin who
// get information about interface status of CISCO Switch (currently tested with Catalyst and Nexus Series switch)
// check return critical if one of interfaces has status "err-disabled"
package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	ict "github.com/tdh-foundation/icinga2-go-checktools"
	"os"
	"strconv"
	"strings"
)

var (
	arguments  docopt.Opts
	err        error
	buildcount string
	usage      string

	paramsstruct {
		command    string
		host       string
		port       int
		username   string
		password   string
		identity   string
		version    bool
		verbose    bool
		switchType string
	}
)

// Init parsing program arguments
func init() {

	usage = `check_ciscoswitch
Check CISCO Switch port status
Usage: 
	check_ciscoswitch (-h | --help | --version)
	check_ciscoswitch status (-H <host> | --host=<host> -u <username> | --username=<username>) [-p <password> | --password=<password> | -i <pkey_file> | --identity=<pkey_file] [-P <port> | --port=<port>] 
Options:
	--version  				Show check_cattools version.
	-h --help  				Show this screen.
	-v --verbose  	Verbose mode
	-H <host> --host=<host>  		Switch hostname or IP Address
	-u <username> --username=<username>  	Username
	-p <password> --password=<password>  	Password
	-i <pkey_file> --identity=<pkey_file>  	Private key file [default: ~/.ssh/id_rsa]
	-P <port> --port=<port>  		Port number [default: 22]
	-s <switch_type> --switch-type=<switch_type> Switch Type (CISCO)
`
	// Don't parse command line argument for testing argument must be passed with OS environment variable
	if os.Getenv("CHECK_MODE") == "TEST" {
		params.version, _ = strconv.ParseBool(os.Getenv("VERSION"))
		params.port, _ = strconv.Atoi(os.Getenv("PORT"))
		if params.port == 0 {
			params.port = 22
		}
		params.host = os.Getenv("HOST")
		params.username = os.Getenv("USERNAME")
		params.password = os.Getenv("PASSWORD")
		params.identity = os.Getenv("IDENTITY")
		if params.identity == "" && params.password == "" {
			params.identity = "~/.ssh/id_rsa"
		}
		params.verbose, _ = strconv.ParseBool(os.Getenv("VERBOSE"))
		params.command = os.Getenv("COMMAND")
		params.switchType = os.Getenv("SWITCH_TYPE")
	} else {
		arguments, err = docopt.ParseDoc(usage)
		if err != nil {
			fmt.Printf("%s: Error parsing command line arguments: %v", ict.UnkMsg, err)
			os.Exit(ict.UnkExit)
		}

		if c, _ := arguments.Bool("status"); c {
			params.command = "status"
		}

		params.version, _ = arguments.Bool("--version")
		params.port, _ = arguments.Int("--port")
		params.host, _ = arguments.String("--host")
		params.username, _ = arguments.String("--username")
		params.password, _ = arguments.String("--password")
		params.identity, _ = arguments.String("--identity")
		params.verbose, _ = arguments.Bool("--verbose")
		params.switchType, _ = arguments.String("--switch-type")
	}
}

func main() {
	var err error
	var icinga ict.Icinga
	var sw ict.SwitchInterface

	//TODO: Currently check are implemented to check CISCO 2960 and CISCO Nexu switch
	//		in this place we can implement other Switch Brand and models
	switch strings.ToUpper(params.switchType) {
	case "CISCO":
		sw = NewCiscoSwitch(params.host)
	default:
		sw = NewCiscoSwitch(params.host)
	}

	// We return version of program and exit with Ok status
	if params.version {
		fmt.Printf("check_ciscoswitch version 0.0.0-build %s\n", buildcount)
		os.Exit(ict.OkExit)
	}

	// Check command arguments and calling method
	switch params.command {
	case "status":
		icinga, err = sw.CheckInterfaceStatus(params.host, params.username, params.password, params.identity, params.port)
		if err != nil {
			fmt.Printf("%s: Error CheckInterfaceStatus => %s", ict.CriMsg, err)
			os.Exit(ict.CriExit)
		}

		fmt.Println(icinga)
		os.Exit(icinga.Exit)
	default:
		fmt.Printf("check_ciscoswitch version 0.0.0-build %s\n", buildcount)
		fmt.Printf("Usage: %s", usage)
		os.Exit(ict.CriExit)
	}

}
