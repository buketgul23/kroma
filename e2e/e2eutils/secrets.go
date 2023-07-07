package e2eutils

import (
	"crypto/ecdsa"
	"fmt"

	hdwallet "github.com/ethereum-optimism/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// DefaultMnemonicConfig is the default mnemonic used in testing.
// We prefer a mnemonic rather than direct private keys to make it easier
// to export all testing keys in external tooling for use during debugging.
var DefaultMnemonicConfig = &MnemonicConfig{
	Mnemonic:         "test test test test test test test test test test test junk",
	Deployer:         "m/44'/60'/0'/0/1",
	CliqueSigner:     "m/44'/60'/0'/0/2",
	TrustedValidator: "m/44'/60'/0'/0/3",
	Challenger:       "m/44'/60'/0'/0/4",
	Batcher:          "m/44'/60'/0'/0/5",
	ProposerP2P:      "m/44'/60'/0'/0/6",
	Alice:            "m/44'/60'/0'/0/7",
	Bob:              "m/44'/60'/0'/0/8",
	Mallory:          "m/44'/60'/0'/0/9",
	SysCfgOwner:      "m/44'/60'/0'/0/10",
}

// MnemonicConfig configures the private keys for the hive testnet.
// It's json-serializable, so we can ship it to e.g. the hardhat script client.
type MnemonicConfig struct {
	Mnemonic string

	Deployer     string
	CliqueSigner string
	SysCfgOwner  string

	// rollup actors
	TrustedValidator string
	Challenger       string
	Batcher          string
	ProposerP2P      string

	// prefunded L1/L2 accounts for testing
	Alice   string
	Bob     string
	Mallory string
}

// Secrets computes the private keys for all mnemonic paths,
// which can then be kept around for fast precomputed private key access.
func (m *MnemonicConfig) Secrets() (*Secrets, error) {
	wallet, err := hdwallet.NewFromMnemonic(m.Mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}
	account := func(path string) accounts.Account {
		return accounts.Account{URL: accounts.URL{Path: path}}
	}

	deployer, err := wallet.PrivateKey(account(m.Deployer))
	if err != nil {
		return nil, err
	}
	cliqueSigner, err := wallet.PrivateKey(account(m.CliqueSigner))
	if err != nil {
		return nil, err
	}
	sysCfgOwner, err := wallet.PrivateKey(account(m.SysCfgOwner))
	if err != nil {
		return nil, err
	}
	trustedValidator, err := wallet.PrivateKey(account(m.TrustedValidator))
	if err != nil {
		return nil, err
	}
	challenger, err := wallet.PrivateKey(account(m.Challenger))
	if err != nil {
		return nil, err
	}
	batcher, err := wallet.PrivateKey(account(m.Batcher))
	if err != nil {
		return nil, err
	}
	proposerP2P, err := wallet.PrivateKey(account(m.ProposerP2P))
	if err != nil {
		return nil, err
	}
	alice, err := wallet.PrivateKey(account(m.Alice))
	if err != nil {
		return nil, err
	}
	bob, err := wallet.PrivateKey(account(m.Bob))
	if err != nil {
		return nil, err
	}
	mallory, err := wallet.PrivateKey(account(m.Mallory))
	if err != nil {
		return nil, err
	}

	return &Secrets{
		Deployer:         deployer,
		CliqueSigner:     cliqueSigner,
		SysCfgOwner:      sysCfgOwner,
		TrustedValidator: trustedValidator,
		Challenger:       challenger,
		Batcher:          batcher,
		ProposerP2P:      proposerP2P,
		Alice:            alice,
		Bob:              bob,
		Mallory:          mallory,
		Wallet:           wallet,
	}, nil
}

// Secrets bundles secp256k1 private keys for all common rollup actors for testing purposes.
type Secrets struct {
	Deployer     *ecdsa.PrivateKey
	CliqueSigner *ecdsa.PrivateKey
	SysCfgOwner  *ecdsa.PrivateKey

	// rollup actors
	TrustedValidator *ecdsa.PrivateKey
	Challenger       *ecdsa.PrivateKey
	Batcher          *ecdsa.PrivateKey
	ProposerP2P      *ecdsa.PrivateKey

	// prefunded L1/L2 accounts for testing
	Alice   *ecdsa.PrivateKey
	Bob     *ecdsa.PrivateKey
	Mallory *ecdsa.PrivateKey

	// Share the wallet to be able to generate more accounts
	Wallet *hdwallet.Wallet
}

// EncodePrivKey encodes the given private key in 32 bytes
func EncodePrivKey(priv *ecdsa.PrivateKey) hexutil.Bytes {
	privkey := make([]byte, 32)
	blob := priv.D.Bytes()
	copy(privkey[32-len(blob):], blob)
	return privkey
}

// Addresses computes the ethereum address of each account,
// which can then be kept around for fast precomputed address access.
func (s *Secrets) Addresses() *Addresses {
	return &Addresses{
		Deployer:         crypto.PubkeyToAddress(s.Deployer.PublicKey),
		CliqueSigner:     crypto.PubkeyToAddress(s.CliqueSigner.PublicKey),
		SysCfgOwner:      crypto.PubkeyToAddress(s.SysCfgOwner.PublicKey),
		TrustedValidator: crypto.PubkeyToAddress(s.TrustedValidator.PublicKey),
		Challenger:       crypto.PubkeyToAddress(s.Challenger.PublicKey),
		Batcher:          crypto.PubkeyToAddress(s.Batcher.PublicKey),
		ProposerP2P:      crypto.PubkeyToAddress(s.ProposerP2P.PublicKey),
		Alice:            crypto.PubkeyToAddress(s.Alice.PublicKey),
		Bob:              crypto.PubkeyToAddress(s.Bob.PublicKey),
		Mallory:          crypto.PubkeyToAddress(s.Mallory.PublicKey),
	}
}

// Addresses bundles the addresses for all common rollup addresses for testing purposes.
type Addresses struct {
	Deployer     common.Address
	CliqueSigner common.Address
	SysCfgOwner  common.Address

	// rollup actors
	TrustedValidator common.Address
	Challenger       common.Address
	Batcher          common.Address
	ProposerP2P      common.Address

	// prefunded L1/L2 accounts for testing
	Alice   common.Address
	Bob     common.Address
	Mallory common.Address
}

func (a *Addresses) All() []common.Address {
	return []common.Address{
		a.Deployer,
		a.CliqueSigner,
		a.SysCfgOwner,
		a.TrustedValidator,
		a.Challenger,
		a.Batcher,
		a.ProposerP2P,
		a.Alice,
		a.Bob,
		a.Mallory,
	}
}
