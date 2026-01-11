// Package utils provides UI utility types and functions for layout,
// styling, and rendering operations.
package utils

// Anchor defines the reference point used for positioning a widget
// within its parent container or the screen.
//
// The anchor system uses a 3x3 grid of reference points:
//
//	TopLeft     TopCenter     TopRight
//	CenterLeft  Center        CenterRight
//	BottomLeft  BottomCenter  BottomRight
type Anchor int

// Anchor constants define the nine possible reference points for widget positioning.
const (
	AnchorTopLeft Anchor = iota
	AnchorTopCenter
	AnchorTopRight
	AnchorCenterLeft
	AnchorCenter
	AnchorCenterRight
	AnchorBottomLeft
	AnchorBottomCenter
	AnchorBottomRight
)

// ComputeAnchoredPosition calculates the top-left position of a widget
// given its anchor point, offset, size, and parent dimensions.
//
// Parameters:
//   - anchor: the reference point within the parent
//   - offsetX, offsetY: additional offset from the anchor point
//   - w, h: widget dimensions
//   - parentW, parentH: parent container dimensions
//
// Returns the (x, y) coordinates of the widget's top-left corner.
func ComputeAnchoredPosition(
	anchor Anchor,
	offsetX, offsetY float64,
	w, h float64,
	parentW, parentH float64,
) (float64, float64) {
	x := offsetX
	y := offsetY

	// Horizontal positioning based on anchor column
	switch anchor {
	case AnchorTopCenter, AnchorCenter, AnchorBottomCenter:
		x = parentW/2 - w/2 + offsetX
	case AnchorTopRight, AnchorCenterRight, AnchorBottomRight:
		x = parentW - w + offsetX
	}

	// Vertical positioning based on anchor row
	switch anchor {
	case AnchorCenterLeft, AnchorCenter, AnchorCenterRight:
		y = parentH/2 - h/2 + offsetY
	case AnchorBottomLeft, AnchorBottomCenter, AnchorBottomRight:
		y = parentH - h + offsetY
	}

	return x, y
}
