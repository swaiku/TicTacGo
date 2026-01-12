package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// HoverMode defines how a widget visually responds to mouse hover.
type HoverMode int

const (
	// HoverFade fades in the background when hovered (alpha animation).
	HoverFade HoverMode = iota

	// HoverColorLerp interpolates between normal and hover colors.
	HoverColorLerp

	// HoverSolid switches instantly to hover color when fully hovered.
	HoverSolid
)

// Hover animation constants.
const (
	// hoverTransitionSpeed controls how fast the hover state changes per frame.
	// A value of 0.1 means 10% progress per frame (~10 frames for full transition).
	hoverTransitionSpeed = 0.1

	// hoverMin is the minimum hover value (not hovered).
	hoverMin = 0.0

	// hoverMax is the maximum hover value (fully hovered).
	hoverMax = 1.0
)

// Hoverable provides hover state management for interactive widgets.
//
// It tracks a smooth hover value (0 to 1) that transitions over time
// as the cursor enters or leaves the widget bounds. This enables
// smooth hover animations rather than instant state changes.
type Hoverable struct {
	hover float64   // Current hover state (0 = not hovered, 1 = fully hovered)
	mode  HoverMode // Animation mode for this hoverable
}

// NewHoverable creates a new Hoverable with the specified hover mode.
func NewHoverable(mode HoverMode) *Hoverable {
	return &Hoverable{
		hover: hoverMin,
		mode:  mode,
	}
}

// Update checks if the cursor is within the given bounds and updates
// the hover state accordingly.
//
// The hover value smoothly transitions toward 1 when hovered and
// toward 0 when not hovered, at a rate of hoverTransitionSpeed per frame.
func (h *Hoverable) Update(x, y, width, height float64) {
	mx, my := ebiten.CursorPosition()
	isHovered := float64(mx) >= x &&
		float64(mx) <= x+width &&
		float64(my) >= y &&
		float64(my) <= y+height

	if isHovered {
		h.hover += hoverTransitionSpeed
		if h.hover > hoverMax {
			h.hover = hoverMax
		}
	} else {
		h.hover -= hoverTransitionSpeed
		if h.hover < hoverMin {
			h.hover = hoverMin
		}
	}
}

// IsHovered returns true if the hover state is greater than zero.
//
// This can be used to check if the widget is currently being hovered
// or is in the middle of a hover transition.
func (h *Hoverable) IsHovered() bool {
	return h.hover > hoverMin
}

// ApplyHoverColor modifies the DrawImageOptions color scale based on
// the current hover state and the widget's style.
//
// The behavior depends on the HoverMode:
//   - HoverFade: applies alpha scaling based on hover progress
//   - HoverColorLerp: interpolates between normal and hover colors
//   - HoverSolid: switches to hover color only when fully hovered
func (h *Hoverable) ApplyHoverColor(op *ebiten.DrawImageOptions, style WidgetStyle) {
	switch h.mode {
	case HoverFade:
		op.ColorScale.ScaleWithColor(style.BackgroundNormal)
		op.ColorScale.ScaleAlpha(float32(h.hover))

	case HoverColorLerp:
		blendedColor := LerpColor(style.BackgroundNormal, style.BackgroundHover, h.hover)
		op.ColorScale.ScaleWithColor(blendedColor)

	case HoverSolid:
		if h.hover >= hoverMax {
			op.ColorScale.ScaleWithColor(style.BackgroundHover)
		} else {
			op.ColorScale.ScaleWithColor(style.BackgroundNormal)
		}
	}
}
