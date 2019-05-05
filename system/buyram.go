package system

import (
	eos "github.com/eoscochain/eos-go"
)

func NewBuyRAMQuantity(payer, receiver eos.AccountName, eosQuantity uint64) *eos.Action {
	return NewBuyRAM(payer, receiver, eos.NewEOSAsset(int64(eosQuantity)))
}

func NewBuyRAM(payer, receiver eos.AccountName, asset eos.Asset) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("buyram"),
		Authorization: []eos.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: asset,
		}),
	}
	return a
}

// BuyRAM represents the `eosio.system::buyram` action.
type BuyRAM struct {
	Payer    eos.AccountName `json:"payer"`
	Receiver eos.AccountName `json:"receiver"`
	Quantity eos.Asset       `json:"quant"` // specified in EOS
}
