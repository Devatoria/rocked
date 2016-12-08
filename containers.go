package rocked

import (
	"io/ioutil"
)

const (
	DockerContainersDir = "/containers"
)

// ListContainers returns the list of existing containers ID
func (r *Rocked) ListContainers() ([]string, error) {
	fi, err := ioutil.ReadDir(r.DockerHome + DockerContainersDir)
	if err != nil {
		return []string{}, err
	}

	var containers []string
	for _, container := range fi {
		containers = append(containers, container.Name())
	}

	return containers, nil
}
