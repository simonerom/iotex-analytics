// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package mimo

import (
	"fmt"
	"io"
	"strconv"
)

type Action struct {
	Hash        string     `json:"hash"`
	Type        ActionType `json:"type"`
	Exchange    string     `json:"exchange"`
	Account     string     `json:"account"`
	IotxAmount  string     `json:"iotxAmount"`
	TokenAmount string     `json:"tokenAmount"`
	Time        string     `json:"time"`
}

type AmountInOneDay struct {
	Amount string `json:"amount"`
	Date   string `json:"date"`
}

type Exchange struct {
	Address                  string `json:"address"`
	Token                    Token  `json:"token"`
	Supply                   string `json:"supply"`
	VolumeInPast24Hours      string `json:"volumeInPast24Hours"`
	VolumeInPast7Days        string `json:"volumeInPast7Days"`
	BalanceOfToken           string `json:"balanceOfToken"`
	BalanceOfToken24HoursAgo string `json:"balanceOfToken24HoursAgo"`
	BalanceOfIotx            string `json:"balanceOfIOTX"`
	BalanceOfIOTX24HoursAgo  string `json:"balanceOfIOTX24HoursAgo"`
}

type Pagination struct {
	Skip  int `json:"skip"`
	First int `json:"first"`
}

type Stats struct {
	NumOfTransations int    `json:"numOfTransations"`
	Volume           string `json:"volume"`
}

type Token struct {
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type ActionType string

const (
	ActionTypeAll    ActionType = "ALL"
	ActionTypeSwap   ActionType = "SWAP"
	ActionTypeAdd    ActionType = "ADD"
	ActionTypeRemove ActionType = "REMOVE"
)

var AllActionType = []ActionType{
	ActionTypeAll,
	ActionTypeSwap,
	ActionTypeAdd,
	ActionTypeRemove,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypeAll, ActionTypeSwap, ActionTypeAdd, ActionTypeRemove:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
