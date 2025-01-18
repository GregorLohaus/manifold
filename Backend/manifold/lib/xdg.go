package lib

import (
	"errors"
	"os"
	"path"
)

type XDG_VAR int

const (
	XDG_CONFIG_HOME XDG_VAR = 1 + iota
	XDG_DATA_HOME
	XDG_HOME
)

func GetXDGVar(name XDG_VAR) (*string, error) {
	switch name {
	case XDG_CONFIG_HOME:
		val := os.Getenv("XDG_CONFIG_HOME")
		if val != "" {
			return &val, nil
		} else {
			return xdgDefault(name)
		}
	case XDG_DATA_HOME:
		val := os.Getenv("XDG_DATA_HOME")
		if val != "" {
			return &val, nil
		} else {
			return xdgDefault(name)
		}
	case XDG_HOME:
		val := os.Getenv("XDG_HOME")
		if val != "" {
			return &val, nil
		} else {
			return xdgDefault(name)
		}
	}
	return nil, errors.New("Invalid XDG_VAR")
}

func xdgDefault(name XDG_VAR) (*string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return nil, errors.New("$HOME not set.")
	}
	switch name {
	case XDG_CONFIG_HOME:
		val := path.Join(home, "/.config")
		return &val, nil
	case XDG_DATA_HOME:
		val := path.Join(home, "/.local/share")
		return &val, nil
	case XDG_HOME:
		return &home, nil
	}
	return nil, errors.New("Invalid XDG_VAR")
}
