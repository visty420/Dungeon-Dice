package game

import (
	"fmt"
	"math/rand"
)

func CombatRound(player *Character, monster *Monster) {
	fmt.Printf("\n%s=====COMBAT ROUND=====%s\n", ColorSubtitle, ColorReset)
	if monster.PreCombatEffect(player) {
		return
	}
	playerRoll := rand.Intn(20) + 1 + player.Weapon.DamageBonus + player.TempAttackBoost
	monsterRoll := rand.Intn(20) + 1

	fmt.Printf("You roll: %d \n", playerRoll)
	Countdown(2)
	fmt.Printf("The %s rolls: %d\n", monster.ColorName(), monsterRoll)

	if player.Stunned {
		fmt.Println("You are stunned and lose your turn!")
		player.Stunned = false
		return
	}

	if player.BurnDuration > 0 {
		burnDamage := 5
		player.HP -= burnDamage
		fmt.Printf("You are burning! You take %d fire damage!\n", burnDamage)
		player.BurnDuration--
	}

	if playerRoll == 1 {
		selfDamage := rand.Intn(5) + 1
		player.HP -= selfDamage
		fmt.Printf("CRITICAL MISS! You stumble and hurt yourself for %d damage! Your HP: %d\n", selfDamage, player.HP)
	}

	if playerRoll == 20 {
		criticalDamage := player.Strength * 2
		monster.HP -= criticalDamage
		fmt.Printf("CRITICAL HIT! You unleash the powers of heavon upon the damned %s, dealing %d damage! Monster's HP: %d\n", monster.ColorName(), criticalDamage, monster.HP)
	}

	if monsterRoll == 1 {
		selfDamage := rand.Intn(5) + 1
		monster.HP -= selfDamage
		fmt.Printf("THE FOOLISH CREATURE MISSES YOU! The %s stumbles and hits itself for %d damage like a moron! Monster's HP: %d\n", monster.ColorName(), selfDamage, monster.HP)
	}

	if monsterRoll == 20 {
		criticalDamage := monster.Strength * 2
		player.HP -= criticalDamage
		fmt.Printf("OUCH! The %s curses upon your bloodline and hits you like a truck for %d damage! Your HP: %d\n", monster.ColorName(), criticalDamage, player.HP)
	}

	if playerRoll > monsterRoll {
		damage := playerRoll - monsterRoll
		monster.HP -= damage
		fmt.Printf("You have hit the %s for %d damage! MonsterHP: %d\n", monster.ColorName(), damage, monster.HP)
	} else {
		if monster.IsBoss {
			monster.UseSpecialMove(player)
		}
		damage := monsterRoll - playerRoll - player.Defense - player.TempDefenseBoost
		if damage < 0 {
			damage = 0
		}
		player.HP -= damage
		fmt.Printf("The %s hits you for %d damage! YourHP: %d\n", monster.ColorName(), damage, player.HP)
	}
	player.TempAttackBoost = 0
	player.DefensePotion = 0
}
