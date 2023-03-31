package rollup

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/wemixkanvas/kanvas/components/node/eth"
)

var (
	ErrBlockTimeZero                 = errors.New("block time cannot be 0")
	ErrMissingChannelTimeout         = errors.New("channel timeout must be set, this should cover at least a L1 block time")
	ErrInvalidProposerWindowSize     = errors.New("proposing window size must at least be 2")
	ErrMissingGenesisL1Hash          = errors.New("genesis L1 hash cannot be empty")
	ErrMissingGenesisL2Hash          = errors.New("genesis L2 hash cannot be empty")
	ErrGenesisHashesSame             = errors.New("achievement get! rollup inception: L1 and L2 genesis cannot be the same")
	ErrMissingGenesisL2Time          = errors.New("missing L2 genesis time")
	ErrMissingBatcherAddr            = errors.New("missing genesis system config batcher address")
	ErrMissingOverhead               = errors.New("missing genesis system config overhead")
	ErrMissingScalar                 = errors.New("missing genesis system config scalar")
	ErrMissingGasLimit               = errors.New("missing genesis system config gas limit")
	ErrMissingBatchInboxAddress      = errors.New("missing batch inbox address")
	ErrMissingDepositContractAddress = errors.New("missing deposit contract address")
	ErrMissingL1ChainID              = errors.New("L1 chain ID must not be nil")
	ErrMissingL2ChainID              = errors.New("L2 chain ID must not be nil")
	ErrChainIDsSame                  = errors.New("L1 and L2 chain IDs must be different")
	ErrL1ChainIDNotPositive          = errors.New("L1 chain ID must be non-zero and positive")
	ErrL2ChainIDNotPositive          = errors.New("L2 chain ID must be non-zero and positive")
)

type Genesis struct {
	// The L1 block that the rollup starts *after* (no derived transactions)
	L1 eth.BlockID `json:"l1"`
	// The L2 block the rollup starts from (no transactions, pre-configured state)
	L2 eth.BlockID `json:"l2"`
	// Timestamp of L2 block
	L2Time uint64 `json:"l2_time"`
	// Initial system configuration values.
	// The L2 genesis block may not include transactions, and thus cannot encode the config values,
	// unlike later L2 blocks.
	SystemConfig eth.SystemConfig `json:"system_config"`
}

type Config struct {
	// Genesis anchor point of the rollup
	Genesis Genesis `json:"genesis"`
	// Seconds per L2 block
	BlockTime uint64 `json:"block_time"`
	// Proposer batches may not be more than MaxProposerDrift seconds after
	// the L1 timestamp of the proposer window end.
	//
	// Note: When L1 has many 1 second consecutive blocks, and L2 grows at fixed 2 seconds,
	// the L2 time may still grow beyond this difference.
	MaxProposerDrift uint64 `json:"max_proposer_drift"`
	// Number of epochs (L1 blocks) per proposer window, including the epoch L1 origin block itself
	ProposerWindowSize uint64 `json:"proposer_window_size"`
	// Number of L1 blocks between when a channel can be opened and when it must be closed by.
	ChannelTimeout uint64 `json:"channel_timeout"`
	// Required to verify L1 signatures
	L1ChainID *big.Int `json:"l1_chain_id"`
	// Required to identify the L2 network and create p2p signatures unique for this chain.
	L2ChainID *big.Int `json:"l2_chain_id"`

	// Note: below addresses are part of the block-derivation process,
	// and required to be the same network-wide to stay in consensus.

	// L1 address that batches are sent to.
	BatchInboxAddress common.Address `json:"batch_inbox_address"`
	// L1 Deposit Contract Address
	DepositContractAddress common.Address `json:"deposit_contract_address"`
	// L1 System Config Address
	L1SystemConfigAddress common.Address `json:"l1_system_config_address"`
}

// ValidateL1Config checks L1 config variables for errors.
func (cfg *Config) ValidateL1Config(ctx context.Context, client L1Client) error {
	// Validate the L1 Client Chain ID
	if err := cfg.CheckL1ChainID(ctx, client); err != nil {
		return err
	}

	// Validate the Rollup L1 Genesis Blockhash
	if err := cfg.CheckL1GenesisBlockHash(ctx, client); err != nil {
		return err
	}

	return nil
}

