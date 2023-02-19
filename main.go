package main

import (
	"ucheniki/ucheniki"

	"github.com/kelindar/tile"
)

func RunSimulation() {
	b := ucheniki.NewBattle()
	b.AddUnit(
		ucheniki.NewUnit(1, "warrior").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeAttack, 10, 10, 0.8)),
		tile.Point{X: 0, Y: 0},
	)
	b.AddUnit(
		ucheniki.NewUnit(2, "warrior").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeAttack, 10, 10, 0.8)),
		tile.Point{X: 0, Y: 1},
	)
	b.AddUnit(
		ucheniki.NewUnit(1, "warrior").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeAttack, 10, 10, 0.8)),
		tile.Point{X: 1, Y: 1},
	)
	b.AddUnit(
		ucheniki.NewUnit(2, "warrior").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeAttack, 10, 10, 0.8)),
		tile.Point{X: 1, Y: 0},
	)
	b.AddUnit(
		ucheniki.NewUnit(1, "healer").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeHeal, 5, 10, 1.0)),
		tile.Point{X: 0, Y: 2},
	)
	b.AddUnit(
		ucheniki.NewUnit(2, "healer").WithSkills(ucheniki.NewSkill(ucheniki.SkillTypeHeal, 5, 20, 1.0)),
		tile.Point{X: 1, Y: 2},
	)

	for {
		if finished, winner := b.Finished(); finished {
			print("winner:", winner)
			return
		}
		b.Step()
	}
}

func main() {
	RunSimulation()
}
