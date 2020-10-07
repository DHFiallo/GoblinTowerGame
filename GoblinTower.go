//Darius Fiallo
//June 17, 2020
//Goblin Tower (roguelike game)

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type human struct {
  name string
	hp      int
	atk     int
	def     int
	potions [5]int
	gold    int
}

type goblin struct {
	hp  int
	atk int
	def int
}

func generateHero(gold int) human {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	var h human
  fmt.Println("Player, choose your name.")
  var user string
  fmt.Scanln(&user)
  time.Sleep(2 * time.Second)
  fmt.Println(user + " you embark on a great and perilous journey. Good luck!")
  time.Sleep(2 * time.Second)


	potions := [5]int{2, 2, 2, 2, 2}
  h.name = user
	h.hp = gen.Intn(11) + 20
	h.atk = gen.Intn(3) + 1
	h.def = gen.Intn(5) + 1
	h.potions = potions
	h.gold = gold

  fmt.Println("Generating stats...")
  time.Sleep(3 * time.Second)

  var health string = "Health: " + strconv.Itoa(h.hp)
  var attack string = "Attack: " + strconv.Itoa(h.atk)
  var defense string = "Defense: " + strconv.Itoa(h.def)
  var stats string = "Name: " + user + "\n" + health + "\n" + attack + "\n" + defense

  fmt.Println(stats)

	return h //Returns the hero generated character to the variable that called this function
}

func generateGoblin() goblin {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	var g goblin

	g.hp = gen.Intn(6) + 5
	g.atk = gen.Intn(2) + 2
	g.def = gen.Intn(2) + 1

	return g
}

//defense goes away when hit, reset when level up
func (h *human) fight(g *goblin) bool {

	for h.hp > 0 && g.hp > 0 {
		var dmg string
		var def int = 0
		var hp int = 0
		for i := 0; i < h.atk; i++ {

			if g.def > 0 {
				g.def--
				def++
			} else if g.hp > 0 {
				g.hp--
				hp++
			}

		}
		dmg = dmg + strconv.Itoa(def) + " defense and " + strconv.Itoa(hp) + " hp."
		fmt.Println("You hit the goblin for ", dmg)
		fmt.Println("It has ", strconv.Itoa(g.def)+" defense and "+strconv.Itoa(g.hp)+" health.")
		if g.hp > 0 {
			def = 0
			hp = 0
			for i := 0; i < g.atk; i++ {
				dmg = ""
				if h.def > 0 {
					h.def--
					def++
				} else if h.hp > 0 {
					h.hp--
					hp++
				}
			}
			dmg = dmg + strconv.Itoa(def) + " defense and " + strconv.Itoa(hp) + " hp."
			fmt.Println("The goblin hit you for ", dmg)
			fmt.Println("You have ", strconv.Itoa(h.def)+" defense and "+strconv.Itoa(h.hp)+" health.")

		}
	}
	if g.hp > 0 {
		fmt.Println("The goblin wins! You are dead.")
		return false
	} else {
		fmt.Println("You win, the goblin is dead!")
		h.gold = h.gold + 2
		fmt.Println("Your current health is ", strconv.Itoa(h.hp))
		return true
	}
}

func (h *human) steps() {
	defense := h.def
	health := h.hp

	steps := 0
	level := 1
	goblins := 0

	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	for h.hp > 0 {

		if health >= h.hp+2 {
			h.usePotions()
		}
		steps++
		if steps%10 == 0 {
			fmt.Println("Level up!")
			fmt.Println("Step #", strconv.Itoa(steps))
			var user string
			fmt.Println("Would you like to visit the potion shop? Y/N")
			fmt.Scanln(&user)
			if user == "Y" {
				h.potionShop()
			}
			level++
			fmt.Println("You have regained your defense.")
			h.def = defense
			fmt.Println("You are back at ", strconv.Itoa(h.def), " defense.")
		}
		chance := gen.Intn(5) + 1
		if chance == 5 {
			fmt.Println("A wild goblin approaches! Get ready to fight!")
			g := generateGoblin()
			if h.fight(&g) {
				goblins++
			} else {
				break
			}
		}
	}
	fmt.Println("Your level was ", strconv.Itoa(level))
	fmt.Println("You killed ", strconv.Itoa(goblins), ".")
}

func (h *human) usePotions() {
	for i := 0; i < len(h.potions); i++ {
		if h.potions[i] == 2 {
			h.potions[i] = 0
			h.hp = h.hp + 2
			fmt.Println("Potion has been used.")
			fmt.Println("Health is now at ", strconv.Itoa(h.hp))

			break
		}
	}
}

func (h *human) potionShop() {
	fmt.Println("Potions cost 4g each. You have ", strconv.Itoa(h.gold))
	fmt.Println("Here is your potion belt.")
	fmt.Println(h.potions)
	for i := 0; i < len(h.potions); i++ {
		if h.potions[i] == 0 {
			var user string
			fmt.Println("Buy a potion? Y/N")
			fmt.Scanln(&user)
			if user == "Y" {
				if h.gold >= 4 {
					h.potions[i] = 2
					h.gold = h.gold - 4
					fmt.Println("This is how many you got now.")
					fmt.Println(h.potions)
				} else {
					fmt.Println("If you ain't got the dough, no po for yo")
					break
				}
			} else {
				fmt.Println("Stop wasting my time.")
				break
			}
		}
	}
}
func main() {

  user := "Y"
  fmt.Println("Would you like to play Goblin Tower? Y/N")
  fmt.Scanln(&user)
	
	
	for user == "Y" {
    h := generateHero(0)
		h.steps()      //Runs the game
		fmt.Println(h) //Prints out player stats
		fmt.Println("You want to play again? Y/N")
		fmt.Scanln(&user)
		if user == "Y" {
			h = generateHero(h.gold)
		} else {
			user = "N"
			fmt.Println("Thank you for playing Darius's Dungeon")
			break
		}
	}
}
