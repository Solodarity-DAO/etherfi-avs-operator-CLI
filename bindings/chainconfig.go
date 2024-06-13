package bindings

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {

	// Eigenlayer core contracts
	AvsDirectoryAddress common.Address
	DelegationManager   common.Address
	StrategyManager     common.Address
	EigenpodManager     common.Address
	BeaconEthStrategy   common.Address
	WethStrategy        common.Address

	OperatorManagerAddress     common.Address
	EigenDARegistryCoordinator common.Address
	EigenDAServiceManager      common.Address
	EOracleRegistryCoordinator common.Address
	EOracleServiceManager      common.Address
	BrevisRegistryCoordinator  common.Address
	BrevisServiceManager       common.Address
	LagrangeService            common.Address
}

var Mainnet = Config{
	AvsDirectoryAddress: common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"),
	DelegationManager:   common.HexToAddress("0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A"),
	StrategyManager:     common.HexToAddress("0x858646372CC42E1A627fcE94aa7A7033e7CF075A"),
	EigenpodManager:     common.HexToAddress("0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338"),
	BeaconEthStrategy:   common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
	WethStrategy:        common.HexToAddress(""),

	OperatorManagerAddress:     common.HexToAddress("0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a"),
	EigenDARegistryCoordinator: common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
	EigenDAServiceManager:      common.HexToAddress("0x870679E138bCdf293b7Ff14dD44b70FC97e12fc0"),
	EOracleRegistryCoordinator: common.HexToAddress("0x757E6f572AfD8E111bD913d35314B5472C051cA8"),
	EOracleServiceManager:      common.HexToAddress("0x23221c5bB90C7c57ecc1E75513e2E4257673F0ef"),
	BrevisRegistryCoordinator:  common.HexToAddress("0x434621cfd8BcDbe8839a33c85aE2B2893a4d596C"),
	BrevisServiceManager:       common.HexToAddress("0x9FC952BdCbB7Daca7d420fA55b942405B073A89d"),
	LagrangeService:            common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2"),
}

var Holesky = Config{
	AvsDirectoryAddress: common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	DelegationManager:   common.HexToAddress("0xA44151489861Fe9e3055d95adC98FbD462B948e7"),
	StrategyManager:     common.HexToAddress("0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6"),
	EigenpodManager:     common.HexToAddress("0x30770d7E3e71112d7A6b7259542D1f680a70e315"),
	BeaconEthStrategy:   common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
	WethStrategy:        common.HexToAddress("0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9"),

	OperatorManagerAddress:     common.HexToAddress("0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4"),
	EigenDARegistryCoordinator: common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"),
	EigenDAServiceManager:      common.HexToAddress("0xD4A7E1Bd8015057293f0D0A557088c286942e84b"),
	BrevisRegistryCoordinator:  common.HexToAddress("0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca"),
	BrevisServiceManager:       common.HexToAddress("0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F"),
}

func ConfigForChain(chainID int64) (*Config, error) {

	var cfg Config
	switch chainID {
	case 1:
		cfg = Mainnet
	case 17000:
		cfg = Holesky
	default:
		return nil, fmt.Errorf("unimplemented chain: %d", chainID)
	}
	return &cfg, nil
}

func AutodetectConfig(rpcClient *ethclient.Client) (*Config, error) {
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("querying chainID from RPC: %w", err)
	}
	return ConfigForChain(chainID.Int64())
}
