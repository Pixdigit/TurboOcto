package turboOcto

import (
	"gitlab.com/Pixdigit/sorted"
	"gitlab.com/Pixdigit/uniqueID"
)

var zSpace sorted.Set

type UniqueRenderable interface {
	Renderable
	ID() uniqueID.ID
}

func init() {
	zSpace = sorted.Set{}
}

//Adds an element to the zSpace.
func AddElement(element UniqueRenderable, z float64) error {
	//Make sure all elements fulfill RenderElement for the type asserstion later
	return zSpace.Insert(element, sorted.Num(z))
}
