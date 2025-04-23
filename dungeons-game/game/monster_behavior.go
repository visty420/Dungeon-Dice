package game

// import (
// 	"fmt"
// 	"math/rand"
// )

// func ApplyMonsterPreCombatEffects(monster *Monster, player *Character) bool {
// 	switch monster.Type {
// 	case Goblin:
// 		if rand.Intn(100) < 20 {
// 			fmt.Println("The Goblin dodged your attack!")
// 			return true
// 		}
// 	case Skeleton:
// 		if rand.Intn(100) < 15 {
// 			fmt.Println("The skeleton blocked your attack!")
// 			return true
// 		}
// 	case Ork:
// 		if rand.Intn(100) < 10 {
// 			fmt.Println("The ork is enraged!")
// 			monster.Strength += 5
// 		}
// 	case Wraith:
// 		fmt.Println("The Wraith phases eerily...your armor becomes dead weight.")
// 	case Demon:
// 		burn := rand.Intn(3) + 2
// 		player.HP -= burn
// 		fmt.Printf("The Demon's fire aura burns you for %d damage!\n", burn)
// 	case Succubus:
// 		fmt.Println("The succubus seduces you and lowers your strength by 2! What a shame.")
// 		player.Strength -= 2
// 	}
// 	return false
// }
