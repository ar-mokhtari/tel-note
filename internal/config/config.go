package config

import (
	"tel-note/internal/storage/memory"
)

const (
	Admin Role = iota + 10
	Manager
	Basic
	AdminString = "admin"
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
