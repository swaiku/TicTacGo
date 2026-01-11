package utils

// SizeMode describes how a widget resolves its dimensions relative to its parent.
type SizeMode int

const (
	// SizeFixed uses the widget's explicit Width/Height values.
	SizeFixed SizeMode = iota

	// SizeFill expands the widget to fill the parent's available space on that axis.
	SizeFill
)

// LayoutRect represents a resolved rectangle in absolute screen coordinates.
//
// This is the result of computing a widget's final position and size
// after applying anchoring, offsets, and size modes.
type LayoutRect struct {
	X, Y          float64 // Top-left corner position
	Width, Height float64 // Dimensions
}

// Insets represents spacing values for all four sides of a rectangle,
// similar to CSS padding or margin.
type Insets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// InsetsAll creates uniform insets with the same value on all sides.
func InsetsAll(v float64) Insets {
	return Insets{Top: v, Right: v, Bottom: v, Left: v}
}

// Horizontal returns the total horizontal space (left + right).
func (i Insets) Horizontal() float64 {
	return i.Left + i.Right
}

// Vertical returns the total vertical space (top + bottom).
func (i Insets) Vertical() float64 {
	return i.Top + i.Bottom
}
