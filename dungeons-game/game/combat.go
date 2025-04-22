package game

import (
	"fmt"
	"math/rand"
)

func CombatRound(player *Character, monster *Monster) {
	fmt.Println("\n---Combat Turn---")
	playerRoll := rand.Intn(player.Strength) + 1
	monsterRoll := rand.Intn(monster.Strength) + 1

	fmt.Printf("You roll: %d \n", playerRoll)
	fmt.Printf("The monster rolls: %d\n", monsterRoll)

	if playerRoll > monsterRoll {
		damage := playerRoll - monsterRoll
		monster.HP -= damage
		fmt.Printf("You have hit the %s for %d damage! MonsterHP: %d\n", monster.Name, damage, monster.HP)
	} else {
		damage := monsterRoll - playerRoll - player.Defense
		if damage < 0 {
			damage = 0
		}
		player.HP -= damage
		fmt.Printf("The %s hits you for %d damage! YourHP: %d\n", monster.Name, damage, player.HP)
	}
}
