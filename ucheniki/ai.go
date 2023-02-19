package ucheniki

import (
	"github.com/kelindar/tile"
)

type AI interface {
	GetMove(b *Battle, u *Unit) *UnitMove
}

type SingleUnitAI struct {
}

func (s *SingleUnitAI) getHealTarget(b *Battle, u *Unit) *Unit {
	if u.hp < u.maxHp {
		return u
	}

	var target *Unit
	for i := range b.units {
		t := b.units[i]
		if t.team != u.team {
			continue
		}
		if !t.Alive() {
			continue
		}
		if t.hp < t.maxHp {
			if target != nil {
				if t.hp/t.maxHp < target.hp/target.maxHp {
					target = t
					continue
				}
			}
			target = t
		}
	}
	return target
}

func (s *SingleUnitAI) getAttackTarget(b *Battle, u *Unit) *Unit {
	var position tile.Point
	b.grid.Each(func(point tile.Point, tile tile.Tile) {
		if tile[0] != byte(u.ID()) {
			return
		}
		position = point
	})

	targets := make([]*Unit, 0, 3)
	b.grid.Neighbors(position.X, position.Y, func(point tile.Point, t tile.Tile) {
		if t[0] == byte(u.ID()) {
			return
		}
		if t[0] == EmptyTile {
			return
		}
		unitAtPos := b.GetUnitByID(int64(t[0]))
		if !unitAtPos.Alive() || unitAtPos.team == u.Team() {
			return
		}
		targets = append(targets, unitAtPos)
	})
	if len(targets) > 0 {
		return targets[0]
	}
	return nil
}

func (s *SingleUnitAI) GetMove(b *Battle, u *Unit) *UnitMove {
	attackSkill := u.Skill(SkillTypeAttack)
	if attackSkill != nil {
		target := s.getAttackTarget(b, u)
		if target != nil {
			return &UnitMove{skill: attackSkill, target: target}
		}
	}

	healSkill := u.Skill(SkillTypeHeal)
	if healSkill != nil {
		target := s.getHealTarget(b, u)
		if target != nil {
			return &UnitMove{skill: healSkill, target: target}
		}
	}

	defendSkill := u.Skill(SkillTypeDefendSelf)
	if defendSkill != nil {
		return &UnitMove{skill: defendSkill, target: u}
	}

	return nil
}
