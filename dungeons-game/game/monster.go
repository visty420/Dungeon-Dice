package game

import "math/rand"

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

// func SpawnMonster() Monster {
// 	monsters := []Monster{
// 		{"Sekelton", 30, 8},
// 		{"Goblin", 30, 10},
// 		{"Ork", 50, 12},
// 	}
// 	return monsters[rand.Intn(len(monsters))]
// }

func SpawnMonsterForLevel(level int) Monster {
	template := MonsterTemplates[rand.Intn(len(MonsterTemplates))]
	hp := template.HP + rand.Intn(level*3)
	str := template.Strength + level/2

	return Monster{template.Name, template.Type, hp, str, template.IsBoss, template.Description}
}
