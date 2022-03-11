package test

import (
	"context"
	"errors"
	"time"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/wtxmgr"
	"github.com/lightninglabs/lndclient"
	"github.com/lightningnetwork/lnd/keychain"
	"github.com/lightningnetwork/lnd/lnwallet"
	"github.com/lightningnetwork/lnd/lnwallet/chainfee"
)

type mockWalletKit struct {
	lndclient.WalletKitClient
	lnd          *LndMockServices
	keyIndex     int32
	feeEstimates map[int32]chainfee.SatPerKWeight
}

var _ lndclient.WalletKitClient = (*mockWalletKit)(nil)

func (m *mockWalletKit) ListUnspent(ctx context.Context, minConfs,
	maxConfs int32) ([]*lnwallet.Utxo, error) {

	return nil, nil
}

func (m *mockWalletKit) LeaseOutput(ctx context.Context, lockID wtxmgr.LockID,
	op wire.OutPoint, leaseTime time.Duration) (time.Time, error) {

	return time.Now(), nil
}

func (m *mockWalletKit) ReleaseOutput(ctx context.Context,
	lockID wtxmgr.LockID, op wire.OutPoint) error {

	return nil
}

func (m *mockWalletKit) DeriveNextKey(ctx context.Context, family int32) (
	*keychain.KeyDescriptor, error) {

	index := m.keyIndex

	_, pubKey := CreateKey(index)
	m.keyIndex++

	return &keychain.KeyDescriptor{
		KeyLocator: keychain.KeyLocator{
			Family: keychain.KeyFamily(family),
			Index:  uint32(index),
		},
		PubKey: pubKey,
	}, nil
}

func (m *mockWalletKit) DeriveKey(ctx context.Context, in *keychain.KeyLocator) (
	*keychain.KeyDescriptor, error) {

	_, pubKey := CreateKey(int32(in.Index))

	return &keychain.KeyDescriptor{
		KeyLocator: *in,
		PubKey:     pubKey,
	}, nil
}

func (m *mockWalletKit) NextAddr(ctx context.Context) (btcutil.Address, error) {
	addr, err := btcutil.NewAddressWitnessPubKeyHash(
		make([]byte, 20), &chaincfg.TestNet3Params,
	)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (m *mockWalletKit) PublishTransaction(_ context.Context, tx *wire.MsgTx,
	_ string) error {

	m.lnd.AddTx(tx)
	m.lnd.TxPublishChannel <- tx
	return nil
}

func (m *mockWalletKit) SendOutputs(ctx context.Context, outputs []*wire.TxOut,
	feeRate chainfee.SatPerKWeight, _ string) (*wire.MsgTx, error) {

	var inputTxHash chainhash.Hash

	tx := wire.MsgTx{}
	tx.AddTxIn(&wire.TxIn{
		PreviousOutPoint: wire.OutPoint{
			Hash:  inputTxHash,
			Index: 0,
		},
	})

	for _, out := range outputs {
		tx.AddTxOut(&wire.TxOut{
			PkScript: out.PkScript,
			Value:    out.Value,
		})
	}

	m.lnd.AddTx(&tx)
	m.lnd.SendOutputsChannel <- tx

	return &tx, nil
}

func (m *mockWalletKit) EstimateFee(ctx context.Context, confTarget int32) (
	chainfee.SatPerKWeight, error) {

	if confTarget <= 1 {
		return 0, errors.New("conf target must be greater than 1")
	}

	feeEstimate, ok := m.feeEstimates[confTarget]
	if !ok {
		return 10000, nil
	}

	return feeEstimate, nil
}

// ListSweeps returns a list of the sweep transaction ids known to our node.
func (m *mockWalletKit) ListSweeps(_ context.Context) ([]string, error) {
	return m.lnd.Sweeps, nil
}
