package es

import "r9/sb"

type Foods struct {
	Pos []sb.Vec2
}

func NewFoods(num int) *Foods {
	fds := make([]sb.Vec2, 0)
	for i := 0; i < num; i++ {
		fds = append(fds, sb.Vec2{sb.RandFloat64n(Width - 8), sb.RandFloat64n(Height - 8)})
	}
	return &Foods{fds}
}

func (f *Foods) Remove(pos int) {
	l := len(f.Pos)
	// Remove the element at index from Food
	f.Pos[pos] = f.Pos[l-1] // Copy last element to index pos
	f.Pos[l-1] = sb.Vec2{}  // Erase last element
	f.Pos = f.Pos[:l-1]     // Truncate slice
}

func (f *Foods) Add(loc sb.Vec2) {
	f.Pos = append(f.Pos, loc)
}

func (f *Foods) Draw() {
	//draw ......
}

func (f *Foods) Run(canvas *sb.Canvas) {
	for i := 0; i < len(f.Pos); i++ {
		canvas.DrawRect(int(f.Pos[i].X), int(f.Pos[i].Y), 8, 8, sb.ColorGreen)
	}
	// There's a small chance food will appear randomly.
	if sb.RandFloat64() < 0.001 {
		f.Add(sb.Vec2{sb.RandFloat64n(Width - 8), sb.RandFloat64n(Height - 8)})
	}
}
