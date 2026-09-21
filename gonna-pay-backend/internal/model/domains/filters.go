package domains

type CostPeriod string

const (
	CostPeriodMonth CostPeriod = "month"
	CostPeriodWeek  CostPeriod = "7days"
)

type CostType string

const (
	CostTypeSolo  CostType = "solo"
	CostTypeGroup CostType = "group"
)

type CostFilters struct {
	Category string
	Period   CostPeriod
	Type     CostType
	MinValue *float64
	MaxValue *float64
}

type ContactFilters struct {
	Category string
	Search   string
}

type GroupFilters struct {
	Category string
	Search   string
}
