package msig

import (
	"github.com/cochainio/eos-go"
)

func init() {
	eos.RegisterAction(eos.AN("eosio.msig"), eos.ActN("propose"), &Propose{})
	eos.RegisterAction(eos.AN("eosio.msig"), eos.ActN("approve"), &Approve{})
	eos.RegisterAction(eos.AN("eosio.msig"), eos.ActN("unapprove"), &Unapprove{})
	eos.RegisterAction(eos.AN("eosio.msig"), eos.ActN("cancel"), &Cancel{})
	eos.RegisterAction(eos.AN("eosio.msig"), eos.ActN("exec"), &Exec{})
}
