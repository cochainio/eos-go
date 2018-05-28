package system

import (
	"github.com/cochainio/eos-go"
	"github.com/cochainio/eos-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract.
func NewNewAccount(creator, newAccount eos.AccountName, ownerKey ecc.PublicKey, activeKey ...ecc.PublicKey) *eos.Action {
	if len(activeKey) == 0 {
		activeKey = []ecc.PublicKey{ownerKey}
	}
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("newaccount"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: ownerKey,
						Weight:    1,
					},
				},
				Accounts: []eos.PermissionLevelWeight{},
			},
			Active: eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: activeKey[0],
						Weight:    1,
					},
				},
				Accounts: []eos.PermissionLevelWeight{},
			},
		}),
	}
}

// NewAccount represents a `newaccount` action on the `eosio.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator  eos.AccountName `json:"creator"`
	Name     eos.AccountName `json:"name"`
	Owner    eos.Authority   `json:"owner"`
	Active   eos.Authority   `json:"active"`
}
