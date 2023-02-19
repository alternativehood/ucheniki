package ucheniki

type SkillType int64

const (
	SkillTypeAttack         = SkillType(0)
	SkillTypeDefendSelf     = SkillType(1)
	SkillTypePassiveDefence = SkillType(2)
	SkillTypeHeal           = SkillType(3)
)

type Skill struct {
	typ      SkillType
	power    int64
	chance   float64
	cooldown int64
}

func NewSkill(t SkillType, power, cooldown int64, chance float64) *Skill {
	return &Skill{typ: t, power: power, cooldown: cooldown, chance: chance}
}
