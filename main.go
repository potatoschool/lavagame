package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"lavagame/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	frameWidth  = 65
	frameHeight = 64

	speed = 10
)

var (
	mainAssetsImage      *ebiten.Image // 1920x1440
	characterAssetsImage *ebiten.Image
)

type Game struct {
	cloudMod float64

	stanceY int
	stanceX int

	posY int
	posX int
}

func (g *Game) Update() error {
	var _keyMoveLeft, _keyMoveRight bool

	g.cloudMod += 0.5

	pressedKeys := inpututil.PressedKeys()

	for _, k := range pressedKeys {
		if k == ebiten.KeyA {
			_keyMoveLeft = true
		} else if k == ebiten.KeyD {
			_keyMoveRight = true
		}
	}

	if _keyMoveLeft {
		g.posX -= speed

		if g.stanceY == 9 && g.stanceX == 0 {
			g.stanceX = 1
		} else {
			g.stanceX = 0
		}

		g.stanceY = 9

		if g.posX <= 0 {
			g.posX = 0
		}
	}

	if _keyMoveRight {
		g.posX += speed

		if g.stanceY == 11 && g.stanceX == 0 {
			g.stanceX = 1
		} else {
			g.stanceX = 0
		}

		g.stanceY = 11

		if g.posX >= screenWidth {
			g.posX = screenWidth
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xa5, 0x45, 0x1e, 0xff})

	//bgsprite
	bgSpriteX := []int{55, 600}
	bgSpriteY := []int{876, 1113}
	bgSpriteWidth := bgSpriteX[1] - bgSpriteX[0]
	//bgSpriteHeight := bgSpriteY[1] - bgSpriteY[0]
	background := mainAssetsImage.SubImage(image.Rect(bgSpriteX[0], bgSpriteY[0], bgSpriteX[1], bgSpriteY[1])).(*ebiten.Image)
	backgroundOptions := &ebiten.DrawImageOptions{}

	numberOfBackgroundTilesX := screenWidth / bgSpriteWidth

	for i := -1; i < numberOfBackgroundTilesX+2; i++ {
		backgroundOptions.GeoM.Reset()
		backgroundOptions.GeoM.Translate(float64(screenWidth-i*bgSpriteWidth), 0)
		screen.DrawImage(background, backgroundOptions)
	}

	// cloud
	cloudSpriteX := []int{54, 597}
	cloudSpriteY := []int{220, 485}
	cloudSpriteWidth := cloudSpriteX[1] - cloudSpriteX[0]
	//cloudSpriteHeight := cloudSpriteY[1] - cloudSpriteY[0]
	cloud := mainAssetsImage.SubImage(image.Rect(cloudSpriteX[0], cloudSpriteY[0], cloudSpriteX[1], cloudSpriteY[1])).(*ebiten.Image)
	cloudOptions := &ebiten.DrawImageOptions{}

	numberOfCloudTilesX := screenWidth / cloudSpriteWidth

	for i := -1; i < numberOfCloudTilesX+2; i++ {
		cloudOptions.GeoM.Reset()
		cloudOptions.GeoM.Translate(float64(screenWidth-i*cloudSpriteWidth)-g.cloudMod, float64(50))
		screen.DrawImage(cloud, cloudOptions)
	}

	// ground
	groundSpriteX := []int{56, 600}
	groundSpriteY := []int{1145, 1385}
	groundSpriteWidth := groundSpriteX[1] - groundSpriteX[0]
	groundSpriteHeight := groundSpriteY[1] - groundSpriteY[0]
	ground := mainAssetsImage.SubImage(image.Rect(groundSpriteX[0], groundSpriteY[0], groundSpriteX[1], groundSpriteY[1])).(*ebiten.Image)
	groundOptions := &ebiten.DrawImageOptions{}

	numberOfGroundTilesX := screenWidth / groundSpriteWidth

	for i := -1; i < numberOfGroundTilesX+2; i++ {
		groundOptions.GeoM.Reset()
		groundOptions.GeoM.Translate(float64(screenWidth-i*groundSpriteWidth), float64(screenHeight-groundSpriteHeight))
		screen.DrawImage(ground, groundOptions)
	}

	sx, sy := g.stanceX*frameWidth, g.stanceY*frameHeight
	frame := characterAssetsImage.SubImage(image.Rect(sx, sy, sx+frameHeight, sy+frameWidth)).(*ebiten.Image)
	charOptions := &ebiten.DrawImageOptions{}

	charOptions.GeoM.Translate(float64(g.posX), float64(g.posY))
	screen.DrawImage(frame, charOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	mainImg, _, err := image.Decode(bytes.NewReader(assets.Assets_png))
	if err != nil {
		panic(err)
	}
	mainAssetsImage = ebiten.NewImageFromImage(mainImg)

	charImg, _, err := image.Decode(bytes.NewReader(assets.Character_png))
	if err != nil {
		panic(err)
	}
	characterAssetsImage = ebiten.NewImageFromImage(charImg)

	newGame := &Game{0, 3, 0, 470, 30}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Lavagame")

	if err := ebiten.RunGame(newGame); err != nil {
		panic(err)
	}

}