// ValidateL2Config checks L2 config variables for errors.
func (cfg *Config) ValidateL2Config(ctx context.Context, client L2Client) error {
	// Validate the L2 Client Chain ID
	if err := cfg.CheckL2ChainID(ctx, client); err != nil {
		return err
	}

	// Validate the Rollup L2 Genesis Blockhash
	if err := cfg.CheckL2GenesisBlockHash(ctx, client); err != nil {
		return err
	}

	return nil
}

type L1Client interface {
	ChainID(context.Context) (*big.Int, error)
	L1BlockRefByNumber(context.Context, uint64) (eth.L1BlockRef, error)
}

// CheckL1ChainID checks that the configured L1 chain ID matches the client's chain ID.
func (cfg *Config) CheckL1ChainID(ctx context.Context, client L1Client) error {
	id, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	if cfg.L1ChainID.Cmp(id) != 0 {
		return fmt.Errorf("incorrect L1 RPC chain id %d, expected %d", cfg.L1ChainID, id)
	}
	return nil
}

// CheckL1GenesisBlockHash checks that the configured L1 genesis block hash is valid for the given client.
func (cfg *Config) CheckL1GenesisBlockHash(ctx context.Context, client L1Client) error {
	l1GenesisBlockRef, err := client.L1BlockRefByNumber(ctx, cfg.Genesis.L1.Number)
	if err != nil {
		return err
	}
	if l1GenesisBlockRef.Hash != cfg.Genesis.L1.Hash {
		return fmt.Errorf("incorrect L1 genesis block hash %d, expected %d", cfg.Genesis.L1.Hash, l1GenesisBlockRef.Hash)
	}
	return nil
}

type L2Client interface {
	ChainID(context.Context) (*big.Int, error)
	L2BlockRefByNumber(context.Context, uint64) (eth.L2BlockRef, error)
}

// CheckL2ChainID checks that the configured L2 chain ID matches the client's chain ID.
func (cfg *Config) CheckL2ChainID(ctx context.Context, client L2Client) error {
	id, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	if cfg.L2ChainID.Cmp(id) != 0 {
		return fmt.Errorf("incorrect L2 RPC chain id, expected from config %d, obtained from client %d", cfg.L2ChainID, id)
	}
	return nil
}

// CheckL2GenesisBlockHash checks that the configured L2 genesis block hash is valid for the given client.
func (cfg *Config) CheckL2GenesisBlockHash(ctx context.Context, client L2Client) error {
	l2GenesisBlockRef, err := client.L2BlockRefByNumber(ctx, cfg.Genesis.L2.Number)
	if err != nil {
		return err
	}
	if l2GenesisBlockRef.Hash != cfg.Genesis.L2.Hash {
		return fmt.Errorf("incorrect L2 genesis block hash %d, expected %d", cfg.Genesis.L2.Hash, l2GenesisBlockRef.Hash)
	}
	return nil
}

// Check verifies that the given configuration makes sense
func (cfg *Config) Check() error {
	if cfg.BlockTime == 0 {
		return ErrBlockTimeZero
	}
	if cfg.ChannelTimeout == 0 {
		return ErrMissingChannelTimeout
	}
	if cfg.ProposerWindowSize < 2 {
		return ErrInvalidProposerWindowSize
	}
	if cfg.Genesis.L1.Hash == (common.Hash{}) {
		return ErrMissingGenesisL1Hash
	}
	if cfg.Genesis.L2.Hash == (common.Hash{}) {
		return ErrMissingGenesisL2Hash
	}
	if cfg.Genesis.L2.Hash == cfg.Genesis.L1.Hash {
		return ErrGenesisHashesSame
	}
	if cfg.Genesis.L2Time == 0 {
		return ErrMissingGenesisL2Time
	}
	if cfg.Genesis.SystemConfig.BatcherAddr == (common.Address{}) {
		return ErrMissingBatcherAddr
	}
	if cfg.Genesis.SystemConfig.Overhead == (eth.Bytes32{}) {
		return ErrMissingOverhead
	}
	if cfg.Genesis.SystemConfig.Scalar == (eth.Bytes32{}) {
		return ErrMissingScalar
	}
	if cfg.Genesis.SystemConfig.GasLimit == 0 {
		return ErrMissingGasLimit
	}
	if cfg.BatchInboxAddress == (common.Address{}) {
		return ErrMissingBatchInboxAddress
	}
	if cfg.DepositContractAddress == (common.Address{}) {
		return ErrMissingDepositContractAddress
	}
	if cfg.L1ChainID == nil {
		return ErrMissingL1ChainID
	}
	if cfg.L2ChainID == nil {
		return ErrMissingL2ChainID
	}
	if cfg.L1ChainID.Cmp(cfg.L2ChainID) == 0 {
		return ErrChainIDsSame
	}
	if cfg.L1ChainID.Sign() < 1 {
		return ErrL1ChainIDNotPositive
	}
	if cfg.L2ChainID.Sign() < 1 {
		return ErrL2ChainIDNotPositive
	}
	return nil
}

