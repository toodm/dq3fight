package main

import (
	"./model"
	"time"
	"fmt"
)

func main() {
	model.SkillConfigInit()
	var fighterAttrs = [][]int{{1,2000, 10, 3, 0, 5, 0}, {1,1000, 20, 5, 0, 5, 2}, {1,6600, 10, 5, 0, 4, 0}, {2,150, 20, 5, 0, 5, 0}, {2,66, 10, 5, 0, 13, 0}}

	var attFighters = GetFightersData(fighterAttrs, 0)
	var defAttrs = [][]int{{1,2000, 15, 2, 0, 10, 0}, {1,1088, 20, 8, 0, 5, 1}, {2,77, 10, 5, 0, 6, 0}, {2,88, 20, 8, 0, 10, 0}, {1,77, 5, 5, 0, 2, 0}}
	var defFighters = GetFightersData(defAttrs, int(len(defAttrs)))
	var battleMap = map[int](*[]model.Fighter){}
	battleMap[model.Atts] = &attFighters
	battleMap[model.Defs] = &defFighters

	var battle = model.Battle{FighterGroupMap: battleMap, Frame: 1, IsOk: false}
	t1 := time.Now()
	battle.Fight()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

}

func GetFightersData(fighterAttrs [][]int, startIndex int) []model.Fighter {
	fighters := []model.Fighter{}
	for i, item := range fighterAttrs {
		var fighter = GetFighterData(item, i)
		fighter.Id = i + 1 + startIndex
		fighter.Pos = i + 1
		fighter.HeroId = i + 1 + startIndex
		frameState := model.FighterFrame{
			State:         model.Star,
			Frame:         fighter.GetProperty(model.IntervalFrame),
			IntervalFrame: fighter.GetProperty(model.IntervalFrame),
		}

		fighter.FrameState = &frameState
		fighters = append(fighters, fighter)
	}
	return fighters
}

func GetFighterData(values []int, index int) model.Fighter {
	var fight1 model.Fighter
	attributes := map[int]*model.FighterProperty{}
	skills := []model.FighterSkill{}
	buffs := []model.FigheterBuff{}
	for i, item := range values {
		attr := &model.FighterProperty{i, item}
		attributes[i] = attr
	}
	//for i := 0; i < 3; i++ {
	//	skill := model.FighterSkill{SkillId: i} //先假设都是技能1
	//	skills = append(skills, skill)
	//}
	for i := 0; i < 3; i++ {
		buff := model.FigheterBuff{BuffId: i, BuffType: i, EndFrame: 1000, Value: 10}
		buffs = append(buffs, buff)
	}
	fight1.Propertys = attributes
	fight1.Skills = skills
	fight1.Buffs = &buffs
	return fight1
}
