package game

import (
	"fmt"
	"math/rand"
)

type MonsterType string

const (
	Goblin   MonsterType = "Goblin"
	Skeleton MonsterType = "Skeleton"
	Ork      MonsterType = "Ork"
	Wraith   MonsterType = "Wraith"
	Demon    MonsterType = "Demon"
	Succubus MonsterType = "Succubus"
)

type Monster struct {
	Name        string
	Type        MonsterType
	HP          int
	Strength    int
	IsBoss      bool
	Description string
	IntroLine   string
	DeathLine   string
	AsciiArt    string
}

var MonsterTemplates = []Monster{
	{"Goblin", Goblin, 30, 8, false, "A sneaky little creature with sharp teeth.", "A green creature approaches you!", "You slay the green creature", GoblinArt},
	{"Skeleton", Skeleton, 40, 10, false, "Clattering bones that won't stay down,", "You encounter a restless pile of bones.", "You have granted peace to the pile of bones", SkeletonArt},
	{"Ork", Ork, 50, 12, false, "Strong, brutal and always looking for a scrap.", "A red angry ork abushes you!", "You have slain the strange creature", OrcArt},
	{"Wraith", Wraith, 35, 14, false, "It strikes from the shadows and ignores your armor,", "You can hardly see a transparent ghost.", "You banished the ghost.", WraithArt},
	{"Demon", Demon, 60, 16, false, "Burns everything it touches.", "A creature from hell is blocking your path", "You exorcize the demon, sending it to the depths of hell.", DemonArt},
	{"Succubus", Succubus, 35, 15, false, "You'd tap that sucubussy, but you know better then to do so.", "You encounter a creature that stirs your interest.", "After coming to your senses, you feel the creature's treachery and banish it from existence.", SuccubusArt},
}

var Bosses = []Monster{
	{"Bone Crusher", Skeleton, 150, 25, true, "An ancient titan of bone and hatred. It swings with unstoppable force.", "-You dare approach me mortal? This carcass has a lot of fight left in it!", "-This cannot happen! DAMN YOU HERO!", BoneCrusherArt},
	{"Inferna, Flame Witch", Demon, 130, 30, true, "Her eyes glow like embers. Her eyes.. a burning whisper.", "You dare stand against me?", "Even though my form from this realm is broken, I will curse upon you from hell!", InfernaArt},
	{"The Lich King", Wraith, 160, 28, true, "The air freezes. Time slows. Death watches you from behind hollow eyes.", "You have finally come to meet me, messenger. Let me teach you something....", "You might have banished me and beaten the game, but my tainted touch is all over this realm!", LichKingArt},
}

func SpawnMonsterForLevel(level int) Monster {
	template := MonsterTemplates[rand.Intn(len(MonsterTemplates))]
	hp := template.HP + rand.Intn(level*3)
	str := template.Strength + level/2

	return Monster{template.Name, template.Type, hp, str, template.IsBoss, template.Description, template.IntroLine, template.DeathLine, template.AsciiArt}
}

func SpawnBossForLevel(level int) Monster {
	index := (level/5 - 1) % len(Bosses)
	return Bosses[index]
}

func (m *Monster) PreCombatEffect(player *Character) bool {
	switch m.Type {
	case Goblin:
		if rand.Intn(100) < 20 {
			fmt.Printf("The %s dodged your attack!\n", m.ColorName())
			return true
		}
	case Skeleton:
		if rand.Intn(100) < 15 {
			fmt.Printf("The %s blocked your attack!\n", m.ColorName())
			return true
		}
	case Ork:
		if rand.Intn(100) < 10 {
			fmt.Printf("The %s is enraged!\n", m.ColorName())
			m.Strength += 5
		}
	case Wraith:
		fmt.Printf("The %s phases eerily...your armor becomes dead weight.\n", m.ColorName())
	case Demon:
		burn := rand.Intn(3) + 2
		player.HP -= burn
		fmt.Printf("The %ss fire aura burns you for %d damage!\n", m.ColorName(), burn)
	case Succubus:
		fmt.Printf("The %s seduces you and lowers your strength by 2! What a shame.\n", m.ColorName())
		player.Strength -= 2
	}
	return false
}

func (m *Monster) UseSpecialMove(player *Character) {
	if !m.IsBoss {
		return
	}
	switch m.Name {
	case "Bonecursher":
		if rand.Intn(100) < 20 {
			fmt.Printf("%s slams the ground and stuns you!\n", m.ColorName())
			player.Stunned = true
		}
	case "Inferna, Flame Witch":
		if rand.Intn(100) < 30 {
			fmt.Printf("%s engulfs you in flames! You are burning!\n", m.ColorName())
			player.BurnDuration = 3
		}
	case "The Lich King":
		if rand.Intn(100) < 25 {
			heal := rand.Intn(10) + 10
			m.HP += heal
			fmt.Printf("%s absorbs the life around him and heals for %d HP!\n", m.ColorName(), heal)

		}
	}
}

func (m *Monster) NameWithType() string {
	var color string
	switch m.Type {
	case Goblin:
		color = ColorGoblin
	case Skeleton:
		color = ColorSkeleton
	case Ork:
		color = ColorOrk
	case Wraith:
		color = ColorWraith
	case Demon:
		color = ColorDemon
	case Succubus:
		color = ColorSuccubus
	}
	return fmt.Sprintf("a random %s%s%s", color, m.Type, ColorReset)
}

func (m *Monster) ColorName() string {
	var color string
	switch m.Type {
	case Goblin:
		color = ColorGoblin
	case Skeleton:
		color = ColorSkeleton
	case Ork:
		color = ColorOrk
	case Wraith:
		color = ColorWraith
	case Demon:
		color = ColorDemon
	case Succubus:
		color = ColorSuccubus
	}
	return fmt.Sprintf("%s%s%s", color, m.Name, ColorReset)
}
