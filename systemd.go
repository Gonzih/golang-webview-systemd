package main

import (
	"fmt"

	"github.com/coreos/go-systemd/dbus"
)

type SystemState struct {
	Units []dbus.UnitStatus `json:"units"`
	Err   error             `json:"err"`
}

func (state *SystemState) Search(search string, user bool) {
	var conn *dbus.Conn
	var err error

	if user {
		conn, err = dbus.NewUserConnection()
	} else {
		conn, err = dbus.New()
	}

	if err != nil {
		state.Err = err
		return
	}

	defer conn.Close()

	if len(search) > 0 {
		wildcardSearch := fmt.Sprintf("%s*", search)

		state.Units, state.Err = conn.ListUnitsByPatterns([]string{}, []string{wildcardSearch, search})
	} else {
		state.Units, state.Err = conn.ListUnits()
	}
}
