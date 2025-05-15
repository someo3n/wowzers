// layout is a package containing everything about layouts in wowzers
package layout

import (
	"wowzers/client"
	"wowzers/util"
)

// LayoutGeometry allows you to convert a layout tree to a map of client handle to size and position.
// It also stores information about how to do such task, providing inner gaps, outer gaps, working area
// and a function to use when we need to query a clients factor
type LayoutGeometry struct {
	// WorkingArea stores a Rect with the working area
	WorkingArea util.Rect

	// GapsInner are the inner gaps
	GapsInner util.Point

	// GapsOuter are the outer gaps
	GapsOuter util.Point

	// ClientFactor is a function called when computing the rects of the clients and client factor is needed
	ClientFactor func(handle client.Handle) float32
}

type GeometryMap = map[client.Handle]util.Rect

func (c *LayoutGeometry) applyOuterGaps(area util.Rect) util.Rect {
	return util.Rect{
		Pos: util.Point{
			X: area.Pos.X + c.GapsOuter.X,
			Y: area.Pos.Y + c.GapsOuter.Y,
		},
		Size: util.Point{
			X: area.Size.X - 2*c.GapsOuter.X,
			Y: area.Size.Y - 2*c.GapsOuter.Y,
		},
	}
}

// Compute computes the map of client handles to sizes and positions
func (c *LayoutGeometry) Compute(root Element) GeometryMap {
	result := make(map[client.Handle]util.Rect)
	adjustedArea := c.applyOuterGaps(c.WorkingArea)
	c.compute(root, adjustedArea, result)
	return result
}

func (c *LayoutGeometry) compute(element Element, area util.Rect, result GeometryMap) {
	switch e := element.(type) {
	case *Client:
		result[e.Handle] = area
	case *Collection:
		switch e.CollectionType {
		case OverlayType:
			for _, i := range e.Children {
				c.compute(i, area, result)
			}
		case VSplitType, HSplitType:
			c.computeSplit(e, area, result, e.CollectionType == HSplitType)
		}
	default:
		panic("unknown element type")
	}
}

func (c *LayoutGeometry) computeSplit(element Element, area util.Rect, result GeometryMap, horizontal bool) {

}
