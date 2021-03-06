package controllers

import (
	"fmt"
	"os"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
)

// TekuBeaconNode is ConsenSys Pegasys Ethereum 2.0 client
type TekuBeaconNode struct{}

const (
	// EnvTekuBeaconNodeImage is the environment variable used for PegaSys Teku beacon node image
	EnvTekuBeaconNodeImage = "TEKU_BEACON_NODE_IMAGE"
	// DefaultTekuBeaconNodeImage is PegaSys Teku beacon node image
	DefaultTekuBeaconNodeImage = "consensys/teku:21.2.0"
)

// HomeDir returns container home directory
func (t *TekuBeaconNode) HomeDir() string {
	return TekuHomeDir
}

// Args returns command line arguments required for client
func (t *TekuBeaconNode) Args(node *ethereum2v1alpha1.BeaconNode) (args []string) {

	args = append(args, TekuDataPath, PathBlockchainData(t.HomeDir()))

	args = append(args, TekuNetwork, node.Spec.Join)

	if len(node.Spec.Eth1Endpoints) != 0 {
		args = append(args, TekuEth1Endpoint, node.Spec.Eth1Endpoints[0])
	}

	if node.Spec.REST {
		args = append(args, TekuRestEnabled)

		if node.Spec.RESTPort != 0 {
			args = append(args, TekuRestPort, fmt.Sprintf("%d", node.Spec.RESTPort))
		}
		if node.Spec.RESTHost != "" {
			args = append(args, TekuRestHost, node.Spec.RESTHost)
		}
	}

	if node.Spec.P2PPort != 0 {
		args = append(args, TekuP2PPort, fmt.Sprintf("%d", node.Spec.P2PPort))
	}

	return
}

// Command returns command for running the client
func (t *TekuBeaconNode) Command() (command []string) {
	return
}

// Image returns teku docker image
func (t *TekuBeaconNode) Image() string {
	if os.Getenv(EnvTekuBeaconNodeImage) == "" {
		return DefaultTekuBeaconNodeImage
	}
	return os.Getenv(EnvTekuBeaconNodeImage)
}
