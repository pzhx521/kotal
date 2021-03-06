package controllers

import (
	"fmt"
	"os"
	"strings"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
)

// LighthouseBeaconNode is SigmaPrime Ethereum 2.0 client
type LighthouseBeaconNode struct{}

// Images
const (
	// EnvLighthouseBeaconNodeImage is the environment variable used for SigmaPrime Ethereum 2.0 beacon node image
	EnvLighthouseBeaconNodeImage = "LIGHTHOUSE_BEACON_NODE_IMAGE"
	// DefaultLighthouseBeaconNodeImage is the default SigmaPrime Ethereum 2.0 beacon node image
	DefaultLighthouseBeaconNodeImage = "sigp/lighthouse:v1.1.3"
)

// HomeDir returns container home directory
func (t *LighthouseBeaconNode) HomeDir() string {
	return LighthouseHomeDir
}

// Args returns command line arguments required for client
func (t *LighthouseBeaconNode) Args(node *ethereum2v1alpha1.BeaconNode) (args []string) {

	args = append(args, LighthouseDataDir, PathBlockchainData(t.HomeDir()))

	args = append(args, LighthouseNetwork, node.Spec.Join)

	if len(node.Spec.Eth1Endpoints) != 0 {
		args = append(args, LighthouseEth1)
		args = append(args, LighthouseEth1Endpoints, strings.Join(node.Spec.Eth1Endpoints, ","))
	}

	if node.Spec.REST {
		args = append(args, LighthouseHTTP)

		if node.Spec.RESTPort != 0 {
			args = append(args, LighthouseHTTPPort, fmt.Sprintf("%d", node.Spec.RESTPort))
		}
		if node.Spec.RESTHost != "" {
			args = append(args, LighthouseHTTPAddress, node.Spec.RESTHost)
		}
	}

	if node.Spec.P2PPort != 0 {
		args = append(args, LighthousePort, fmt.Sprintf("%d", node.Spec.P2PPort))
		args = append(args, LighthouseDiscoveryPort, fmt.Sprintf("%d", node.Spec.P2PPort))
	}

	return
}

// Command returns command for running the client
func (t *LighthouseBeaconNode) Command() (command []string) {
	command = []string{"lighthouse", "bn"}
	return
}

// Image returns prysm docker image
func (t *LighthouseBeaconNode) Image() string {
	if os.Getenv(EnvLighthouseBeaconNodeImage) == "" {
		return DefaultLighthouseBeaconNodeImage
	}
	return os.Getenv(EnvLighthouseBeaconNodeImage)
}
