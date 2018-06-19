package model

//玩家数值
const (
	HeroType            = iota
	HP
	ATK
	AttackFrame
	SkillPoint
	IntervalFrame
	MaxHP
)

//阵营
const (
	Atts = 0
	Defs = 1
)

//施法动作
const (
	NonAction = iota
	Star
	Attack
)


//伤害=((A基础攻击力x(1+A攻击加成)+A额外攻击力)x(1+A伤害加成)+A技能伤害)x(1+A暴击伤害加成)

const (
	BaseAttack     = iota
	ATKAddition
	ExtATK
	DamageAddition
	SkillDamage
	SuckBlood
)

const (
	Non= iota
	Jin
	Mu
	Shui
	Huo
	Tu
)