package ucheniki

type StatusType int64

const (
	StatusTypeDefence = StatusType(0)
)

type Status struct {
	target     *Unit
	statusType StatusType
	value      int64
	endTurn    int64
}
