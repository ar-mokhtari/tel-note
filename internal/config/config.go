package config

import (
	"tel-note/internal/storage/memory"
)

const (
	Regulator Role = iota + 10
	Manager
	Basic
	RegulatorString = "regulator"
)

type (
	Config struct {
		DB Database
	}
	Database struct {
		Memory memory.AllData
	}

	Role           int
	ResponseStatus struct {
		State  bool
		String string
	}
)
