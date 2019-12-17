# check_ciscoswitch
## Description
check_ciscoswitch is Icinga2/nagios monitoring plugin to check interface status of cisco switch.
Connection to switch is made by SSH and user must be defined on switch (defined user must have privilege level 5).

Plugin is currently tested on CISCO Catalyst 2960X Series and Cisco Nexus 3100 Series

This plugin is written in Go build with (Go version 1.13)
 
##  Installation
Plugin can be built with "build.sh" on Linux system or "build.cmd" on Windows. 
Executable must be copied to Icinga/nagios plugins directory (usualy on /usr/lib/nagios/plugins)

## Testing
Test file contain mock data and online test. Before running test following environment variable must be defined:
 
    CHECK_MODE=TEST
    HOST={Switch IP address or DNS name}
    USERNAME={Cisco switch username}
    COMMAND=status

### Running test


## Usage
`check_ciscoswitch (-h | --help | --version)`

   `check_ciscoswitch status (-H <host> | --host=<host>) (-u <username> | --username=<username>) [-p <password> | --password=<password> | -i <pkey_file> | --identity=<pkey_file] [-P <port> | --port=<port>]` 

   	--version				        Show check_cattools version.
	-h --help  		    		        Show this screen.
	-v --verbose  	                                Verbose mode
	-H <host> --host=<host>  		        Switch hostname or IP Address
	-u <username> --username=<username>  	        Username
	-p <password> --password=<password>  	        Password
	-i <pkey_file> --identity=<pkey_file>  	        Private key file [default: ~/.ssh/id_rsa]
	-P <port> --port=<port>  		        Port number [default: 22]
	-s <switch_type> --switch-type=<switch_type>    Switch Type (CISCO)

## Example
### Command
`./check_ciscoswitch status --host=10.10.100.42 --username=icinga`
### Result
`OK: Fa0 Info[disabled] |'Exception'=0;;;0;102 'Up'=39;;;0;102 'Down'=63;;;0;102`
