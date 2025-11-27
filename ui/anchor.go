package ui

type Anchor int

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

// ComputeAnchoredPosition computes the position of an anchored widget.
// returns the x and y coordinates of the widget's top-left corner.
func ComputeAnchoredPosition(
	anchor Anchor,
	offsetX, offsetY float64,
	w, h float64,
	screenW, screenH int,
) (float64, float64) {

	x := offsetX
	y := offsetY

	switch anchor {
	case AnchorTopCenter:
		x = float64(screenW)/2 - w/2 + offsetX
	case AnchorTopRight:
		x = float64(screenW) - w + offsetX

	case AnchorCenterLeft:
		y = float64(screenH)/2 - h/2 + offsetY
	case AnchorCenter:
		x = float64(screenW)/2 - w/2 + offsetX
		y = float64(screenH)/2 - h/2 + offsetY
	case AnchorCenterRight:
		x = float64(screenW) - w + offsetX
		y = float64(screenH)/2 - h/2 + offsetY

	case AnchorBottomLeft:
		y = float64(screenH) - h + offsetY
	case AnchorBottomCenter:
		x = float64(screenW)/2 - w/2 + offsetX
		y = float64(screenH) - h + offsetY
	case AnchorBottomRight:
		x = float64(screenW) - w + offsetX
		y = float64(screenH) - h + offsetY
	}

	return x, y
}
