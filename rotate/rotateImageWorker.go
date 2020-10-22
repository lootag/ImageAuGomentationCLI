package rotate

import (
	"image"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func rotateImageWorker(imageToRotate entities.ImageInfo,
	wg *sync.WaitGroup,
	direction entities.Direction,
	rotatedImages chan entities.ImageInfo) {

	defer (*wg).Done()
	if direction == entities.LEFT {
		var left entities.ImageInfo
		left.ImageSource = rotateImageLeft(imageToRotate.ImageSource)
		left.OriginalFileName = imageToRotate.OriginalFileName
		left.NewName = "rotateleft" + imageToRotate.OriginalFileName
		rotatedImages <- left
	} else if direction == entities.RIGHT {
		var right entities.ImageInfo
		right.ImageSource = rotateImageRight(imageToRotate.ImageSource)
		right.OriginalFileName = imageToRotate.OriginalFileName
		right.NewName = "rotateright" + imageToRotate.OriginalFileName
		rotatedImages <- right
	} else if direction == entities.FLIP {
		var flipped entities.ImageInfo
		flipped.ImageSource = flipImage(imageToRotate.ImageSource)
		flipped.OriginalFileName = imageToRotate.OriginalFileName
		flipped.NewName = "flipped" + imageToRotate.OriginalFileName
		rotatedImages <- flipped
	} else if direction == entities.ALL {
		var left entities.ImageInfo
		left.ImageSource = rotateImageLeft(imageToRotate.ImageSource)
		left.OriginalFileName = imageToRotate.OriginalFileName
		left.NewName = "rotateleft" + imageToRotate.OriginalFileName
		rotatedImages <- left
		var right entities.ImageInfo
		right.ImageSource = rotateImageRight(imageToRotate.ImageSource)
		right.OriginalFileName = imageToRotate.OriginalFileName
		right.NewName = "rotateright" + imageToRotate.OriginalFileName
		rotatedImages <- right
		var flipped entities.ImageInfo
		flipped.ImageSource = flipImage(imageToRotate.ImageSource)
		flipped.OriginalFileName = imageToRotate.OriginalFileName
		flipped.NewName = "flipped" + imageToRotate.OriginalFileName
		rotatedImages <- flipped
	} else {
		panic("The argument provided in the rotate field is not valid. Exiting.")
	}

}

func rotateImageLeft(imageToRotate image.Image) image.Image {
	rotatedImage := imaging.Rotate90(imageToRotate)
	return rotatedImage
}

func rotateImageRight(imageToRotate image.Image) image.Image {
	rotatedImage := imaging.Rotate270(imageToRotate)
	return rotatedImage
}

func flipImage(imageToRotate image.Image) image.Image {
	rotatedImage := imaging.Rotate180(imageToRotate)
	return rotatedImage
}
