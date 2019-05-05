package system

import (
	"github.com/eoscochain/eos-go"
)

func NewInit(symbol string, precision uint8) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("init"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Init{
			Version: 0,
			Core: eos.Symbol{
				Precision: precision,
				Symbol:    symbol,
			},
		}),
	}
}

type Init struct {
	Version eos.Varuint32 `json:"version"`
	Core    eos.Symbol    `json:"core"`
}
