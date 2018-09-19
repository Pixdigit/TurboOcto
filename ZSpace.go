package turboOcto

import (
	"github.com/pkg/errors"
    "gitlab.com/Pixdigit/uniqueID"
    "gitlab.com/Pixdigit/sortedList"
)

//A short comment about the implementation:
//While a map would be easier to implement I go with an array approach
//This is because with a map order would have to determined each frame
//With an array only while changing order

type zSpace struct {
	sortedList.
	elements []*zSpaceElement
}

var ZSpace zSpace

type zSpaceElement struct {
	zValue int
	RenderElement
}
type RenderElement interface {
	Render() error
	ElementID() uniqueID.ID
}

func init() {
	ZSpace = zSpace{}
}

func AddElement(element RenderElement, z int) error {
	if ZSpace.Contains(element) {
		return errors.New("element already exists within ZSpace")
	}

	newElement := zSpaceElement{
		z,
		element,
	}

SEARCH:
	for i, existingElement := range ZSpace.elements {
		if existingElement.zValue > newElement.zValue {
			previousSpace := ZSpace.elements[:i]
			forwardSpace := make([]*zSpaceElement, len(ZSpace.elements[i:]))
			copy(forwardSpace, ZSpace.elements[i:])

			ZSpace.elements = append(previousSpace, &newElement)
			ZSpace.elements = append(ZSpace.elements, forwardSpace...)
			break SEARCH
		}
	}
	return nil
}

func (z *zSpace) Contains(element RenderElement) bool {
	id := element.ElementID()
	for _, existingElement := range z.elements {
		if existingElement.ElementID() == id {
			//prematurely exit since element was found
			return true
		}
	}
	return false
}

func (z *zSpace) Render()
