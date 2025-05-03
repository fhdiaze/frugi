package commands

import (
	"time"

	"github.com/fhdiaze/frugi/internal/pkg/types"
)

type RunConvertCmd struct {
	HourlyWage float64 `form:"hourly-wage"`
	Amount     float64 `form:"amount"`
}

type RunCompoundCmd struct {
	Principal types.Money `form:"principal"`
}

func HandleRunConvert(cmd *RunConvertCmd) (time.Duration, error) {
	amount := types.MoneyFromMajor(cmd.Amount)
	hours := time.Duration(amount.Div(cmd.HourlyWage).ToFloat64())

	return time.Hour * hours, nil
}

func HandleRunCompound(cmd *RunCompoundCmd) (types.Money, error) {
	result := types.MoneyFromMajor(0)

	return result, nil
}
