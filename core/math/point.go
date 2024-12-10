package math

type Point[T Number] struct {
	X T
	Y T
}

func (p *Point[T]) Move(v Vector2[T], times T) {
	p.X += v.X * times
	p.Y += v.Y * times
}

func (p *Point[T]) IsWithinBounds(p2 Point[T]) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	if p.X >= p2.X || p.Y >= p2.Y {
		return false
	}

	return true
}
