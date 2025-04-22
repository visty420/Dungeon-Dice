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
	Level           int
	Weapon          Weapon
	UnlockedWeapons []Weapon
}

type Weapon struct {
	Name        string
	DamageBonus int
	Description string
}

type Inventory struct {
	HealingPotion int
	DefensePotion int
	AttackPotion  int
}

var AllWeapons = []Weapon{{"Sword", 5, "A rusty looking blade used for close combat."}, {"Staff", 3, "A sturdy staff that can unleash anger upon your enemies. Nerd."},
	{"Dagger", 2, "A pocket dagger used for staby staby actions. It's quite shit but it sounds cool!"}, {"GreatAxe", 8, "Big, brutal and slow."},
	{"Magic Wand", 6, "A wand bestowed to you by ancient powers."}, {"Poisoned Blade", 4, "It stings after you swing. Edgelord."}}

func CreateCharacter() Character {
	var name string
	var class string
	//var weaponChoice string
	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)

	var selectedClass Character
	starterWeapons := []Weapon{AllWeapons[0], AllWeapons[1], AllWeapons[2]}

	for {
		fmt.Println("Please choose your class [Warrior|Mage|Rogue]")
		fmt.Scan(&class)

		switch class {
		case "Warrior":
			selectedClass = Character{name, class, 100, 100, 15, 10, Inventory{}, 1, starterWeapons[0], starterWeapons}
		case "Mage":
			selectedClass = Character{name, class, 70, 70, 20, 5, Inventory{}, 1, starterWeapons[1], starterWeapons}
		case "Rogue":
			selectedClass = Character{name, class, 80, 80, 12, 8, Inventory{}, 1, starterWeapons[2], starterWeapons}
		default:
			fmt.Println("Please enter a valid class!")
			continue
		}
		break
	}
	// fmt.Println("Choose a weapon: ")
	// fmt.Println("[Sword](+5 dmg), [Staff](+3 dmg), [Dagger](+2 dmg, sounds cool)")
	// fmt.Println("You choice hero: ")
	// fmt.Scanln(&weaponChoice)

	// switch weaponChoice {
	// case "Sword":
	// 	selectedClass.Weapon = Weapon{"Sword", 5, "A rusty looking blade used for close combat."}
	// case "Staff":
	// 	selectedClass.Weapon = Weapon{"Staff", 3, "A sturdy staff that can unleash anger upon your enemies. Nerd."}
	// case "Dagger":
	// 	selectedClass.Weapon = Weapon{"Dagger", 2, "A pocket dagger used for staby staby actions. It's quite shit but it sounds cool!"}
	// default:
	// 	selectedClass.Weapon = Weapon{"Fists", 0, "You are a man of culture, you enjoy an honest brawl."}
	// }
	selectedClass.UnlockedWeapons = []Weapon{AllWeapons[0], AllWeapons[1], AllWeapons[2]}
	return selectedClass
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
