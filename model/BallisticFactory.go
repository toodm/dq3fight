package model

//技能
const (
	CommonAttack       = iota
	BallisticAttack
	AddHP
	DoubleAttack
	MultAttack
	Silent
)


func AttackMult(fighter Fighter, battle Battle,skillId int,toArr []int) []Ballistic{
	ballistics := []Ballistic{}
	dd := Ballistic{
		Form:     fighter.Id,
		To:       toArr,
		SkillId:  skillId,
		EndFrame: battle.Frame + fighter.GetProperty(skillId),
		Value:    fighter.GetProperty(ATK),
	}
	ballistics = append(ballistics, dd)
	return  nil
}


func BuildBallistic(fighter Fighter, skillId int,battle Battle) []Ballistic {

	ballistics := []Ballistic{}
	//目标选择
	//对单个敌人造成100%攻击伤害，每回合额外造成48%攻击伤害，持续2回合。
	//技能效果
	defs:=*battle.GetDefs(fighter)

	switch skillId {
	case DoubleAttack:
		dd := BuildBallisticModel(fighter, defs[0].Id, skillId, battle.Frame)
		ballistics = append(ballistics, dd)
		ballistics = append(ballistics, dd)
	case MultAttack:
		var toArr =[]int{}
		for _, item := range defs {
			toArr=	append(toArr, item.Id)
		}
		ballistics =  AttackMult(fighter,battle,skillId,toArr)
	case Silent:
		defs[0].SetSilent(battle.Frame,battle.Frame + 5)
		dd := BuildBallisticModel(fighter, defs[0].Id, skillId, battle.Frame)
		ballistics = append(ballistics, dd)
	default:
		dd := BuildBallisticModel(fighter, defs[0].Id, skillId, battle.Frame)
		ballistics = append(ballistics, dd)
	}
	return ballistics
}

func BuildBallisticModel(fighter Fighter, to int, skillId int, frame int) Ballistic {
	var arr=[]int{to}
	return Ballistic{
		Form:     fighter.Id,
		To:       arr,
		SkillId:  skillId,
		EndFrame: frame + GetSkillBallisticFly(skillId),
		Value:   10,//技能伤害
	}

}


