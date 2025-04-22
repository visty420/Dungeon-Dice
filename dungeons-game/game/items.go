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

var rarityDropChances = map[string]int{
	RarityLegendary: 10,
	RarityEpic:      15,
	RarityRare:      30,
	RarityCommon:    45,
}

func getWeaponByRarity(rarity string) []Weapon {
	var filtered []Weapon
	for _, w := range AllWeapons {
		if w.Rarity == rarity {
			filtered = append(filtered, w)
		}
	}
	return filtered
}

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

func getLockedWeaponsByRarity(player *Character, rarity string) []Weapon {
	unlocked := make(map[string]bool)
	for _, w := range player.UnlockedWeapons {
		unlocked[w.Name] = true
	}
	var locked []Weapon
	for _, w := range AllWeapons {
		if w.Rarity == rarity && !unlocked[w.Name] {
			locked = append(locked, w)
		}
	}
	return locked
}

func TryDropWeapon(player *Character) {
	roll := rand.Intn(100)
	if roll >= 50 {
		return
	}
	rarityRoll := rand.Intn(100)
	var selectedRarity string
	switch {
	case rarityRoll < rarityDropChances[RarityLegendary]:
		selectedRarity = RarityLegendary
	case rarityRoll < rarityDropChances[RarityLegendary]+rarityDropChances[RarityEpic]:
		selectedRarity = RarityEpic
	case rarityRoll < rarityDropChances[RarityLegendary]+rarityDropChances[RarityEpic]+rarityDropChances[RarityRare]:
		selectedRarity = RarityRare
	default:
		selectedRarity = RarityCommon
	}
	candidates := getLockedWeaponsByRarity(player, selectedRarity)
	if len(candidates) == 0 {
		return
	}
	newWeapon := candidates[rand.Intn(len(candidates))]
	player.UnlockedWeapons = append(player.UnlockedWeapons, newWeapon)

	var flavor string
	switch newWeapon.Rarity {
	case RarityCommon:
		flavor = "You find a simple but useful weapon."
	case RarityRare:
		flavor = "You stumble upon a weapon radiating strange power."
	case RarityEpic:
		flavor = "Your hands tramble. This weapon is radiating with energy."
	case RarityLegendary:
		flavor = "The world goes silent. Destiny has chosen you."
	}

	fmt.Printf("\n **** You found a %s! (%s) ****\n", ColorizeWeapon(newWeapon), newWeapon.Description)
	fmt.Println(flavor)
}
