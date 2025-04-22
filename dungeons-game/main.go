package main

import (
	"dungeons-game/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Welcome to Dungeon Dice! ===")
	player := game.CreateCharacter()
	player.Level = 1

	for {
		fmt.Printf("\n --- Level %d --- \n", player.Level)
		monster := game.SpawnMonsterForLevel(player.Level)
		for monster.HP > 0 && player.HP > 0 {
			game.CombatRound(&player, &monster)
		}
		if player.HP <= 0 {
			fmt.Println("You died a horrible death! RIP my homie.")
			break
		}
		fmt.Printf("You defeated the %s!\n", monster.Name)
		game.TryDropItem(&player)
		player.Level++
	}
}
