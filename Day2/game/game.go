package main

import "fmt"

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		// X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	p1.Move(400, 600)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k: ", k)

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Crystal)
	fmt.Println(p1.Keys)
}

// implementing the fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	// if nothing else, return what has been given
	return fmt.Sprintf("<key> %d>", k)
}

// go's version of enum
const (
	// iota is a constanting increasing sequence in Go
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal
)

type Key byte

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
	// Move(int int)
}

type Player struct {
	Name string
	// X	int
	Item // Embed Item
	Keys []Key
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key %v", k)
	}
	if !containKeys(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containKeys(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

// i is called 'the receiver'
// if you want to mutate, use the pointer receiver - * is the key here
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
// captitals are exportable
type Item struct {
	X int
	Y int
}

//learning pointers
// userTickets := 50

// fmt.Println("tickets %v", userTickets)
