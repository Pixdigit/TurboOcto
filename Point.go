package turboOcto

type Point struct {
	X, Y int32
}
type Vector Point

func (p *Point) Add(v Vector) error {
	p.X += v.X
	p.Y += v.Y
	return nil
}
func (p *Point) Sub(v Vector) error {
	p.X -= v.X
	p.Y -= v.Y
	return nil
}
func (p *Point) MoveTo(p2 Point) error {
	p.X = p2.X
	p.Y = p2.Y
	return nil
}
func (p *Point) InRect(r Rect) (bool, error) {
	return (r.topLeftReference.X <= p.X &&
		p.X <= r.topLeftReference.X &&
		r.topLeftReference.Y <= p.Y &&
		p.Y <= r.topLeftReference.Y), nil
}
func (p *Point) Equals(p2 Point) (bool, error) {
	return (p.X == p2.X && p.Y == p2.Y), nil
}
