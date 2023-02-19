package ucheniki

type Unit struct {
	ai    AI
	team  int64
	hp    int64
	maxHp int64
	name  string

	id     int64
	skills []*Skill
}

func NewUnit(team int64, name string) *Unit {
	return &Unit{ai: &SingleUnitAI{}, team: team, hp: 100, maxHp: 100, skills: make([]*Skill, 0), name: name}
}

func (u *Unit) WithSkills(skills ...*Skill) *Unit {
	u.skills = append(u.skills, skills...)
	return u
}

func (u *Unit) GetOwnDefenceValue() int64 {
	value := int64(0)
	for i := range u.skills {
		if u.skills[i].typ == SkillTypePassiveDefence {
			value += u.skills[i].power
		}
	}
	return value
}

func (u *Unit) AI() AI {
	return u.ai
}

func (u *Unit) Team() int64 {
	return u.team
}

func (u *Unit) Alive() bool {
	return u.hp > 0
}

func (u *Unit) ID() int64 {
	return u.id
}

func (u *Unit) SetID(id int64) {
	u.id = id
}

func (u *Unit) Skill(skillType SkillType) *Skill {
	for i := range u.skills {
		if u.skills[i].typ == skillType {
			return u.skills[i]
		}
	}
	return nil
}

func (u *Unit) AddSkill(s *Skill) {
	u.skills = append(u.skills, s)
}
