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
			fmt.Printf("You found a new weapon: %s. (%s)\n", newWeapon.Name, newWeapon.Description)
		} else {
			fmt.Println("No weapons left to unlock. You are the weapon master.")
		}
	}
}
