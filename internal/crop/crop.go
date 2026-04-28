package crop

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func CropPNGToAlphaBounds(inPath, outPath string, keepEmpty bool) (bool, image.Rectangle, image.Rectangle, error) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return false, image.Rectangle{}, image.Rectangle{}, fmt.Errorf("open input: %w", err)
	}
	defer inFile.Close()

	srcImg, err := png.Decode(inFile)
	if err != nil {
		return false, image.Rectangle{}, image.Rectangle{}, fmt.Errorf("decode png: %w", err)
	}

	srcBounds := srcImg.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, srcBounds.Dx(), srcBounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), srcImg, srcBounds.Min, draw.Src)

	cropRect, empty := findAlphaBounds(rgba)

	var outImg *image.RGBA

	if empty {
		if !keepEmpty {
			return true, srcBounds, image.Rectangle{}, nil
		}

		outImg = image.NewRGBA(image.Rect(0, 0, 1, 1))
	} else {
		outImg = image.NewRGBA(image.Rect(0, 0, cropRect.Dx(), cropRect.Dy()))
		draw.Draw(outImg, outImg.Bounds(), rgba, cropRect.Min, draw.Src)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return empty, srcBounds, cropRect, fmt.Errorf("create output: %w", err)
	}

	encodeErr := png.Encode(outFile, outImg)
	closeErr := outFile.Close()

	if encodeErr != nil {
		return empty, srcBounds, cropRect, fmt.Errorf("encode png: %w", encodeErr)
	}
	if closeErr != nil {
		return empty, srcBounds, cropRect, fmt.Errorf("close output: %w", closeErr)
	}

	return empty, srcBounds, outImg.Bounds(), nil
}

func findAlphaBounds(img *image.RGBA) (image.Rectangle, bool) {
	bounds := img.Bounds()

	minX := bounds.Max.X
	minY := bounds.Max.Y
	maxX := bounds.Min.X - 1
	maxY := bounds.Min.Y - 1

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			alpha := img.RGBAAt(x, y).A

			if alpha > 0 {
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	if maxX < minX || maxY < minY {
		return image.Rectangle{}, true
	}

	return image.Rect(minX, minY, maxX+1, maxY+1), false
}
