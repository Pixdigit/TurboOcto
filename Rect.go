package turboOcto

import (
	"github.com/pkg/errors"
)

type xPlane int
type yPlane int

const (
	TOP = yPlane(iota)
	CENTERY
	BOTTOM
)
const (
	LEFT = xPlane(iota)
	CENTERX
	RIGHT
)

type AnchorPoint struct {
	xAnchor xPlane
	yAnchor yPlane
}

type Rect struct {
	FixPoint         AnchorPoint
	topLeftReference Point
	size             Size
}

type rectQuery struct {
	Width, Height, W, H                              int32
	Top, Left, Bottom, Right                         int32
	CenterX, CenterY                                 int32
	TopLeft, TopRight, BottomLeft, BottomRight       Point
	Center                                           Point
	CenterTop, CenterLeft, CenterBottom, CenterRight Point
}

func NewRect(p Point, s Size, fixPoint AnchorPoint) (*Rect, error) {
	//Calculate reference after size is defined
	r := &Rect{fixPoint, Point{0, 0}, s}
	p2, err := r.getTopLeftFromPoint(p);	if err != nil {return &Rect{}, errors.Wrap(err, "could not calculate rect Position")}
	r.topLeftReference = p2

	err = r.validateSize();	if err != nil {return &Rect{}, errors.Wrap(err, "unable to create Rect")}

	return r, nil
}

func (r Rect) getTopLeftFromPoint(p Point) (Point, error) {
	//Pretend this rect's anchor point is at p. Where would top left be?

	topLeft := Point{0, 0}
	switch r.FixPoint.xAnchor {
	case LEFT:
		topLeft.X = p.X
	case CENTERX:
		topLeft.X = p.X - r.size.Width/2
	case RIGHT:
		topLeft.X = p.X - r.size.Width
	}
	switch r.FixPoint.yAnchor {
	case TOP:
		topLeft.Y = p.Y
	case CENTERY:
		topLeft.Y = p.Y - r.size.Height/2
	case BOTTOM:
		topLeft.Y = p.Y - r.size.Height
	}
	return topLeft, nil
}

func (r *Rect) validateSize() error {
	ok, err := r.size.AreaIsPositive();	if err != nil {return errors.Wrap(err, "could not validate size")}
	if !ok {
		return errors.New("Rect can not have negative size")
	}
	return nil
}

func (r *Rect) Empty() (bool, error) {
	return (r.size.Width == 0 && r.size.Height == 0), nil
}

func (r *Rect) Equals(r2 Rect) (bool, error) {
	ok, err := r.topLeftReference.Equals(r2.topLeftReference);	if err != nil {return false, errors.Wrap(err, "could not check if rects are equal")}
	ok2 := (r.FixPoint.xAnchor == r2.FixPoint.xAnchor) && (r.FixPoint.yAnchor == r2.FixPoint.yAnchor)
	return (ok &&
		ok2 &&
		r.size.Width == r2.size.Width &&
		r.size.Height == r2.size.Height), err
}

func (r *Rect) HasIntersection(r2 Rect) (bool, error) {
	//Check all conditions where no collision occurs
	//Then invert
	return !(r2.topLeftReference.X > r.topLeftReference.X+r.size.Width ||
		r2.topLeftReference.X+r2.size.Width < r.topLeftReference.X ||
		r2.topLeftReference.Y > r.topLeftReference.Y+r.size.Height ||
		r2.topLeftReference.Y+r2.size.Height < r.topLeftReference.Y), nil
}
func (r *Rect) IntersectRect(r2 Rect) (Rect, bool, error) {
	ok, err := r.HasIntersection(r2);	if err != nil {return Rect{}, false, errors.Wrap(err, "Could not check for intersection")}
	if !ok {
		return Rect{}, false, nil
	}

	top, err := max(r.topLeftReference.Y, r2.topLeftReference.Y);	if err != nil {return Rect{}, false, err}
	left, err := max(r.topLeftReference.X, r2.topLeftReference.X);	if err != nil {return Rect{}, false, err}
	bottom, err := min(r.topLeftReference.Y+r.size.Height, r2.topLeftReference.Y+r2.size.Height);	if err != nil {return Rect{}, false, err}
	right, err := min(r.topLeftReference.X+r.size.Width, r2.topLeftReference.X+r2.size.Width);	if err != nil {return Rect{}, false, err}

	intersection, err := NewRect(Point{top, left}, Size{right - left, bottom - top}, AnchorPoint{LEFT, TOP});	if err != nil {return Rect{}, false, errors.Wrap(err, "could not get intersection")}

	return *intersection, true, nil
}

func (r *Rect) GetAnchorPlane(pl interface{}) (int32, error) {
	var scalar int32

	switch pl.(type) {
	case xPlane:
		switch pl {
		case LEFT:
			scalar = r.topLeftReference.X
		case CENTERX:
			scalar = r.topLeftReference.X + r.size.Width/2
		case RIGHT:
			scalar = r.topLeftReference.X + r.size.Width
		default:
			return 0, errors.New("xPlane is unknown")
		}
	case yPlane:
		switch pl {
		case TOP:
			scalar = r.topLeftReference.Y
		case CENTERY:
			scalar = r.topLeftReference.Y + r.size.Height/2
		case BOTTOM:
			scalar = r.topLeftReference.Y + r.size.Height
		default:
			return 0, errors.New("yPlane is unknown")
		}
	default:
		return 0, errors.New("plane is of unknown dimension")
	}
	return scalar, nil
}

func (r *Rect) GetAnchorPoint(fixPoint AnchorPoint) (Point, error) {
	x, err := r.GetAnchorPlane(fixPoint.xAnchor);	if err != nil {return Point{}, err}
	y, err := r.GetAnchorPlane(fixPoint.yAnchor);	if err != nil {return Point{}, err}

	return Point{x, y}, nil
}
func (r *Rect) GetPosition() (Point, error) {
	return r.GetAnchorPoint(r.FixPoint)
}

func (r *Rect) MoveTo(p Point) error {
	newTopLeft, err := r.getTopLeftFromPoint(p);	if err != nil {return errors.Wrap(err, "could not calculate new rect position")}
	r.topLeftReference = newTopLeft
	return nil
}

func (r *Rect) MoveRelative(v Vector) error {
	r.topLeftReference.Add(v)
	return nil
}

//SetSize changes the size of the Rect in such a way, that the anchor will stay in place
func (r *Rect) SetSize(s Size) error {
    anchor, err := r.GetAnchorPoint(r.FixPoint)
    if err != nil {
        return errors.Wrap(err, "could not change size")
    }
    r.size = s
    err = r.MoveTo(anchor)
    if err != nil {
        return errors.Wrap(err, "could not retain position in place after size change")
    }
    return nil
}
