package system

import (
	"github.com/cochainio/eos-go"
	"strconv"
)

func NewSetGlobal(name, value string) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setglobal"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetGlobal{
			Name: name,
			Value: value,
		}),
	}
}

type SetGlobal struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

func NewSetProducerScheduleSize(schedSize uint8) *eos.Action {
	return NewSetGlobal("max_producer_schedule_size", strconv.FormatUint(uint64(schedSize), 10))
}

func NewSetPervoteDailyPay(vpay int64) *eos.Action {
	return NewSetGlobal("min_pervote_daily_pay", strconv.FormatInt(vpay, 10))
}

func NewSetMinActivatedStake(minStake int64) *eos.Action {
	return NewSetGlobal("min_activated_stake", strconv.FormatInt(minStake, 10))
}

func NewSetContinuousRate(rate float64) *eos.Action {
	return NewSetGlobal("continuous_rate", strconv.FormatFloat(rate, 'f', -1, 64))
}

func NewSetToProducersRate(rate float64) *eos.Action {
	return NewSetGlobal("to_producers_rate", strconv.FormatFloat(rate, 'f', -1, 64))
}

func NewSetToBpayRate(rate float64) *eos.Action {
	return NewSetGlobal("to_bpay_rate", strconv.FormatFloat(rate, 'f', -1, 64))
}
