package rocked

import (
	"os"
)

// Rocked represents a configured rocked library
type Rocked struct {
	DockerHome string
}

// NewRocked returns an initialized rocked struct, or an error if the given Docker home can't be red
func NewRocked(dockerHome string) (*Rocked, error) {
	_, err := os.Stat(dockerHome)
	if err != nil {
		return nil, err
	}

	return &Rocked{
		DockerHome: dockerHome,
	}, nil
}
