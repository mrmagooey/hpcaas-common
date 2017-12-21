package common

// ContainerAddresses maps the container id to a string in the format ip<port>
type ContainerAddresses map[int]string

// and what the address:port is for every other container
type extraPort struct {
	InternalPort           int
	Name                   string
	ExternalContainerPorts ContainerAddresses
}

// DaemonInfo What the daemon stores
type DaemonInfo struct {
	// The
	DaemonState *DaemonState `json:"daemonState"`

	// updated by the orchestrator
	CodeParams    *map[string]string `json:"codeParams"`
	CodeName      *string            `json:"codeName"`
	CodeArguments *[]string          `json:"codeArguments"`

	// only written by the daemon
	CodeExitStatus *int    `json:"codeExitStatus"`
	CodeStdout     *string `json:"codeStdout"`
	CodeStderr     *string `json:"codeStderr"`
	CodePID        *int    `json:"codePID"`

	// updated by the orchestrator
	ExtraPorts    *[]extraPort        `json:"extraPorts"`
	SSHAddresses  *ContainerAddresses `json:"sshAddresses"`
	WorldRank     *int                `json:"worldRank"`
	WorldSize     *int                `json:"worldSize"`
	SSHPrivateKey *string             `json:"sshPrivateKey"`
	SSHPublicKey  *string             `json:"sshPublicKey"`

	// written on startup, pulled from env variables
	AuthorizationKey *string `json:"authorizationKey"`
}
