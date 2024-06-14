package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
)

/*
// Needs a bunch of refactoring
func registerBLS(ctx context.Context, cli *cli.Command) error {

	var cfg bindings.Config
	chainID := cli.Int("chain-id")
	switch chainID {
	case 1:
		cfg = bindings.Mainnet
	case 17000:
		cfg = bindings.Holesky
	default:
		return fmt.Errorf("unimplemented chain: %d", chainID)
	}

	// load key for signing tx
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	// toggle whether tx is broadcast to network
	transactor.NoSend = !cli.Bool("broadcast")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	// bind rpc to contract abi
	operatorContract, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding contract: %w", err)
	}

	// TODO: validate params
	operatorID := big.NewInt(cli.Int("operator-id"))
	registryCoordinator := common.HexToAddress(cli.String("registry-coordinator"))
	socket := cli.String("socket")
	var quorumNumbers []uint8
	for _, v := range cli.IntSlice("quorum-numbers") {
		quorumNumbers = append(quorumNumbers, uint8(v))
	}

	// load bls signature and convert to format expected by contract
	params, err := BLSJsonToRegistrationParams(cli.String("bls-signature-file"))
	if err != nil {
		return fmt.Errorf("parsing bls signature file: %w", err)
	}

	// Sign transaction and broadcast if requested
	tx, err := operatorContract.RegisterBlsKeyAsDelegatedNodeOperator(transactor, operatorID, registryCoordinator, quorumNumbers, socket, *params)
	if err != nil {
		return fmt.Errorf("failed to sign and/or broadcast tx: %w", err)
	}
	var buf bytes.Buffer
	tx.EncodeRLP(&buf)
	fmt.Printf("raw tx: %s\n", hex.EncodeToString(buf.Bytes()))

	return nil
}
*/

func BLSJsonToRegistrationParams(filepath string) (*contracts.IBLSApkRegistryPubkeyRegistrationParams, error) {

	buf, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading signature file: %w", err)
	}

	var input types.BLSPubkeyRegistrationParams
	if err := json.Unmarshal(buf, &input); err != nil {
		return nil, fmt.Errorf("unmarshalling signature file: %w", err)
	}

	/*
		g1x, _ := big.NewInt(0).SetString(input.G1.X, 10)
		g1y, _ := big.NewInt(0).SetString(input.G1.Y, 10)

		g2x0, _ := big.NewInt(0).SetString(input.G2.X[0], 10)
		g2x1, _ := big.NewInt(0).SetString(input.G2.X[1], 10)
		g2y0, _ := big.NewInt(0).SetString(input.G2.Y[0], 10)
		g2y1, _ := big.NewInt(0).SetString(input.G2.Y[1], 10)

		sigX, _ := big.NewInt(0).SetString(input.Signature.X, 10)
		sigY, _ := big.NewInt(0).SetString(input.Signature.Y, 10)
	*/

	contractParams := contracts.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: contracts.BN254G1Point{
			X: input.Signature.X,
			Y: input.Signature.Y,
		},
		PubkeyG1: contracts.BN254G1Point{
			X: input.G1.X,
			Y: input.G1.Y,
		},
		PubkeyG2: contracts.BN254G2Point{
			X: input.G2.X,
			Y: input.G2.Y,
		},
	}

	return &contractParams, nil
}
