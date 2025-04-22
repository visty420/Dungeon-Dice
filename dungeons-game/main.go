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

		fmt.Println("Do you wish to equip a weapon? y/n")
		var choice string
		fmt.Scan(&choice)
		if choice == "y" {
			fmt.Println("Unlocked weapons: ")
			for i, w := range player.UnlockedWeapons {
				fmt.Printf("[%d] %s (+%d dmg) -%s\n", i+1, w.Name, w.DamageBonus, w.Description)
			}
			fmt.Println("Please enter weapon number: ")
			var pick int
			fmt.Scan(&pick)
			if pick > 0 && pick <= len(player.UnlockedWeapons) {
				player.Weapon = player.UnlockedWeapons[pick-1]
				fmt.Printf("You have equipped the %s\n", player.Weapon.Name)
			} else {
				fmt.Println("Invalid selection, keeping current weapon.")
			}
		}

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
		game.TryDropWeapon(&player)
		player.Level++
	}
}
