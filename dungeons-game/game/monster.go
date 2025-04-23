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
}

var MonsterTemplates = []Monster{
	{"Goblin", Goblin, 30, 8, false, "A sneaky little creature with sharp teeth."},
	{"Skeleton", Skeleton, 40, 10, false, "Clattering bones that won't stay down,"},
	{"Ork", Ork, 50, 12, false, "Strong, brutal and always looking for a scrap."},
	{"Wraith", Wraith, 35, 14, false, "It strikes from the shadows and ignores your armor,"},
	{"Demon", Demon, 60, 16, false, "Burns everything it touches."},
	{"Succubus", Succubus, 35, 15, false, "You'd tap that sucubussy, but you know better then to do so."},
}

func SpawnMonsterForLevel(level int) Monster {
	template := MonsterTemplates[rand.Intn(len(MonsterTemplates))]
	hp := template.HP + rand.Intn(level*3)
	str := template.Strength + level/2

	return Monster{template.Name, template.Type, hp, str, template.IsBoss, template.Description}
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
