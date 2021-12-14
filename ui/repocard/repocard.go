package repocard

import (
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	cardWidth  = 350
	cardHeight = 200
)

var (
	cards []*card
	green = color.RGBA{0, 255, 0, 255}
	gray  = color.RGBA{200, 200, 200, 255}
)

type card struct {
	name    string
	desc    string
	release string
	units   []int
	content *fyne.Container
}

func Add(name, desc string) *card {
	c := card{name: name, desc: desc}
	cards = append(cards, &c)
	return &c
}

func (c *card) Release(info string) {
	c.release = info
}

func (c *card) MakeCard() {
	// text := canvas.NewText("Unit-3. 3 commits after", color.RGBA{0, 255, 0, 255})
	// text.Alignment = fyne.TextAlignTrailing
	// text.TextStyle = fyne.TextStyle{Italic: true}
	desc := canvas.NewText(c.desc, gray)
	var s strings.Builder
	for _, v := range c.units {
		s.WriteString("Unit-" + strconv.Itoa(v) + "; ")
	}
	units := canvas.NewText(s.String(), green)
	units.Move(fyne.NewPos(20, 20))
	c.content = container.NewWithoutLayout(desc, units)
}

func MakeGrid() *fyne.Container {
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(cardWidth, cardHeight)))
	for _, c := range cards {
		r := widget.NewCard(c.name, c.desc, c.content)
		grid.Add(r)
	}
	return grid
}
