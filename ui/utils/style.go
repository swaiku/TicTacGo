package utils

import "image/color"

// WidgetStyle defines the visual appearance of a widget.
//
// It includes colors for normal and hover states, text appearance,
// border styling, and the hover animation behavior.
type WidgetStyle struct {
	// Background colors
	BackgroundNormal color.Color // Background color in normal state
	BackgroundHover  color.Color // Background color when hovered

	// Text appearance
	TextColor color.Color // Color used for text rendering

	// Border styling
	BorderColor color.Color // Border stroke color
	BorderWidth float64     // Border stroke width in pixels

	// Interaction behavior
	HoverMode HoverMode // Animation mode when hovering
}

// Predefined widget styles for common UI patterns.
var (
	// DefaultWidgetStyle provides a neutral gray appearance with hover lerp.
	DefaultWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 40, G: 40, B: 40, A: 255},
		BackgroundHover:  color.RGBA{R: 70, G: 70, B: 70, A: 255},
		TextColor:        color.White,
		BorderColor:      color.RGBA{R: 255, G: 255, B: 255, A: 50},
		BorderWidth:      2,
		HoverMode:        HoverColorLerp,
	}

	// DisabledWidgetStyle provides a grayed-out appearance with no hover effect.
	DisabledWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 40, G: 40, B: 40, A: 255},
		BackgroundHover:  color.RGBA{R: 40, G: 40, B: 40, A: 255},
		TextColor:        color.Gray{Y: 255},
		BorderColor:      color.RGBA{R: 255, G: 255, B: 255, A: 50},
		BorderWidth:      2,
		HoverMode:        HoverSolid,
	}

	// TransparentWidgetStyle provides an invisible background with fade hover.
	TransparentWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		BackgroundHover:  color.RGBA{R: 255, G: 255, B: 255, A: 20},
		TextColor:        color.White,
		BorderColor:      color.RGBA{R: 255, G: 255, B: 255, A: 100},
		BorderWidth:      1,
		HoverMode:        HoverFade,
	}

	// NormalWidgetStyle provides a blue button appearance for primary actions.
	NormalWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 64, G: 92, B: 245, A: 100},
		BackgroundHover:  color.RGBA{R: 43, G: 61, B: 163, A: 255},
		TextColor:        color.White,
		BorderColor:      color.RGBA{R: 60, G: 0, B: 0, A: 255},
		BorderWidth:      2,
		HoverMode:        HoverColorLerp,
	}

	// DangerWidgetStyle provides a red appearance for destructive actions.
	DangerWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 200, G: 30, B: 30, A: 255},
		BackgroundHover:  color.RGBA{R: 200, G: 30, B: 30, A: 255},
		TextColor:        color.White,
		BorderColor:      color.RGBA{R: 60, G: 0, B: 0, A: 255},
		BorderWidth:      2,
		HoverMode:        HoverSolid,
	}

	// SuccessWidgetStyle provides a green appearance for positive/confirm actions.
	SuccessWidgetStyle = WidgetStyle{
		BackgroundNormal: color.RGBA{R: 40, G: 170, B: 80, A: 255},
		BackgroundHover:  color.RGBA{R: 30, G: 140, B: 60, A: 255},
		TextColor:        color.White,
		BorderColor:      color.RGBA{R: 0, G: 80, B: 0, A: 255},
		BorderWidth:      2,
		HoverMode:        HoverColorLerp,
	}
)
