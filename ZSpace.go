package turboOcto

import (
	"gitlab.com/Pixdigit/sorted"
	"gitlab.com/Pixdigit/uniqueID"
)

//A short comment about the implementation:
//While a map would be easier to implement I go with an array approach
//This is because with a map order would have to determined each frame
//With an array only while changing order

type zSpace struct {
	sorted.Set
}

var zSpaceSingelton zSpace

type RenderElement interface {
	Render() error
	ID() uniqueID.ID
}

func init() {
	zSpaceSingelton = zSpace{}
}

func AddElement(element RenderElement, z int) error {
	return zSpaceSingelton.Insert(element, sorted.Num(z))
}

func (z *zSpace) Render() {
	for _, elem := range z.Set.Elems() {
		elem.(RenderElement).Render()
	}
}
