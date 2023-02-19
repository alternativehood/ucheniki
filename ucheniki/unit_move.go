package ucheniki

import (
	"math"
	"math/rand"
)

type UnitMove struct {
	skill  *Skill
	target *Unit
}

func (m *UnitMove) Execute(b *Battle) {
	if m.skill.chance < 1.0 {
		if rand.Float64() > m.skill.chance {
			return
		}
	}

	if m.skill.typ == SkillTypeAttack {
		damage := float64(m.skill.power)
		targetDefence := b.GetUnitDefenceValue(m.target.ID())
		damage *= 1 - targetDefence/100.0
		m.target.hp = int64(math.Max(float64(m.target.hp)-damage, 0))
		return
	}

	if m.skill.typ == SkillTypeDefendSelf {
		b.AddStatus(&Status{value: 20, target: m.target, statusType: StatusTypeDefence, endTurn: b.turn + m.Cooldown()})
		return
	}

	if m.skill.typ == SkillTypeHeal {
		m.target.hp = int64(math.Min(float64(m.target.maxHp), float64(m.target.hp+m.skill.power)))
		return
	}
}

func (m *UnitMove) Cooldown() int64 {
	return m.skill.cooldown
}
