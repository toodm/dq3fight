package model

import (
	"sort"
	"fmt"
	"github.com/json-iterator/go"
)


type Battle struct {
	Frame             int
	IsOk              bool
	FighterGroupMap   map[int]*[]Fighter // 阵营，
	FrameBallisticMap map[int]*[]Ballistic //每帧 技能
	FighterMap        map[int]int
	Log               *FighterLog

}

func (battle *Battle) PopRand() int {
	return 50
}

func (battle *Battle) GetDefsToForm(form int) *[]Fighter {
	return battle.FighterGroupMap[1-battle.FighterMap[form]]
}

func (battle *Battle) GetDefs(fighter Fighter) *[]Fighter {
	return battle.GetDefsToForm(fighter.Id)
}

func (battle *Battle) GetOneDef(fighter Fighter) int {
	return (*battle.GetDefs(fighter))[0].Id
}

func (battle *Battle) GetAlls() *[]Fighter {
	alls := *battle.FighterGroupMap[1]
	for _, item := range *battle.FighterGroupMap[0] {
		alls = append(alls, item)
	}
	return &alls
}

func (battle *Battle) remove(index int) {
	slice := *battle.FighterGroupMap[battle.FighterMap[index]]
	for i, item := range slice {
		if item.Id == int(index) {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	battle.FighterGroupMap[battle.FighterMap[index]] = &slice
}

func (battle *Battle) AddBallistic(result *Ballistic) {
	if val, ok := battle.FrameBallisticMap[result.EndFrame]; ok {
			*val = append(*val, *result)
	} else {
		arrs := []Ballistic{*result}
		battle.FrameBallisticMap[result.EndFrame] = &arrs

	}
}

func (battle *Battle) Fight() {
	battle.FrameBallisticMap = make(map[int]*[]Ballistic)
	battle.Log = &FighterLog{
		Sequence: &[][]map[string]interface{}{},
		LastFrame: 0,
		Alignment: GetAlignment(*battle),
	}

	battle.FighterMap = make(map[int]int)
	for key, value := range battle.FighterGroupMap {
		for _, item := range *value {
			item.InCeaseProperty(MaxHP, item.GetProperty(HP))
			battle.FighterMap[item.Id] = key
		}
	}

	for _, value := range battle.FighterGroupMap {
		printFighter(*battle, *value)
	}

	fmt.Println("ok")
	for {
		battle.Frame++
		println(battle.Frame)


		//生成当前帧所有英雄要出手的技能,并加入对应的帧
		for _, value := range battle.FighterGroupMap {
			for _, item := range *value {
				item.Attack(battle)
			}
		}

		if val, ok := battle.FrameBallisticMap[battle.Frame]; ok {
			for _, item := range *val {

				result := item.AttackFighter(battle)
				if result {
					battle.IsOk = len(*battle.FighterGroupMap[Atts]) == 0 || len(*battle.FighterGroupMap[Defs]) == 0
				}
			}
		}

		//当本轮能出手的时候，生成对应技能的弹道
		//处理本帧的弹道
		//找到对手，对其进行攻击,死亡判定，战斗结束判定 有点BUG，先打A再打B了，不公平

		if battle.Frame > 100000 {
			battle.IsOk = true
		}
		if battle.IsOk {
			battle.IsOk = true
			for _, value := range battle.FighterGroupMap {
				printFighter(*battle, *value)
			}
			log := map[string]interface{}{}
			if len(*battle.FighterGroupMap[Atts]) == 0 {
				battle.Log.SetWinner(1)
			} else
			{
				battle.Log.SetWinner(0)
			}
			log["sequence"] = battle.Log.Sequence
			log["alignment"] = battle.Log.Alignment
			log["lastFrame"] = battle.Log.LastFrame
			log["result"] = battle.Log.result
			log["timeUnit"] = 600
			println("共", battle.Frame, "帧")
			result, err := jsoniter.ConfigFastest.Marshal(log)
			if err == nil {
				//println(string(result))
				Write("E:\\daqing3\\client\\card0.1.0\\battle.json", string(result))
			}
			break;
		}
	}
}

func printFighter(battle Battle, fighters []Fighter) {
	for _, item := range fighters {
		println(battle.FighterMap[item.Id], item.Pos, "血量", item.GetProperty(HP), "攻击", item.GetProperty(ATK))
	}
}

func OrderSkill(ballistics []Ballistic) []Ballistic {
	ballisticMap := map[int][]Ballistic{}
	for _, dd := range ballistics {
		ballisticMap[GetSkillOrder(dd.SkillId)] = append(ballisticMap[GetSkillOrder(dd.SkillId)], dd)
	}
	var keys []int
	for k := range ballisticMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	result := []Ballistic{}
	for _, k := range keys {
		for _, item := range ballisticMap[ int(k)] {
			result = append(result, item)
		}
	}
	return result
}

func GetSkillOrder(skillId int) int {
	return skillId
}
