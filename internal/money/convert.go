package money

import (
	"time"
)

type RunConvertCmd struct {
	HourlyWage float64 `form:"hourly-wage"`
	Amount     float64 `form:"amount"`
}

func HandleRunConvert(cmd *RunConvertCmd) (time.Duration, error) {
	amount := FromMajor(cmd.Amount)
	wage := FromMajor(cmd.HourlyWage)
	hours := time.Duration(amount / wage)

	return time.Hour * hours, nil
}
