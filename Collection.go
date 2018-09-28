package turboOcto

import (
	geo "gitlab.com/Pixdigit/geometry"
)

//May also have Update Method but is no requirent
type PositionedRenderable interface {
	Renderable
	MoveTo(geo.Point)
	MoveRel(geo.Vector)
}

type Collection struct {
	Renderables []PositionedRenderable
	anchor      *geo.Point
}

func (c *Collection) MoveTo(p geo.Point) {
	vec := p.ToVector()
	c.anchor.Sub(*vec)

	relMotion := c.anchor.ToVector()

	for _, elem := range c.Renderables {
		elem.MoveRel(*relMotion)
	}

	c.anchor.MoveTo(p)
}

func (c *Collection) MoveRel(v geo.Vector) {
	c.anchor.Add(v)

	for _, elem := range c.Renderables {
		elem.MoveRel(v)
	}
}

func (c *Collection) Position() geo.Point {
	return *c.anchor
}

func (c *Collection) Render() []error {
	errs := []error{}
	for _, elem := range c.Renderables {
		err := elem.Render()
		errs = append(errs, err)
	}
	if len(errs) >= 1 {
		return errs
	}
	return nil
}
