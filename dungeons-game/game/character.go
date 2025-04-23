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
	Rarity      string
}

type Inventory struct {
	HealingPotion int
	DefensePotion int
	AttackPotion  int
}

const (
	RarityCommon    = "Common"
	RarityRare      = "Rare"
	RarityEpic      = "Epic"
	RarityLegendary = "Legendary"
)

const (
	ColorReset  = "\033[0m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorYellow = "\033[33m"
)

func ColorizeWeapon(w Weapon) string {
	switch w.Rarity {
	case RarityRare:
		return fmt.Sprintf("%s[%s]%s %s", ColorBlue, w.Rarity, ColorReset, w.Name)
	case RarityEpic:
		return fmt.Sprintf("%s[%s]%s %s", ColorPurple, w.Rarity, ColorReset, w.Name)
	case RarityLegendary:
		return fmt.Sprintf("%s[%s]%s %s", ColorYellow, w.Rarity, ColorReset, w.Name)
	default:
		return fmt.Sprintf("[%s] %s", w.Rarity, w.Name)
	}
}

var AllWeapons = []Weapon{{"Sword", 5, "A rusty looking blade used for close combat.", RarityCommon}, {"Staff", 3, "A sturdy staff that can unleash anger upon your enemies. Nerd.", RarityCommon},
	{"Dagger", 2, "A pocket dagger used for staby staby actions. It's quite shit but it sounds cool!", RarityCommon}, {"GreatAxe", 8, "Big, brutal and slow.", RarityRare},
	{"Magic Wand", 6, "A wand bestowed to you by ancient powers.", RarityRare}, {"Poisoned Blade", 4, "It stings after you swing. Edgelord.", RarityRare},
	{"Flaming sword", 10, "Engulfed in the God Emperor's holy fire", RarityEpic}, {"Blade of eternity", 15, "Forged by the gods.", RarityLegendary}}

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
			starterWeapons = []Weapon{AllWeapons[0]}
			selectedClass = Character{name, class, 100, 100, 15, 10, Inventory{}, 1, AllWeapons[0], starterWeapons}
		case "Mage":
			starterWeapons = []Weapon{AllWeapons[1]}
			selectedClass = Character{name, class, 70, 70, 20, 5, Inventory{}, 1, AllWeapons[1], starterWeapons}
		case "Rogue":
			starterWeapons = []Weapon{AllWeapons[2]}
			selectedClass = Character{name, class, 80, 80, 12, 8, Inventory{}, 1, AllWeapons[2], starterWeapons}
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
