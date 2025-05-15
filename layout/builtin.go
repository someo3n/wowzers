// layout is a package containing everything about layouts in wowzers
package layout

// NewTileLayout defines a layout where the working area is split into
// 2 halves horizontally and then on the first half, the masters
// get placed vertically and on the second half, the slaves get
// placed vertically as well. This is called a tiling layout
func NewTileLayout() *Layout {
	return &Layout{
		Name:                   "tile",
		RearrangeOnFocusChange: false,
		Arrange: func(ctx *Context) Element {
			//
			return NewHSplit([]Element{
				NewVSplit(ToElementSlice(ctx.Masters)),
				NewVSplit(ToElementSlice(ctx.Slaves)),
			})
		},
	}
}

// NewSpiralLayout defines a layout where the working area is split into
// 2 halves horizontally and then on the first half, the masters get placed
// vertically and on the second half, the slaves get placed in a spiraling
// layout where the most prominent one is biggest and the least prominent
// one is smallest and off to the corner
func NewSpiralLayout() *Layout {
	return &Layout{
		Name:                   "spiral",
		RearrangeOnFocusChange: false,
		Arrange: func(ctx *Context) Element {
			return NewHSplit([]Element{
				NewVSplit(ToElementSlice(ctx.Masters)),
				Spiral(ToElementSlice(ctx.Slaves), false),
			})
		},
	}
}

// NewCenteredMasterLayout defines a layout where masters is centered and
// the halves of the slaves are put on the left and the right.
// This is called a centered master layout
func NewCenteredMasterLayout() *Layout {
	return &Layout{
		Name:                   "centeredmaster",
		RearrangeOnFocusChange: false,
		Arrange: func(ctx *Context) Element {
			if len(ctx.Slaves) == 0 {
				return NewVSplit(ToElementSlice(ctx.Masters))
			}

			mid := len(ctx.Slaves) / 2 // find middle index

			leftSlaves := ctx.Slaves[:mid]  // the first half of slaves
			rightSlaves := ctx.Slaves[:mid] // the second half of slaves

			return NewHSplit([]Element{
				NewVSplit(ToElementSlice(leftSlaves)),
				NewVSplit(ToElementSlice(ctx.Masters)),
				NewVSplit(ToElementSlice(rightSlaves)),
			})
		},
	}
}
