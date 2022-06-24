package slaterest

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/senv"
)

const (
	// ContainerID defines the slate.slaterest package container entry id base string.
	ContainerID = slate.ContainerID + ".slaterest"

	// ContainerEngineID defines the default id used to register the
	// application gin engine instance in the application container.
	ContainerEngineID = ContainerID + ".engine"
)

const (
	// EnvID defines the slate.slaterest package base environment variable name.
	EnvID = slate.EnvID + "_SREST"
)

var (
	// ConfigPortPath contains the config path of the server port to be used.
	ConfigPortPath = senv.String(EnvID+"_CONFIG_PORT_PATH", "server.port")

	// LogChannel contains the name of the logging channel to be used on
	// the slaterest application messages.
	LogChannel = senv.String(EnvID+"_LOG_CHANNEL", "exec")
)
