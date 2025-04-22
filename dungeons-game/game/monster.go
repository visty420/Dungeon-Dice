package game

import "math/rand"

type Monster struct {
	Name     string
	HP       int
	Strength int
}

func SpawnMonster() Monster {
	monsters := []Monster{
		{"Sekelton", 30, 8},
		{"Goblin", 30, 10},
		{"Ork", 50, 12},
	}
	return monsters[rand.Intn(len(monsters))]
}

func SpawnMonsterForLevel(level int) Monster {
	base := 30 + (level * 10)
	strength := 8 + (level * 2)
	names := []string{"Goblin", "Skeleton", "Ork", "Wraith", "Demon"}
	name := names[rand.Intn(len(names))]

	return Monster{name, base, strength}
}
