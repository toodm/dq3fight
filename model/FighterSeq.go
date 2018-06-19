package model

type FighterSeq struct {
	BallisticsMap map[int][]Ballistic
}

func (seq *FighterSeq) AddSeq(frame int, seqType int, ballistic Ballistic) {
	if val, ok := seq.BallisticsMap[frame]; ok {
		val = append(val, ballistic)
	} else {
		seq.BallisticsMap[frame] = []Ballistic{}
	}
}

func (fighter *Fighter) Get1Hurt(config []BuffConfig){

	return
}

func (ballistic Ballistic) GetAtk()int{

	return ballistic.Value
}
func (ballistic Ballistic) GetBuffs(buffType int)int{

	return 10
}
func (ballistic Ballistic) HurtAfter(damage int,battle *Battle)int{

	suckBlood := ballistic.GetBuffs(SuckBlood)
	if suckBlood > 0 {
		treatValue := damage * suckBlood / 100
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
	return ballistic.Value
}


func (fighter *Fighter) ExecuteSeq(seq *FighterSeq,battle *Battle) {
	if val, ok := seq.BallisticsMap[battle.Frame]; ok {
		for _,item:=range val  {
			for key,_ :=range *item.Buffs  {
				switch key {
				case Damage:
					value := item.GetAtk()
					fighter.DeCeaseProperty(HP, value)
					battle.Log.HurtSide(battle.Frame, fighter.Id, value, fighter.GetProperty(HP), fighter.FrameState.SkillPoint, item.Form)
					if fighter.GetProperty(HP) <= 0 {
					//	//isDie = true
						//battle.remove(fighter.Id)
						//println(battle.Frame, ballistic.Form, "  击杀了  ", fighter.Id)
					}
					default:
				}
			}
		}
	}
}

const (
	FightSeq = iota
	DendSeq
)

const(
	Damage =iota
	Recover
	Contrl
	Such
)