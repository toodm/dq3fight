package model

type Ballistic struct {
	Form     int
	To       []int
	EndFrame int
	SkillId  int
	Value    int
	Buffs    *map[int][]BuffConfig
}

func (ballistic Ballistic) GetSkillType() int {
	if ballistic.SkillId == AddHP {
		return Recover
	}
	return Damage
}

func (ballistic Ballistic) AttackFighter(battle *Battle) bool {
	isDie := false
	for _, ele := range ballistic.To {

		atts := battle.GetAlls()
		for _, item := range *atts {

			if item.Id == ele {
				//防守和被恢复治疗，可写到fighter类中
				isDie = item.Defend(ballistic, battle)
			}
		}
	}
	if isDie {
		return true
	}
	return false
}
