package rotate

import (
	"github.com/disintegration/imaging"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"sync"
)

func rotateImageWorker(imageToRotate entities.ImageInfo,
	wg *sync.WaitGroup,
	direction entities.Direction,
	rotated chan entities.ImageInfo) {

	defer (*wg).Done()
	if direction == entities.LEFT {
		var left entities.ImageInfo
		left.ImageSource = rotateImageLeft(imageToRotate.ImageSource)
		left.OriginalFileName = imageToRotate.OriginalFileName
		left.NewName = "rotateleft" + imageToRotate.OriginalFileName
		rotated <- left
	} else if direction == entities.RIGHT {
		var right entities.ImageInfo
		right.ImageSource = rotateImageRight(imageToRotate.ImageSource)
		right.OriginalFileName = imageToRotate.OriginalFileName
		right.NewName = "rotateright" + imageToRotate.OriginalFileName
		rotated <- right
	} else if direction == entities.FLIP {
		var flipped entities.ImageInfo
		flipped.ImageSource = flipImage(imageToRotate.ImageSource)
		flipped.OriginalFileName = imageToRotate.OriginalFileName
		flipped.NewName = "flipped" + imageToRotate.OriginalFileName
		rotated <- flipped
	} else if direction == entities.ALL {
		var left entities.ImageInfo
		left.ImageSource = rotateImageLeft(imageToRotate.ImageSource)
		left.OriginalFileName = imageToRotate.OriginalFileName
		left.NewName = "rotateleft" + imageToRotate.OriginalFileName
		rotated <- left
		var right entities.ImageInfo
		right.ImageSource = rotateImageRight(imageToRotate.ImageSource)
		right.OriginalFileName = imageToRotate.OriginalFileName
		right.NewName = "rotateright" + imageToRotate.OriginalFileName
		rotated <- right
		var flipped entities.ImageInfo
		flipped.ImageSource = flipImage(imageToRotate.ImageSource)
		flipped.OriginalFileName = imageToRotate.OriginalFileName
		flipped.NewName = "flipped" + imageToRotate.OriginalFileName
		rotated <- flipped
	} else {
		panic("The argument provided in the rotate field is not valid. Exiting.")
	}

}

func rotateImageLeft(imageToRotate image.Image) image.Image {
	dstImage := imaging.Rotate90(imageToRotate)
	return dstImage
}

func rotateImageRight(imageToRotate image.Image) image.Image {
	dstImage := imaging.Rotate270(imageToRotate)
	return dstImage
}

func flipImage(imageToRotate image.Image) image.Image {
	dstImage := imaging.Rotate180(imageToRotate)
	return dstImage
}
