package price

type RunScaleCmd struct {
	Items int     `form:"items"`
	Size  float64 `form:"size"`
	Cost  float64 `form:"cost"`
}

type UnitPrice struct {
	Amount float64
}

func HandleRunScale(cmd *RunScaleCmd) (UnitPrice, error) {
	unitPrice := cmd.Cost / (float64(cmd.Items) * cmd.Size)

	return UnitPrice{Amount: unitPrice}, nil
}
