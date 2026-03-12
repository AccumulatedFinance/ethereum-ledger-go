package ledger

import (
	"github.com/AccumulatedFinance/ethereum-ledger-go/accounts"
	"github.com/AccumulatedFinance/ethereum-ledger-go/usbwallet"
)

type EthereumLedger struct {
	hub *usbwallet.Hub
}

func (el EthereumLedger) Wallets() []accounts.Wallet {
	return el.hub.Wallets()
}

func (el *EthereumLedger) Close() error {
	if el.hub != nil {
		return el.hub.Close()
	}
	return nil
}

func New() (*EthereumLedger, error) {
	l := &EthereumLedger{}
	hub, err := usbwallet.NewLedgerHub()
	if err != nil {
		return &EthereumLedger{}, err
	}

	l.hub = hub
	return l, nil
}
