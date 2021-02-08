package installer

import (
	"context"
)

type System string

const (
	SystemOpenHab System = "OpenHab"
	SystemUbuntu  System = "Ubuntu"
)

type Step struct {
	Description string
	FilePath    string
	Content     string
}

type HasInstaller interface {
	InstallersSupport() []System
	InstallerSteps(context.Context, System) ([]Step, error)
}
