package system

import "github.com/cochainio/eos-go"

func NewSetSymbol(symbol string) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setsymbol"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetSymbol{
			Symbol: symbol,
		}),
	}
}

type SetSymbol struct {
	Symbol string `json:"symbol"`
}
