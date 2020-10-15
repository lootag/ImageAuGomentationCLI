package entities

import (
	"image"
)

type ImageInfo struct {
	ImageSource      image.Image
	OriginalFileName string
	NewName          string
}
