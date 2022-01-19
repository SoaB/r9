package es

import "r9/sb"

type Bloop struct {
	Dna      *DNA
	Location sb.Vec2
	Health   float64
	Xoffset  float64
	Yoffset  float64
	R        float64
	MaxSpeed float64
}

func NewBloop(l sb.Vec2, dna_ *DNA) *Bloop {
	bl := new(Bloop)
	bl.Dna = dna_
	bl.Location = l.Clone()
	bl.Health = 200
	bl.Xoffset = sb.RandFloat64n(1000)
	bl.Yoffset = sb.RandFloat64n(1000)
	// Gene 0 determines maxspeed and r
	// The bigger the bloop,the slower it is
	bl.MaxSpeed = sb.Remap(bl.Dna.Genes[0], 0, 1, 15, 0)
	bl.R = sb.Remap(bl.Dna.Genes[0], 0, 1, 0, 25)
	return bl
}

func (b *Bloop) Eat(f *Foods) {
	for i := 0; i < len(f.Pos); i++ {
		foodLoc := f.Pos[i].Clone()
		distance := b.Location.Distance(foodLoc)
		if distance < b.R/2 {
			b.Health = b.Health + 100
			f.Remove(i)
			i--
		}
	}
}

func (b *Bloop) Reproduce() *Bloop {
	if sb.RandFloat64() < 0.0007 {
		childDna := b.Dna.Clone()
		childDna.Mutate(0.01)
		return NewBloop(b.Location, childDna)
	}
	return nil
}

func (b *Bloop) Update() {
	// Simple movement based on perlin noise
	vx := sb.Remap(sb.Noise1(b.Xoffset), -1, 1, -b.MaxSpeed, b.MaxSpeed)
	vy := sb.Remap(sb.Noise1(b.Yoffset), -1, 1, -b.MaxSpeed, b.MaxSpeed)
	velocity := sb.NewVec2(int(vx), int(vy))
	b.Xoffset = b.Xoffset + 0.01
	b.Yoffset = b.Yoffset + 0.01
	b.Location = b.Location.Add(velocity)
	// Death always looming
	b.Health = b.Health - 0.2
}

func (b *Bloop) Borders() {
	if b.Location.X < -b.R {
		b.Location.X = Width + b.R
	}
	if b.Location.Y < -b.R {
		b.Location.Y = Height + b.R
	}
	if b.Location.X > Width+b.R {
		b.Location.X = -b.R
	}
	if b.Location.Y > Height+b.R {
		b.Location.Y = -b.R
	}
}

func (b *Bloop) IsDead() bool {
	if b.Health < 0.0 {
		return true
	}
	return false
}

func (b *Bloop) Draw(canvas *sb.Canvas) {
	//draw function
	//col := sb.NewRgba(255, 0, 0, int(b.Health))
	canvas.DrawDisk(int(b.Location.X), int(b.Location.Y), int(b.R), sb.ColorRed)
}

func (b *Bloop) Run(canvas *sb.Canvas) {
	b.Update()
	b.Borders()
	b.Draw(canvas)
}
