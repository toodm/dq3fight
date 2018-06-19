package model

type FighterSkill struct {
	SkillId int
}

type FighterProperty struct {
	PropertyId int
	Value      int
}

type FigheterBuff struct {
	BuffId       int
	BuffType     int
	EndFrame     int
	BuffProperty int
	Value        int
}

type Fighter struct {
	FrameState *FighterFrame
	Pos        int
	Id         int
	HeroType   int
	HeroId     int
	Skills     []FighterSkill
	Buffs      *[]FigheterBuff
	Propertys  map[int]*FighterProperty
}

func GetHeroAttackInterval(hero int) int {
	return 2
}

//返回 第几个技能，技能ID
func (fighter Fighter) GetSkill(UseSkill bool) (int, int) {
	index := 0
	if UseSkill {
		index = 1
		return 1, AddHP
	} else {
		if fighter.GetProperty(HeroType) == 1 {
			return index, CommonAttack
		} else {
			return index, BallisticAttack
		}
	}

}

func (fighter Fighter) GetSkillIndex(attackType int, skillId int) int {
	if attackType == 0 {
		return 0
	}
	return 1
}

func (fighter *Fighter) Attack(battle *Battle) {
	fighter.FrameState.Attack(fighter, battle)

}
func (fighter *Fighter) GetBuffs(ballistic Ballistic, buffType int) int {

	return 20
}



func (fighter *Fighter) Defend(ballistic Ballistic, battle *Battle) bool {
	//判定技能类型
	//伤害
	//叠技能BUFF
	//治疗
	isDie := false
	if ballistic.GetSkillType() == Damage {
		value := fighter.GetHurt(ballistic)
		fighter.DeCeaseProperty(HP, value)
		suckBlood := fighter.GetBuffs(ballistic, SuckBlood)
		if suckBlood > 0 {
			treatValue := value * suckBlood / 100
			b := Ballistic{
				Form:     ballistic.Form,
				SkillId:  AddHP,
				Value:    treatValue,
				To:       []int{ballistic.Form},
				Buffs:    &map[int][]BuffConfig{},
				EndFrame: battle.Frame + 1,
			}
			battle.AddBallistic(&b)
		}
		battle.Log.HurtSide(battle.Frame, fighter.Id, ballistic.Value, fighter.GetProperty(HP), fighter.FrameState.SkillPoint, ballistic.Form)
		if fighter.GetProperty(HP) <= 0 {
			isDie = true
			battle.remove(fighter.Id)
			println(battle.Frame, ballistic.Form, "  击杀了  ", fighter.Id)
		}
	} else {
		treat := fighter.GetProperty(MaxHP) - fighter.GetProperty(HP)
		if treat > ballistic.Value {
			treat = ballistic.Value
		}
		if treat > 0 {
			fighter.InCeaseProperty(HP, treat)
			battle.Log.TreatSide(battle.Frame, fighter.Id, ballistic.Value, fighter.GetProperty(HP), fighter.FrameState.SkillPoint)
		}
	}

	return isDie
	//释放技能
}

func (fighter Fighter) getHPPercent(ballistic Ballistic) int {
	return fighter.GetProperty(HP) * 100 / fighter.GetProperty(MaxHP)
}

func (fighter Fighter) GetHurt(ballistic Ballistic) int {
	//是否做组合技能伤害判定

	//是否做血量不足伤害判定

	return ballistic.Value
}

func (fighter Fighter) GetAtk(skillId int) int {

	Violent := 0
	//伤害=((A基础攻击力x(1+A攻击加成)+A额外攻击力)x(1+A伤害加成)+A技能伤害)x(1+A暴击伤害加成)x(1-D免疫伤害比例)
	//println(fighter.GetProperty(ATK) ,
	//	fighter.GetBuffValue(ATKAddition),
	//	fighter.GetBuffValue(ExtATK),
	//	fighter.GetBuffValue(DamageAddition),
	//	fighter.GetBuffValue(SkillDamage))
	return (
		(fighter.GetProperty(ATK) *
			(100 + fighter.GetBuffValue(ATKAddition)) / 100 +
			fighter.GetBuffValue(ExtATK)) *
			(100 + fighter.GetBuffValue(DamageAddition)) / 100 +
			fighter.GetBuffValue(SkillDamage)) *
		(100 + Violent) / 100
}

func (fighter Fighter) GetBuffValue(buffType int) int {
	value := 0
	for _, item := range *fighter.Buffs {
		if item.BuffType == buffType {
			value += item.Value
		}
	}
	return value
}

func (fighter *Fighter) SetSilent(curFrame int, endFrame int) {
	fighter.FrameState.SetSilent(curFrame, endFrame)
}

func (fighter *Fighter) SetStop(endFrame int) {
	fighter.FrameState.SetStop(endFrame)
}

func (fighter *Fighter) InCeaseProperty(propType int, propValue int) {
	if val, ok := fighter.Propertys[propType]; ok {
		val.Value += propValue
	}
}

func (fighter *Fighter) DeCeaseProperty(propType int, propValue int) {
	if val, ok := fighter.Propertys[propType]; ok {
		val.Value -= propValue
	}
}

func (fighter *Fighter) GetProperty(propType int) int {
	if val, ok := fighter.Propertys[propType]; ok {
		return val.Value
	}
	return 0
}
