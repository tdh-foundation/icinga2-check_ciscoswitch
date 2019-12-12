package main

const (
	Connected    = "connected"
	NotConnected = "notconnect"
	Disabled     = "disabled"
	ErrDisabled  = "err-disabled"
)

type CiscoInterfaceStatus struct {
	Port   string
	Name   string
	Status string
	Vlan   string
	Duplex string
	Speed  string
	Type   string
}
