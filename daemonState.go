package common

// ContainerAddresses maps the container id to a string in the format ip<port>
type ContainerAddresses map[int]string

// and what the address:port is for every other container
type extraPort struct {
	InternalPort           int
	Name                   string
	ExternalContainerPorts ContainerAddresses
}

// DaemonState What the daemon stores
type DaemonState struct {
	CodeParams        *map[string]string  `json:"codeParams"`
	ExtraPorts        *[]extraPort        `json:"extraPorts"`
	CodeName          *string             `json:"codeName"`
	CodeArguments     *[]string           `json:"codeArguments"`
	CodeStatus        *CodeStatus         `json:"codeState"`
	DaemonStatus      *DaemonStatus       `json:"daemonState"`
	ResultStatus      *ResultStatus       `json:"resultState"`
	CodeStartedStatus *StartedStatus      `json:"codeStartedMethod"`
	SSHAddresses      *ContainerAddresses `json:"sshAddresses"`
	WorldRank         *int                `json:"worldRank"`
	WorldSize         *int                `json:"worldSize"`
	ResultsDirectory  *string             `json:"resultsDirectory"`
	ResultsURL        *string             `json:"resultsUrl"`
	CodeExitStatus    *int                `json:"codeExitStatus"`
	AuthorizationKey  *string             `json:"authorizationKey"`
	CodeStdout        *string             `json:"codeStdout"`
	CodeStderr        *string             `json:"codeStderr"`
	CodePID           *int                `json:"codePID"`
	SSHPrivateKey     *string             `json:"sshPrivateKey"`
	SSHPublicKey      *string             `json:"sshPublicKey"`
}
