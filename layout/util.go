// layout is a package containing everything about layouts in wowzers
package layout

// ToElementSlice is a utility function that converts
// a slice of values where their type is compatible
// with the Element interface to a generic Element slice
func ToElementSlice[T Element](items []T) []Element {
	elements := make([]Element, len(items))
	for i, item := range items {
		elements[i] = item
	}
	return elements
}

// Spiral makes a spiral from these elements
func Spiral(elements []Element, horizontal bool) Element {
	if len(elements) == 0 {
		return nil
	}
	return buildSpiral(elements, horizontal)
}

// buildSpiral is a utility recursive function used by the Spiral function
func buildSpiral(elements []Element, horizontal bool) Element {
	// base case
	if len(elements) == 1 {
		return elements[0]
	}

	// split elements into two
	mid := len(elements) / 2

	// first half
	left := elements[:mid]
	// second half
	right := elements[mid:]

	if horizontal {
		return NewHSplit([]Element{
			buildSpiral(left, !horizontal),
			buildSpiral(right, !horizontal),
		})
	} else {
		return NewVSplit([]Element{
			buildSpiral(left, !horizontal),
			buildSpiral(right, !horizontal),
		})
	}
}
