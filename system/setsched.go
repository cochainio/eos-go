package system

import "github.com/cochainio/eos-go"

func NewSetSched(schedSize uint8) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setsched"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetSched{
			SchedSize: schedSize,
		}),
	}
}

type SetSched struct {
	SchedSize uint8 `json:"sched_size"`
}
