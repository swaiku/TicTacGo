package ui

import (
	"GoTicTacToe/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	Label   string
	OffsetX float64
	OffsetY float64
	Width   float64
	Height  float64

	hover   float64
	onClick func()

	image  *ebiten.Image
	Anchor Anchor

	Style WidgetStyle
}

func NewButton(
	label string,
	offsetX, offsetY float64,
	anchor Anchor,
	width, height float64,
	radius float64,
	style WidgetStyle,
	onClick func(),
) *Button {

	// Background normal
	bg := CreateRoundedRect(
		int(width), int(height),
		radius, style.BackgroundNormal,
	)

	return &Button{
		Label:   label,
		OffsetX: offsetX,
		OffsetY: offsetY,
		Width:   width,
		Height:  height,
		Anchor:  anchor,
		onClick: onClick,
		hover:   0,
		image:   bg,
		Style:   style,
	}
}

func (b *Button) AbsPosition() (float64, float64) {
	sw, sh := ebiten.WindowSize()
	return ComputeAnchoredPosition(
		b.Anchor,
		b.OffsetX, b.OffsetY,
		b.Width, b.Height,
		sw, sh,
	)
}

func (b *Button) Update() {
	mx, my := ebiten.CursorPosition()

	x, y := b.AbsPosition()

	hover := float64(mx) >= x &&
		float64(mx) <= x+b.Width &&
		float64(my) >= y &&
		float64(my) <= y+b.Height

	// Smooth hover animation
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

	// Handle click
	if hover && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.onClick != nil {
			b.onClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	x, y := b.AbsPosition()

	// === Background with hover style ===
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	ApplyHoverColor(op, b.Style, b.hover)
	screen.DrawImage(b.image, op)

	// === Border (independent image) ===
	// if b.Style.BorderWidth > 0 {
	// 	border := CreateBorderRect(
	// 		int(b.Width), int(b.Height),
	// 		b.Style.BorderWidth, b.Style.BorderColor,
	// 	)
	// 	op2 := &ebiten.DrawImageOptions{}
	// 	op2.GeoM.Translate(x, y)
	// 	screen.DrawImage(border, op2)
	// }

	// === Text ===
	opts := &text.DrawOptions{}
	opts.PrimaryAlign = text.AlignCenter
	opts.SecondaryAlign = text.AlignCenter
	opts.GeoM.Translate(x+b.Width/2, y+b.Height/2)
	opts.ColorScale.ScaleWithColor(b.Style.TextColor)

	text.Draw(screen, b.Label, assets.NormalFont, opts)
}
