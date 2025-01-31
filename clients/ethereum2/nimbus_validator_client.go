package ethereum2

import (
	"fmt"
	"os"
	"strings"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
	corev1 "k8s.io/api/core/v1"
)

// NimbusValidatorClient is Status Ethereum 2.0 client
// https://github.com/status-im/nimbus-eth2
type NimbusValidatorClient struct {
	validator *ethereum2v1alpha1.Validator
}

// Images
const (
	// EnvNimbusValidatorImage is the environment variable used for Status Ethereum 2.0 validator client image
	EnvNimbusValidatorImage = "NIMBUS_VALIDATOR_CLIENT_IMAGE"
	// DefaultNimbusValidatorImage is the default Status Ethereum 2.0 validator client image
	DefaultNimbusValidatorImage = "kotalco/nimbus:v1.6.0"
)

// HomeDir returns container home directory
func (t *NimbusValidatorClient) HomeDir() string {
	return NimbusHomeDir
}

// Command returns environment variables for the client
func (t *NimbusValidatorClient) Env() []corev1.EnvVar {
	return nil
}

// Args returns command line arguments required for client
func (t *NimbusValidatorClient) Args() (args []string) {

	validator := t.validator

	args = append(args, NimbusNonInteractive)

	args = append(args, argWithVal(NimbusLogging, string(t.validator.Spec.Logging)))

	args = append(args, argWithVal(NimbusDataDir, shared.PathData(t.HomeDir())))

	args = append(args, argWithVal(NimbusValidatorsDir, fmt.Sprintf("%s/kotal-validators/validator-keys", shared.PathData(t.HomeDir()))))

	args = append(args, argWithVal(NimbusSecretsDir, fmt.Sprintf("%s/kotal-validators/validator-secrets", shared.PathData(t.HomeDir()))))

	args = append(args, argWithVal(NimbusBeaconNodes, strings.Join(validator.Spec.BeaconEndpoints, ",")))

	if validator.Spec.Graffiti != "" {
		args = append(args, argWithVal(NimbusGraffiti, validator.Spec.Graffiti))
	}

	return
}

// Command returns command for running the client
func (t *NimbusValidatorClient) Command() (command []string) {
	command = []string{"nimbus_validator_client"}
	return
}

// Image returns prysm docker image
func (t *NimbusValidatorClient) Image() string {
	if os.Getenv(EnvNimbusValidatorImage) == "" {
		return DefaultNimbusValidatorImage
	}
	return os.Getenv(EnvNimbusValidatorImage)
}
