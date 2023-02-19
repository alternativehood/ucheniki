package ucheniki

import (
	"fmt"
	"math"

	"github.com/kelindar/tile"
)

type Battle struct {
	units     []*Unit
	unitTurns map[int64]int64
	grid      *tile.Grid

	turn     int64
	statuses []*Status
}

func NewBattle() *Battle {
	return &Battle{
		grid:      tile.NewGrid(6, 3),
		units:     make([]*Unit, 0),
		statuses:  make([]*Status, 0),
		unitTurns: make(map[int64]int64, 0),
	}
}

func (b *Battle) AddStatus(s *Status) {
	b.statuses = append(b.statuses, s)
}

func (b *Battle) Step() {
	b.prepareForStep()
	currentUnit := b.getCurrentUnit()
	b.turn = b.unitTurns[currentUnit.ID()]
	unitMove := currentUnit.AI().GetMove(b, currentUnit)
	if unitMove != nil {
		unitMove.Execute(b)
		b.unitTurns[currentUnit.ID()] += unitMove.Cooldown()
	} else {
		b.unitTurns[currentUnit.ID()] += idleCooldown
	}
	b.cleanupStep()
}

func (b *Battle) AddUnit(u *Unit, position tile.Point) {
	b.units = append(b.units, u)
	unitID := int64(len(b.units))
	u.SetID(unitID)
	b.grid.WriteAt(position.X, position.Y, tile.Tile{byte(unitID)})
	b.unitTurns[unitID] = unitID
}

func (b *Battle) Finished() (finished bool, winner int64) {
	teamsAlive := make(map[int64]struct{})
	for i := range b.units {
		if !b.units[i].Alive() {
			continue
		}
		teamsAlive[b.units[i].Team()] = struct{}{}
		winner = b.units[i].Team()
	}
	if len(teamsAlive) > 1 {
		return false, 0
	}
	return true, winner
}

func (b *Battle) cleanupStep() {
	fmt.Println(fmt.Sprintf("---step %d results---", b.turn))
	for i := range b.units {
		print(b.units[i].name, " ", b.units[i].id, " ", b.units[i].Team(), " ", b.units[i].hp, "\n")
	}
}

func (b *Battle) prepareForStep() {}

func (b *Battle) getCurrentUnit() *Unit {
	resultID := int64(0)
	for uID := range b.unitTurns {
		if !b.GetUnitByID(uID).Alive() {
			continue
		}

		if b.unitTurns[resultID] < b.turn && b.unitTurns[uID] > b.unitTurns[resultID] || resultID == 0 {
			resultID = uID
			continue
		}

		if b.unitTurns[uID] < b.unitTurns[resultID] {
			resultID = uID
		}
	}
	return b.GetUnitByID(resultID)
}

func (b *Battle) GetUnitByID(id int64) *Unit {
	return b.units[id-1]
}

func (b *Battle) GetUnitDefenceValue(id int64) float64 {
	value := b.GetUnitByID(id).GetOwnDefenceValue()
	for i := range b.statuses {
		s := b.statuses[i]
		if s.target.ID() == id && s.statusType == StatusTypeDefence {
			value += s.value
		}
	}
	return math.Min(float64(value), 95.0)
}
