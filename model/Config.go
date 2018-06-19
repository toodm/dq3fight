package model

type SkillConfig struct {
	Id string
	Index int
	Stage int
	Duration int	//触发帧增量
	BallisticFly	int		//弹道飞行帧数
	Property int	//技能附带属性
}

type BuffConfig struct {
	Id string
	SkillId string
	BuffProperty int	//属性（冰、金、木..）
	BuffType     int // 伤害，恢复，控制
	Value        int	//数值
}

type HeroConfig struct {
	HeroId int
}

//施法间隔
func GetSkillInterval(skillId int) int {
	return 2
}

func GetDamageInterval(skillId int) int {
	if skillId == 0 {
		return 0
	}
	return 1
}

var skillMap = map[int][]SkillConfig{}

var BuffMap = map[int]map[int][]BuffConfig{}

func SkillConfigInit(){

	//b1:=BuffConfig{
	//
	//}
	BuffMap=make(map[int]map[int][]BuffConfig)
	skillMap = make(map[int][]SkillConfig)

	s1:=SkillConfig{
		Stage:0,
		Duration:2,
		BallisticFly :0,
		Property:Non,
	}
	skillMap[CommonAttack] = []SkillConfig{s1}
	BuffMap[CommonAttack]=map[int][]BuffConfig{}

	s2:=SkillConfig{
		Stage:0,
		Duration:2,
		BallisticFly :2,
		Property:Non,
	}
	skillMap[BallisticAttack] = []SkillConfig{s2}
	//b2:=BuffConfig{
	//}
	BuffMap[BallisticAttack]=map[int][]BuffConfig{}


	s3:=SkillConfig{
		Stage:0,
		Duration:2,//
		BallisticFly :2,
		Property:Non,
	}
	skillMap[AddHP] = []SkillConfig{s3}
	//b2:=BuffConfig{
	//}
	BuffMap[AddHP]=map[int][]BuffConfig{}
	//BallisticAttack
	//AddHP
	//DoubleAttack
	//MultAttack
	//Silent
	//BuildBallistic//

}

func getSkill(skillId int) []SkillConfig{
	if val, ok := skillMap[skillId]; ok {
		return  val
	}
	return  nil
}

func GetSkillDuration(skillId int) int {
	return  getSkill(skillId)[0].Duration
}

func GetSkillStage(skillId int,stage int)int {
	return  getSkill(skillId)[stage].Duration
}

func GetSkillBallisticFly(skillId int) int {
	return  getSkill(skillId)[0].BallisticFly
}

//伤害类型
//属性buff

func GetSkillDamage(skillId int) map[int][]BuffConfig{
	if val, ok := BuffMap[skillId]; ok {
		return  val
	}
	return  nil
}