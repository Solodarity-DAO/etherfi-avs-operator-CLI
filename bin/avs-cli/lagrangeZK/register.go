package lagrangezk

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	lagrangezk "github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/lagrangeZK"
	"github.com/urfave/cli/v3"
)

var RegisterCmd = &cli.Command{
	Name:   "register",
	Usage:  "(Admin) Register target operator to LagrangeZK Coprocessor AVS",
	Action: handleRegister,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required registration data
	var input lagrangezk.RegistrationInfo
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input, missing operatorID")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// load eip-1271 admin signing key
	signingKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return lagrangeZKAPI.RegisterOperator(operator, input, signingKey)
}