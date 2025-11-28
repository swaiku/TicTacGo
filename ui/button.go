package ui

import (
	"GoTicTacToe/assets"
	"GoTicTacToe/ui/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Button represents a UI button with a label, hover animation, and click handler.
// It embeds Widget, so all Widget fields and methods are promoted and accessible directly.
type Button struct {
	Widget                // Embedded base widget
	Label    string       // Text displayed on the button
	hover    float64      // Hover animation parameter (0–1)
	onClick  func()       // Callback executed when the button is clicked
}

// NewButton creates and returns a new Button instance.
func NewButton(
	label string,
	offsetX, offsetY float64,
	anchor utils.Anchor,
	width, height float64,
	radius float64,
	style utils.WidgetStyle,
	onClick func(),
) *Button {

	// Pre-render the button background (rounded rectangle)
	bg := utils.CreateRoundedRect(
		int(width), int(height),
		radius, style.BackgroundNormal,
	)

	return &Button{
		Widget: Widget{
			OffsetX: offsetX,
			OffsetY: offsetY,
			Width:   width,
			Height:  height,
			Anchor:  anchor,
			image:   bg, // internal widget image
			Style:   style,
		},
		Label:   label,
		onClick: onClick,
		hover:   0,
	}
}

// Update handles hover detection, hover animation, and click events.
func (b *Button) Update() {
	mx, my := ebiten.CursorPosition()
	x, y := b.AbsPosition()

	// Hover detection: check if cursor is inside the button bounds
	hover := float64(mx) >= x &&
		float64(mx) <= x+b.Width &&
		float64(my) >= y &&
		float64(my) <= y+b.Height

	// Smooth hover animation (0 → not hovered, 1 → fully hovered)
	if hover {
		b.hover += 0.1
		if b.hover > 1 {
			b.hover = 1
		}
	} else {
		b.hover -= 0.1
		if b.hover < 0 {
			b.hover = 0
		}
	}

	// Handle click events
	if hover && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.onClick != nil {
			b.onClick()
		}
	}
}

// Draw renders the button background and its centered label.
func (b *Button) Draw(screen *ebiten.Image) {
	x, y := b.AbsPosition()

	// --- Draw background ---
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	// Apply interactive hover color effect
	utils.ApplyHoverColor(op, b.Style, b.hover)

	screen.DrawImage(b.image, op)

	// --- Draw text ---
	opts := &text.DrawOptions{}
	opts.PrimaryAlign = text.AlignCenter
	opts.SecondaryAlign = text.AlignCenter

	// Center the text inside the button
	opts.GeoM.Translate(x+b.Width/2, y+b.Height/2)
	opts.ColorScale.ScaleWithColor(b.Style.TextColor)

	text.Draw(screen, b.Label, assets.NormalFont, opts)
}
