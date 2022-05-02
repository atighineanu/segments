package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	HP     int
	Lives  int
	Armour int
}

func (c *Character) TakeDmg(damage int) {
	if damage > c.Armour {
		c.HP -= damage - c.Armour
	}
}

func (C *Character) PutShieldOn(value int) {
	C.Armour += value
}

type Enemy struct {
	Character
}

func (e *Enemy) Dodge(rate int) bool {
	if rate > 100 {
		fmt.Println("Dodge rate cannot exceed 100%")
	} else {
		rand.Seed(time.Now().UnixNano())
		//fmt.Printf("%v\n", a)
		if rand.Intn(100) <= rate {
			fmt.Println("dodged")
			return true
		}
	}
	fmt.Println("didn't dodge")
	return false
}

type Creature struct {
	Enemy
	Flies     bool
	DodgeRate int
}

func (c *Creature) StartFlying() {
	c.Flies = true
}

func (c *Creature) StartFlyingWhileDodgingandCarryingShield() {
	c.Flies = true
	c.PutShieldOn(10)
	c.DodgeRate += 10
}

func (c *Creature) TakeDamage(damage int) {
	if !c.Dodge(c.DodgeRate) {
		c.TakeDmg(damage)
	}
}

func main() {
	troll := Enemy{
		Character{
			HP:     11,
			Lives:  1,
			Armour: 1,
		},
	}

	fmt.Printf("%+v\n", troll)
	troll.PutShieldOn(3)
	fmt.Printf("%+v\n", troll)
	troll.TakeDmg(7)
	fmt.Printf("%+v\n", troll)
	gargoyle := Creature{
		Enemy{
			Character{
				HP:     11,
				Lives:  1,
				Armour: 1,
			},
		},
		false,
		10,
	}
	fmt.Printf("%+v\n", gargoyle)
	gargoyle.StartFlying()
	fmt.Printf("%+v\n", gargoyle)
	gargoyle.StartFlyingWhileDodgingandCarryingShield()
	fmt.Printf("%+v\n", gargoyle)
	gargoyle.TakeDamage(20)
	fmt.Printf("%+v\n", gargoyle)
}
