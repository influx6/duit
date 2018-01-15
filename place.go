package duit

import (
	"image"

	"9fans.net/go/draw"
)

type Place struct {
	Place func(self *Kid, sizeAvail image.Point) `json:"-"`
	Kids  []*Kid

	kidsReversed []*Kid
	size         image.Point
}

var _ UI = &Place{}

func (ui *Place) ensure() {
	if len(ui.kidsReversed) == len(ui.Kids) {
		return
	}
	ui.kidsReversed = make([]*Kid, len(ui.Kids))
	for i, k := range ui.Kids {
		ui.kidsReversed[len(ui.Kids)-1-i] = k
	}
}

func (ui *Place) Layout(dui *DUI, self *Kid, sizeAvail image.Point, force bool) {
	ui.ensure()
	dui.debugLayout("Place", self)

	ui.Place(self, sizeAvail)
}

func (ui *Place) Draw(dui *DUI, self *Kid, img *draw.Image, orig image.Point, m draw.Mouse, force bool) {
	// xxx place should copies of its kids images, so it doesn't have to ask them to redraw all the time
	if self.Draw == DirtyKid {
		force = true
	}
	kidsDraw("Place", dui, self, ui.Kids, ui.size, img, orig, m, force)
}

func (ui *Place) Mouse(dui *DUI, self *Kid, m draw.Mouse, origM draw.Mouse, orig image.Point) (r Result) {
	return kidsMouse(dui, self, ui.kidsReversed, m, origM, orig)
}

func (ui *Place) Key(dui *DUI, self *Kid, k rune, m draw.Mouse, orig image.Point) (r Result) {
	return kidsKey(dui, self, ui.kidsReversed, k, m, orig)
}

func (ui *Place) FirstFocus(dui *DUI) (warp *image.Point) {
	return kidsFirstFocus(dui, ui.Kids)
}

func (ui *Place) Focus(dui *DUI, o UI) (warp *image.Point) {
	return kidsFocus(dui, ui.Kids, o)
}

func (ui *Place) Mark(self *Kid, o UI, forLayout bool) (marked bool) {
	return kidsMark(self, ui.Kids, o, forLayout)
}

func (ui *Place) Print(self *Kid, indent int) {
	PrintUI("Place", self, indent)
	kidsPrint(ui.Kids, indent+1)
}