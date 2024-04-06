package data

import "math/rand/v2"

type Board struct {
	ID        string   `db:"id,primarykey" form-type:"hidden" form:"id"`
	Name      string   `db:"name" form-type:"text" label:"name" form:"name"`
	Colors    string   `db:"colors" form-type:"color-list" label:"colors" form:"colors"`
	Hexes     []Hex    `db:"omit"`
	ColorList []string `db:"omit"` //may want to remove this later
}

func NewBoard() Board {
	b := Board{ID: NewUniqueID(), Name: "GameBoard", Hexes: make([]Hex, 0)}

	return b
}
func NewBoardTemplate(x int, y int) Board {
	var hexes []Hex
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			hexes = append(hexes, Hex{ID: "Template", BoardID: "Template", X: i, Y: j, Color: "#FFFFFF", Type: "placeholder"})
		}
	}
	b := Board{ID: "Template", Name: "Template", Hexes: hexes}

	return b
}

type Hex struct {
	//Discovered bool
	ID      string `db:"id,primarykey" form-type:"hidden" form:"id"`
	BoardID string `db:"boardid" form-type:"hidden" form:"boardid"`
	Color   string `db:"color" form-type:"color" form:"color" label:"color"`
	Type    string `db:"type" form-type:"text" form:"type" label:"type"`
	Value   string `db:"value" form-type:"text" form:"value" label:"value"`
	X       int    `db:"x" form-type:"number" form:"x" label:"position x"`
	Y       int    `db:"y" form-type:"number" form:"y" label:"position y"`
	//Bank       Bank
}

func NewHex(x int, y int) Hex {
	r := rand.IntN(16)
	var resourceType string
	switch r {
	case 1:
		resourceType = "wood"
	case 2:
		resourceType = "ore"
	case 3:
		resourceType = "wool"
	case 4:
		resourceType = "adder"
	default:
		resourceType = "empty"
	}
	c := "#333" // color.RGBA{R: 22, G: 22, B: 22, A: 255}

	/* grid coloring
	if x > 5 && x < 11 && y > 3 && y < 7 { //center circle
		c = color.RGBA{R: 114, G: 114, B: 160, A: 255}
	} else if (x < 5 || x > 11) && (y < 1 || y > 8) || (x < 3 || x > 12) && (y < 2 || y > 7) || (x < 1 || x > 14) && (y < 3 || y > 6) {
		c = color.RGBA{R: 180, G: 70, B: 70, A: 255}
	} else if x < 9 && y < 5 {
		c = color.RGBA{R: 135, G: 122, B: 89, A: 255}
	} else if x >= 9 && y < 5 {
		c = color.RGBA{R: 135, G: 89, B: 111, A: 255}
	} else if x < 8 && y >= 5 {
		c = color.RGBA{R: 94, G: 124, B: 129, A: 255}
	} else if x >= 8 && y >= 5 {
		c = color.RGBA{R: 82, G: 99, B: 72, A: 255}
	}
	*/
	return Hex{ID: NewUniqueID(), Type: resourceType, X: x, Y: y, Color: c}
}
