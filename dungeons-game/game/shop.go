package game

import (
	"fmt"
	"math/rand"
)

func OpenShop(player *Character) {
	for {
		fmt.Printf("\n%s=====WELCOME TO THE SHOP, TAINTED ONE=====%s\n", ColorTitle, ColorReset)
		fmt.Printf("%sGold%s: %d\n", ColorImportant, ColorReset, player.Gold)
		fmt.Println("1. Healing potion (10g)")
		fmt.Println("2. Attack potion (15g)")
		fmt.Println("3. Defense potion (15g)")
		fmt.Println("4. Mystery weapon chest (30g)")
		fmt.Println("5. Close shop.")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if player.Gold >= 10 {
				player.Gold -= 10
				fmt.Printf("You've bought a healing potion. Your HP: %d\n", player.HP)
				player.HealingPotion++
			} else {
				fmt.Println("You don't have enough gold brokie.")
			}
		case 2:
			if player.Gold >= 15 {
				player.Gold -= 15
				fmt.Println("You've bought an attack potion.")
				player.AttackPotion++
			} else {
				fmt.Println("You don't have enough gold brokie.")
			}
		case 3:
			if player.Gold >= 15 {
				player.Gold -= 15
				fmt.Println("You've bought a defense potion")
				player.DefensePotion++
			} else {
				fmt.Println("You don't have enough gold, brokie.")
			}
		case 4:
			if player.Gold >= 30 {
				player.Gold -= 30
				dropRandomWeapon(player)
			} else {
				fmt.Println("You don't have enough gold, brokie.")
			}
		case 5:
			fmt.Println("Exiting the shop...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func dropRandomWeapon(player *Character) {
	locked := getLockedWeapons(player)
	if len(locked) == 0 {
		fmt.Println("You already have all the weapons unlocked")
		return
	}
	weapon := locked[rand.Intn(len(locked))]
	player.UnlockedWeapons = append(player.UnlockedWeapons, weapon)
	fmt.Printf("Mystery chest reward: %s! (%s)\n", ColorizeWeapon(weapon), weapon.Description)
}
