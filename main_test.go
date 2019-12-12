package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCheck_InterfaceStatus(t *testing.T) {
	_, err := Check_InterfaceStatus(params.host, params.username, params.password, params.identity, params.port)
	if err != nil {
		t.Errorf("Error Check_InterfaceStatus: %s", err)
	}
}
