package main

import "fmt"

type Character struct {
	HP     int
	Lives  int
	Armour int
}

func (C *Character) PutShieldOn(value int) {
	C.Armour += value
}

type Enemy struct {
	Character
}

func (e *Enemy) Dodge() {

}

type Creature struct {
	Enemy
	Flies bool
}

func (c *Creature) StartFlying() {
	c.Flies = true
}

func (c *Creature) StartFlyingWhileDodgingandCarryingShield() {
	c.Flies = true
	c.Dodge()
	c.PutShieldOn(10)
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

	gargoyle := Creature{
		Enemy{
			Character{
				HP:     11,
				Lives:  1,
				Armour: 1,
			},
		},
		false,
	}
	fmt.Printf("%+v\n", gargoyle)
	gargoyle.StartFlying()
	fmt.Printf("%+v\n", gargoyle)
	gargoyle.StartFlyingWhileDodgingandCarryingShield()
	fmt.Printf("%+v\n", gargoyle)
}
