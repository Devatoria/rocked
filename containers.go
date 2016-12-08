package rocked

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	DockerContainersDir           = "/containers"
	DockerContainerConfigFileName = "config.v2.json"
)

// Container represents a container config JSON
type Container struct {
	State struct {
		Running           bool        `json:"Running"`
		Paused            bool        `json:"Paused"`
		Restarting        bool        `json:"Restarting"`
		OOMKilled         bool        `json:"OOMKilled"`
		RemovalInProgress bool        `json:"RemovalInProgress"`
		Dead              bool        `json:"Dead"`
		Pid               int         `json:"Pid"`
		StartedAt         time.Time   `json:"StartedAt"`
		FinishedAt        time.Time   `json:"FinishedAt"`
		Health            interface{} `json:"Health"`
	} `json:"State"`
	ID      string    `json:"ID"`
	Created time.Time `json:"Created"`
	Managed bool      `json:"Managed"`
	Path    string    `json:"Path"`
	Args    []string  `json:"Args"`
	Config  struct {
		Hostname     string              `json:"Hostname"`
		Domainname   string              `json:"Domainname"`
		User         string              `json:"User"`
		AttachStdin  bool                `json:"AttachStdin"`
		AttachStdout bool                `json:"AttachStdout"`
		AttachStderr bool                `json:"AttachStderr"`
		ExposedPorts map[string]struct{} `json:"ExposedPorts"`
		Tty          bool                `json:"Tty"`
		OpenStdin    bool                `json:"OpenStdin"`
		StdinOnce    bool                `json:"StdinOnce"`
		Env          []string            `json:"Env"`
		Cmd          []string            `json:"Cmd"`
		Image        string              `json:"Image"`
		Volumes      interface{}         `json:"Volumes"`
		WorkingDir   string              `json:"WorkingDir"`
		Entrypoint   interface{}         `json:"Entrypoint"`
		OnBuild      interface{}         `json:"OnBuild"`
		Labels       struct {
		} `json:"Labels"`
	} `json:"Config"`
	Image           string `json:"Image"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		SandboxID              string `json:"SandboxID"`
		HairpinMode            bool   `json:"HairpinMode"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		Networks               struct {
			Bridge struct {
				IPAMConfig          interface{} `json:"IPAMConfig"`
				Links               interface{} `json:"Links"`
				Aliases             interface{} `json:"Aliases"`
				NetworkID           string      `json:"NetworkID"`
				EndpointID          string      `json:"EndpointID"`
				Gateway             string      `json:"Gateway"`
				IPAddress           string      `json:"IPAddress"`
				IPPrefixLen         int         `json:"IPPrefixLen"`
				IPv6Gateway         string      `json:"IPv6Gateway"`
				GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
				GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
				MacAddress          string      `json:"MacAddress"`
			} `json:"bridge"`
		} `json:"Networks"`
		Service                interface{}         `json:"Service"`
		Ports                  map[string]struct{} `json:"Ports"`
		SandboxKey             string              `json:"SandboxKey"`
		SecondaryIPAddresses   interface{}         `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses interface{}         `json:"SecondaryIPv6Addresses"`
		IsAnonymousEndpoint    bool                `json:"IsAnonymousEndpoint"`
	} `json:"NetworkSettings"`
	LogPath                string              `json:"LogPath"`
	Name                   string              `json:"Name"`
	Driver                 string              `json:"Driver"`
	MountLabel             string              `json:"MountLabel"`
	ProcessLabel           string              `json:"ProcessLabel"`
	RestartCount           int                 `json:"RestartCount"`
	HasBeenStartedBefore   bool                `json:"HasBeenStartedBefore"`
	HasBeenManuallyStopped bool                `json:"HasBeenManuallyStopped"`
	MountPoints            map[string]struct{} `json:"MountPoints"`
	AppArmorProfile        string              `json:"AppArmorProfile"`
	HostnamePath           string              `json:"HostnamePath"`
	HostsPath              string              `json:"HostsPath"`
	ShmPath                string              `json:"ShmPath"`
	ResolvConfPath         string              `json:"ResolvConfPath"`
	SeccompProfile         string              `json:"SeccompProfile"`
	NoNewPrivileges        bool                `json:"NoNewPrivileges"`
}

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

// InspectContainer inspects a container by reading its config JSON file
func (r *Rocked) InspectContainer(id string) (Container, error) {
	// Check if container exists
	containerPath := r.DockerHome + DockerContainersDir + "/" + id
	_, err := os.Stat(containerPath)
	if err != nil {
		return Container{}, err
	}

	// Read config file
	data, err := ioutil.ReadFile(strings.Join([]string{containerPath, DockerContainerConfigFileName}, "/"))
	if err != nil {
		return Container{}, err
	}

	// Unmarshal
	var containerInfo Container
	err = json.Unmarshal(data, &containerInfo)

	return containerInfo, err
}
