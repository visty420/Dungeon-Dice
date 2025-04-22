package game

import "fmt"

type Character struct {
	Name     string
	Class    string
	HP       int
	MaxHP    int
	Strength int
	Defense  int
	Inventory
	Level int
}

type Inventory struct {
	HealingPotion int
	DefensePotion int
	AttackPotion  int
}

func CreateCharacter() Character {
	var name string
	var class string
	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)

	for {
		fmt.Println("Please choose your class [Warrior|Mage|Rogue]")
		fmt.Scan(&class)

		switch class {
		case "Warrior":
			return Character{name, class, 100, 100, 15, 10, Inventory{}, 1}
		case "Mage":
			return Character{name, class, 70, 70, 20, 5, Inventory{}, 1}
		case "Rogue":
			return Character{name, class, 80, 80, 12, 8, Inventory{}, 1}
		default:
			fmt.Println("Please enter a valid class!")
		}
	}
}

func (c *Character) UseItem(item Item) {
	switch item {
	case HealingPotion:
		if c.HealingPotion > 0 {
			heal := 30
			c.HP += heal
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
		}
		c.HealingPotion--
		fmt.Println("You have used a healing potion! You restored 30HP.")
	case AttackPotion:
		if c.AttackPotion > 0 {
			c.Strength += 5
			c.AttackPotion--
			fmt.Println("You have used an attack potion. Your strength grows -- now your attacks deal 5 more damage!")
		}
	case DefensePotion:
		if c.DefensePotion > 0 {
			c.Defense += 5
			c.DefensePotion--
			fmt.Println("You have used a defense potion! Defense boosted by 5 for this round.")
		}
	}
}