func (c *Config) L1Signer() types.Signer {
	return types.NewLondonSigner(c.L1ChainID)
}

// Description outputs a banner describing the important parts of rollup configuration in a human-readable form.
// Optionally provide a mapping of L2 chain IDs to network names to label the L2 chain with if not unknown.
// The config should be config.Check()-ed before creating a description.
func (c *Config) Description(l2Chains map[string]string) string {
	// Find and report the network the user is running
	var banner string
	networkL2 := ""
	if l2Chains != nil {
		networkL2 = l2Chains[c.L2ChainID.String()]
	}
	if networkL2 == "" {
		networkL2 = "unknown L2"
	}
	networkL1 := params.NetworkNames[c.L1ChainID.String()]
	if networkL1 == "" {
		networkL1 = "unknown L1"
	}
	banner += fmt.Sprintf("L2 Chain ID: %v (%s)\n", c.L2ChainID, networkL2)
	banner += fmt.Sprintf("L1 Chain ID: %v (%s)\n", c.L1ChainID, networkL1)
	// Report the genesis configuration
	banner += "Kanvas starting point:\n"
	banner += fmt.Sprintf("  L2 starting time: %d ~ %s\n", c.Genesis.L2Time, fmtTime(c.Genesis.L2Time))
	banner += fmt.Sprintf("  L2 block: %s %d\n", c.Genesis.L2.Hash, c.Genesis.L2.Number)
	banner += fmt.Sprintf("  L1 block: %s %d\n", c.Genesis.L1.Hash, c.Genesis.L1.Number)
	return banner
}

// Description outputs a banner describing the important parts of rollup configuration in a log format.
// Optionally provide a mapping of L2 chain IDs to network names to label the L2 chain with if not unknown.
// The config should be config.Check()-ed before creating a description.
func (c *Config) LogDescription(log log.Logger, l2Chains map[string]string) {
	// Find and report the network the user is running
	networkL2 := ""
	if l2Chains != nil {
		networkL2 = l2Chains[c.L2ChainID.String()]
	}
	if networkL2 == "" {
		networkL2 = "unknown L2"
	}
	networkL1 := params.NetworkNames[c.L1ChainID.String()]
	if networkL1 == "" {
		networkL1 = "unknown L1"
	}
	log.Info("Rollup Config", "l2_chain_id", c.L2ChainID, "l2_network", networkL2, "l1_chain_id", c.L1ChainID,
		"l1_network", networkL1, "l2_start_time", c.Genesis.L2Time, "l2_block_hash", c.Genesis.L2.Hash.String(),
		"l2_block_number", c.Genesis.L2.Number, "l1_block_hash", c.Genesis.L1.Hash.String(),
		"l1_block_number", c.Genesis.L1.Number)
}

func fmtForkTimeOrUnset(v *uint64) string {
	if v == nil {
		return "(not configured)"
	}
	if *v == 0 { // don't output the unix epoch time if it's really just activated at genesis.
		return "@ genesis"
	}
	return fmt.Sprintf("@ %-10v ~ %s", *v, fmtTime(*v))
}

func fmtTime(v uint64) string {
	return time.Unix(int64(v), 0).Format(time.UnixDate)
}

type Epoch uint64
