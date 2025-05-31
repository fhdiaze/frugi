package core

import (
	"fmt"
	"math"

	"github.com/fhdiaze/frugi/internal/pkg/types"
)

type RunCompoundCmd struct {
	Principal    types.Money       `form:"principal"`
	Contribution types.Money       `form:"contribution"`
	InterestRate float64           `form:"interest-rate"`
	Frequency    CompoundFrequency `form:"frequency"`
	Years        int64             `form:"years"`
}

type CompoundFrequency int

func (f *CompoundFrequency) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Monthly":
		*f = Monthly
		return nil
	case "Annually":
		*f = Annually
		return nil
	default:
		return fmt.Errorf("invalid CompoundFrequency: %s", text)
	}
}

const (
	Monthly CompoundFrequency = iota
	Annually
)

var FrequencyNames = map[CompoundFrequency]string{
	Monthly:  "Monthly",
	Annually: "Annually",
}

func AllFrequencyNames() []string {
	var names = make([]string, 0, len(FrequencyNames))
	for _, name := range FrequencyNames {
		names = append(names, name)
	}

	return names
}

var annualCompounds = map[CompoundFrequency]int{
	Monthly:  12,
	Annually: 1,
}

func HandleRunCompound(cmd *RunCompoundCmd) types.Money {
	compounds := annualCompounds[cmd.Frequency]
	periodicInterestRate := cmd.InterestRate / 100.0 / float64(compounds)
	accumulation := math.Pow(1+periodicInterestRate, float64(cmd.Years)*float64(compounds))
	composedPrincipal := *cmd.Principal.Mul(accumulation)
	composedContribution := *cmd.Contribution.Mul(accumulation - 1.0).Div(periodicInterestRate)

	return *composedPrincipal.Add(&composedContribution)
}
