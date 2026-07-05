package run

import (
	"fmt"

	"github.com/containers/podman/v6/pkg/specgen"
	"github.com/julioln/sandman/config"
)

func Network(spec *specgen.SpecGenerator, containerConfig config.ContainerConfig) {
	// Configure network namespace
	var networkNS specgen.Namespace
	if containerConfig.Run.Network != "" {
		var err error
		if networkNS, _, _, err = specgen.ParseNetworkFlag([]string{containerConfig.Run.Network}); err != nil {
			fmt.Println("Error parsing network, defaulting to none: ", err)
			networkNS.NSMode = specgen.None
		}
	} else if containerConfig.Run.Net {
		networkNS.NSMode = specgen.Pasta
	} else {
		networkNS.NSMode = specgen.None
	}
	spec.NetNS = networkNS
}
