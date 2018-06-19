package model

type FighterFrame struct {
	State         int
	Frame         int
	UseSkill      bool //释放技能
	Stage         int //阶段，默认只有1个阶段
	Ballistics    *Ballistic
	NoSkillFrame  int
	SkillPoint    int
	IntervalFrame int
}

//生成弹道，谁的弹道 什么技能 什么时候造成伤害  多少伤害
func (frameState *FighterFrame) BuildBallistic(fighter Fighter, battle Battle) {

	frameState.UseSkill = frameState.NoSkillFrame < battle.Frame && frameState.SkillPoint >= 100
	index, skillId := fighter.GetSkill(frameState.UseSkill)
	frameState.Ballistics = &Ballistic{
		Form:    fighter.Id,
		SkillId: skillId,
		Value:   0,
		To:      []int{battle.GetOneDef(fighter)},
		Buffs:   &map[int][]BuffConfig{},
	}
	battle.Log.AddAttSide(battle.Frame, fighter.Id, index, frameState.SkillPoint, frameState.Ballistics.To)
}

func (frameState *FighterFrame) AddBallistic(fighter Fighter, battle *Battle) {

	ballistics := frameState.Ballistics
	ballistics.EndFrame = battle.Frame + GetDamageInterval(frameState.Ballistics.SkillId)
	ballistics.Value = fighter.GetAtk(frameState.Ballistics.SkillId)
	battle.AddBallistic(ballistics)
	if ballistics.SkillId != CommonAttack {
		battle.Log.AddAttack(battle.Frame, fighter.Id, ballistics.SkillId, frameState.SkillPoint, ballistics.To, ballistics.EndFrame)
	}
}

func (frameState *FighterFrame) Attack(fighter *Fighter, battle *Battle)  {
	//如果玩家当前处于施法中
	if frameState.Frame == battle.Frame {
		if frameState.State == Star {
			frameState.BuildBallistic(*fighter, *battle)
			frameState.StarContrl(*battle)
		} else if frameState.State == Attack {
			frameState.AddBallistic(*fighter, battle)
			frameState.AttackContrl(*fighter)

		}
	}
}

func (frameState *FighterFrame) StarContrl(battle Battle) {
	if !frameState.UseSkill {
		frameState.SkillPoint += 20
	}
	frameState.Stage = -1
	frameState.Frame = battle.Frame + GetSkillDuration(frameState.Ballistics.SkillId)
	frameState.State = Attack
}

func (frameState *FighterFrame) AttackContrl(fighter Fighter) {
	if frameState.Stage == -1 {
		//弹道释放初始化
		if frameState.UseSkill {
			frameState.SkillPoint = 0
		}
		//在这里判断技能阶段
		frameState.Stage = 1


	}
	frameState.Stage--

	if frameState.Stage == 0 {
		//结束释放
		frameState.Ballistics = nil
		frameState.State = Star
		frameState.Frame += GetHeroAttackInterval(fighter.HeroId)
	} else {
		frameState.Frame += GetSkillStage(frameState.Ballistics.SkillId, frameState.Stage)
	}
}

func (frameState *FighterFrame) AttackAfter() {

}

func (frameState *FighterFrame) InCreasePoint(value int) {
	frameState.SkillPoint += value
}

func (frameState *FighterFrame) SetSilent(curFrame int, endFrame int) {

	if frameState.State == Attack {
		frameState.Frame = curFrame + frameState.IntervalFrame
		frameState.State = Star
	}

	if frameState.NoSkillFrame < endFrame {
		frameState.NoSkillFrame = endFrame
	}
}

func (frameState *FighterFrame) SetStop(endFrame int) {
	if frameState.Frame < endFrame {
		frameState.Frame = endFrame
	}

	frameState.State = Star
}
