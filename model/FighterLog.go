package model

type FighterLog struct {
	Sequence *[][]map[string]interface{} // 帧，事件，事件内容
	Alignment map[string][]interface{}
	LastFrame int
	result    map[string]interface{}
}

func GetBaatleName(mark int) string {
	if mark == 0 {
		return "blue"
	} else {
		return "red"
	}
}

func (fighterLog *FighterLog) SetWinner(mark int) {
	fighterLog.result = make(map[string]interface{})
	fighterLog.result["winner"] = GetBaatleName(mark)
	fighterLog.result["award"] =  []interface{}{}
}

func GetAlignment(battle Battle) map[string][]interface{} {

	alignment := make(map[string][]interface{})
	for key, value := range battle.FighterGroupMap {
		name := GetBaatleName(key)
		alignment[name] = []interface{}{}
		for _, item := range *value {
			obj := map[string]interface{}{
				"pos":    item.Pos,
				"id":     item.Id,
				"heroId": item.HeroId,
				"HP":     item.GetProperty(HP),
				"MP":     item.FrameState.SkillPoint,
			}
			alignment[name] = append(alignment[name], obj)
		}

	}
	return alignment
}

// aFlag: 进攻方标记位
// skillId: 技能ID，0标示普攻
// orbId: 法球ID，规则：站位-帧序号-技能id
// HP, MP: 剩余血量，能量//
func (fighterLog *FighterLog) AddAttSide(frame int, id int, skillId int, mp int, target []int) {

	obj := map[string]interface{}{
		"action":  "attack",
		"id":      id,
		"skillId": skillId,
		"MP":      mp,
		"target":  target,
	}
	println(frame, "id :", id, " 抬手", "使用技能 ", skillId, " 当前能量 ", mp, "target", target[0])
	fighterLog.add(frame, obj)
}

func (fighterLog *FighterLog) AddAttack(frame int, id int, skillId int, mp int, target []int, endFrame int) {

	obj := map[string]interface{}{
		"action":  "orb",
		"id":      id,
		"skillId": skillId,
		"frame":   endFrame - frame,
		"target":  target,
	}
	println(frame, "id :", id, "出手", "使用技能 ", skillId, " 当前能量 ", mp, "target", target[0], "命中帧：", endFrame)
	fighterLog.add(frame, obj)
}

func (fighterLog *FighterLog) add(frame int, obj map[string]interface{}) {
	if fighterLog.LastFrame < frame {
		for i := fighterLog.LastFrame + 1; i <= frame; i++ {
			*fighterLog.Sequence = append(*fighterLog.Sequence, []map[string]interface{}{})
		}

		fighterLog.LastFrame = frame
	}
	index := frame - 1
	(*fighterLog.Sequence)[index] = append((*fighterLog.Sequence)[index], obj)
}

func (fighterLog *FighterLog) HurtSide(frame int, id int, damage int, hp int, mp int, skillId int) {

	obj := map[string]interface{}{
		"action": "hurt",
		"id":     id,
		"damage": damage,
		"HP":     hp,
		"MP":     mp,
	}
	println(frame, "id :", id, "受到 ", damage, " 伤害， 当前血量 ", hp, " 当前能量 ", mp, "技能ID", skillId)

	fighterLog.add(frame, obj)

}
func (fighterLog *FighterLog) TreatSide(frame int, id int, recover int, hp int, mp int) {

	obj := map[string]interface{}{
		"action": "cure",
		"id":     id,
		"recover": recover,
		"HP":     hp,
		"MP":     mp,
	}
	println(frame, "id :", id, "受到 ", recover, " 治疗， 当前血量 ", hp, " 当前能量 ", mp)

	fighterLog.add(frame, obj)

}
