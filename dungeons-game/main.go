package main

import (
	"dungeons-game/game"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game.ShowSplashScreen()
	fmt.Println("Do you want to (1) Start New Game or (2)Load Saved Game")
	var saveLoadChoice int
	fmt.Scanln(&saveLoadChoice)

	var player game.Character

	if saveLoadChoice == 2 {
		loadedPlayer, err := game.LoadGameWithChoice()
		if err != nil {
			fmt.Println("Error loading save: ", err)
			fmt.Println("Starting a new game instead...")
			player = game.CreateCharacter()
			player.Level = 1
		} else {
			player = *loadedPlayer
		}
	} else {
		player = game.CreateCharacter()
		player.Level = 1
	}

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

		isBoss := player.Level%5 == 0 && player.Level > 0
		var monster game.Monster

		game.ClearScreen()

		if isBoss {
			monster = game.SpawnBossForLevel(player.Level)
			game.ClearScreen()
			game.CenterTextSmart(monster.AsciiArt)
			fmt.Println()
			game.CenterTextSmart(fmt.Sprintf("\nYou enocunter a %s\n", monster.ColorName()))
			fmt.Println()
			game.CenterTextSmart(monster.Description)
			fmt.Println()
		} else {
			monster = game.SpawnMonsterForLevel(player.Level)
			game.ClearScreen()
			game.CenterTextSmart(monster.AsciiArt)
			fmt.Println()
			game.CenterTextSmart(fmt.Sprintf("\nYou enocunter a %s\n", monster.ColorName()))
			fmt.Println()
			game.CenterTextSmart(monster.Description)
			fmt.Println()
		}
		fmt.Printf("\nYou are facing %s", monster.NameWithType())
		for monster.HP > 0 && player.HP > 0 {
			game.CombatRound(&player, &monster)
		}
		if player.HP <= 0 {
			fmt.Println("You died a horrible death! RIP my homie.")
			break
		}
		if monster.IsBoss {
			fmt.Printf("\n%s: %s\n", monster.ColorName(), monster.DeathLine)
			bonusGold := 100 + player.Level*2
			player.Gold += bonusGold
			fmt.Printf("You defeated the boss and earned %d bonus gold!", bonusGold)
			game.TryDropWeapon(&player)
		}
		goldEarned := 10 + rand.Intn(5) + player.Level*2
		player.Gold += goldEarned
		fmt.Printf("You defeated the %s!\n", monster.ColorName())
		fmt.Printf("You have earned %d gold coins. Total gold coins: %d\n", goldEarned, player.Gold)
		game.TryDropItem(&player)
		game.TryDropWeapon(&player)

		fmt.Println("You have survived this level... your journey continues...")
		fmt.Println("Would you like to save your progress? (y/n)")
		var saveChoice string
		fmt.Scanln(&saveChoice)

		if saveChoice == "y" || saveChoice == "Y" {
			err := game.AutoSaveGame(&player)
			if err != nil {
				fmt.Println("Warning: Could not save the game:", err)
			}
		} else {
			fmt.Println("Continuing without saving...")
		}

		var continueChoice string
		fmt.Println()
		fmt.Println("Do you wish to (c) Continue exploring the dungeon or do you wish to (q) Quit while you are still able to...")
		fmt.Scanln(&continueChoice)
		if continueChoice == "q" || continueChoice == "Q" {
			fmt.Print("Your journey ends here...")
			game.Countdown(3)
			os.Exit(0)
		}

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

//test commit
