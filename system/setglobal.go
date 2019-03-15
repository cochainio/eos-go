package system

import (
	"strconv"

	"github.com/cochainio/eos-go"
)

func NewSetGlobal(name, value string) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setglobal"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetGlobal{
			Name:  name,
			Value: value,
		}),
	}
}

type SetGlobal struct {
	Name  string `json:"name"`
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

func NewSetUsecondsPerDay(usecondsPerDay int64) *eos.Action {
	return NewSetGlobal("useconds_per_day", strconv.FormatInt(usecondsPerDay, 10))
}

func NewSetContinuousRate(rate float64) *eos.Action {
	rate6 := int64(rate * 1e6)
	return NewSetGlobal("continuous_rate", strconv.FormatInt(rate6, 10))
}

func NewSetToProducersRate(rate float64) *eos.Action {
	rate6 := int64(rate * 1e6)
	return NewSetGlobal("to_producers_rate", strconv.FormatInt(rate6, 10))
}

func NewSetToBpayRate(rate float64) *eos.Action {
	rate6 := int64(rate * 1e6)
	return NewSetGlobal("to_bpay_rate", strconv.FormatInt(rate6, 10))
}

func NewSetToVoterBonusRate(rate float64) *eos.Action {
	rate6 := int64(rate * 1e6)
	return NewSetGlobal("to_voter_bonus_rate", strconv.FormatInt(rate6, 10))
}
