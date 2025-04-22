package game

import (
	"fmt"
	"math/rand"
)

type Item string

const (
	HealingPotion Item = "Healing Potion"
	AttackPotion  Item = "Attack Potion"
	DefensePotion Item = "Defense Potion"
)

func TryDropItem(player *Character) {
	roll := rand.Intn(100)
	if roll < 30 {
		switch rand.Intn(3) {
		case 0:
			fmt.Println("You got a healing potion!")
			player.HealingPotion++
		case 1:
			fmt.Println("You got an attack potion!")
			player.AttackPotion++
		case 2:
			fmt.Println("You got a defense potion!")
			player.DefensePotion++
		default:
			fmt.Println("The RNG Gods frown upon you -- no items this time!")
		}
	}
}

func getLockedWeapons(player *Character) []Weapon {
	unlockedMap := make(map[string]bool)
	for _, w := range player.UnlockedWeapons {
		unlockedMap[w.Name] = true
	}
	var locked []Weapon
	for _, w := range AllWeapons {
		if !unlockedMap[w.Name] {
			locked = append(locked, w)
		}
	}
	return locked
}

func TryDropWeapon(player *Character) {
	roll := rand.Intn(100)
	if roll < 25 {
		locked := getLockedWeapons(player)
		if len(locked) > 0 {
			newWeapon := locked[rand.Intn(len(locked))]
			player.UnlockedWeapons = append(player.UnlockedWeapons, newWeapon)

			var flavor string
			switch newWeapon.Rarity {
			case RarityCommon:
				flavor = "You find a basic but useful weapon."
			case RarityRare:
				flavor = "Your hands tremble with a surge of power!"
			case RarityEpic:
				flavor = "You feel like a one man army."
			case RarityLegendary:
				flavor = "You found a strange items that seems to be forged by the old gods. It radiates with power."
			}

			fmt.Printf("\n*** You found a %s! (%s) ***\n", ColorizeWeapon(newWeapon), newWeapon.Description)
			fmt.Println(flavor)
		} else {
			fmt.Println("No weapons left to unlock. You are the weapon master.")
		}
	}
}
