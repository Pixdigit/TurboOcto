package turboOcto

import (
	"testing"

	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestRectCreation(t *testing.T) {
	_, err := NewRect(Point{0, 0}, Size{50, 50}, AnchorPoint{LEFT, TOP})
	if err != nil {
		tools.WrapErr(err, "Can not create new Rect", t)
	}
}

func TestRectMovement(t *testing.T) {
	r, _ := NewRect(Point{0, 0}, Size{50, 50}, AnchorPoint{LEFT, TOP})

	r.FixPoint = AnchorPoint{RIGHT, BOTTOM}
	err := r.MoveTo(Point{25, 25})
	if err != nil {
		tools.WrapErr(err, "could not move Rect to Point", t)
	}
	err = r.MoveRelative(Vector{-25, -25})
	if err != nil {
		tools.WrapErr(err, "could not move Rect relatively", t)
	}

	expectedAnchors := map[AnchorPoint]Point{
		AnchorPoint{LEFT, TOP}:        Point{-50, -50},
		AnchorPoint{CENTERX, TOP}:     Point{-25, -50},
		AnchorPoint{RIGHT, TOP}:       Point{0, -50},
		AnchorPoint{LEFT, CENTERY}:    Point{-50, -25},
		AnchorPoint{CENTERX, CENTERY}: Point{-25, -25},
		AnchorPoint{RIGHT, CENTERY}:   Point{0, -25},
		AnchorPoint{LEFT, BOTTOM}:     Point{-50, 0},
		AnchorPoint{CENTERX, BOTTOM}:  Point{-25, 0},
		AnchorPoint{RIGHT, BOTTOM}:    Point{0, 0},
	}

	for aP, pTest := range expectedAnchors {
		p, err := r.GetAnchorPoint(aP)
		if err != nil {
			tools.WrapErr(err, "could not get anchor point", t)
		}
		ok, err := p.Equals(pTest)
		if err != nil {
			tools.WrapErr(err, "could not check if points are equal", t)
		}
		tools.Test(ok, "anchor point is not at expected Position", t)
	}
}

func TestRectIntersection(t *testing.T) {
	r1, _ := NewRect(Point{0, 0}, Size{50, 50}, AnchorPoint{LEFT, TOP})
	r2, _ := NewRect(Point{25, 25}, Size{50, 50}, AnchorPoint{LEFT, TOP})
	intersection1, ok, err := r1.IntersectRect(*r2)
	if err != nil {
		tools.WrapErr(err, "could not calculate intersection", t)
	}
	tools.Test(ok, "could not calculate intersection where one exists", t)
	intersection2, ok, err := r2.IntersectRect(*r1)
	if err != nil {
		tools.WrapErr(err, "could not calculate intersection", t)
	}
	tools.Test(ok, "could not calculate intersection where one exists", t)

	ok, err = intersection1.Equals(intersection2)
	if err != nil {
		tools.WrapErr(err, "could not check if intersections are equal", t)
	}
	tools.Test(ok, "intersections of two rects are not identical depending of order", t)

}
func TestRectSizeChange(t *testing.T) {
	r, _ := NewRect(Point{0, 0}, Size{50, 50}, AnchorPoint{LEFT, TOP})
	r.FixPoint = AnchorPoint{RIGHT, BOTTOM}
	r.SetSize(Size{25, 25})
	testPoint, _ := r.GetAnchorPoint(AnchorPoint{LEFT, TOP})
	ok, _ := testPoint.Equals(Point{25, 25})
	tools.Test(ok, "size did not change with position remaining in place", t)
}
