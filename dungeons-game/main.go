package main

import (
	"dungeons-game/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%s=== Welcome to Dungeon Dice! ===%s\n", game.ColorTitle, game.ColorReset)
	player := game.CreateCharacter()
	player.Level = 1

	for {
		fmt.Printf("\n%s--- Level %d ---%s\n", game.ColorInfo, player.Level, game.ColorReset)

		fmt.Printf("Do you wish to change your weapon?(y/n)\nYou currently have the %s equipped.\n", game.ColorizeWeapon(player.Weapon))
		var choice string
		fmt.Scan(&choice)
		if choice == "y" {
			fmt.Println("Unlocked weapons: ")
			for i, w := range player.UnlockedWeapons {
				fmt.Printf("[%d] %s (+%d dmg) - %s\n", i+1, game.ColorizeWeapon(w), w.DamageBonus, w.Description)
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

		fmt.Println("\nWould you like to open your inventory before this fight? (y/n)")
		var inventoryChoice string
		fmt.Scanln(&inventoryChoice)
		if inventoryChoice == "y" {
			game.UsePotionsMenu(&player)
		} else {
			fmt.Println("Cheeky choice mate -- your next encounter might be your last. But fine, continuing without opening the inventory")
		}

		//isBoss := player.Level%5 == 0

		monster := game.SpawnMonsterForLevel(player.Level)
		fmt.Printf("\nYou are facing %s", monster.NameWithType())
		for monster.HP > 0 && player.HP > 0 {
			game.CombatRound(&player, &monster)
		}
		if player.HP <= 0 {
			fmt.Println("You died a horrible death! RIP my homie.")
			break
		}
		goldEarned := 10 + rand.Intn(5) + player.Level*2
		player.Gold += goldEarned
		fmt.Printf("You defeated the %s!\n", monster.ColorName())
		fmt.Printf("You have earned %d gold coins. Total gold coins: %d\n", goldEarned, player.Gold)
		game.TryDropItem(&player)
		game.TryDropWeapon(&player)
		player.Level++
		fmt.Println("Would you like to open the shop (y/n)?")
		var shopChoice string
		fmt.Scanln(&shopChoice)
		if shopChoice == "y" {
			game.OpenShop(&player)
		} else if shopChoice == "n" {
			fmt.Println("You go past the travelling merchant, ignoring it.")
			continue
		} else {
			fmt.Println("It's a yes or no question, yet you weren't able to input a simple letter. The merchant doesn't want to sell you his items. Good job, dork.")
			continue
		}

	}
}
