package rotate

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"sync"
)

func rotateWorker(imageToRotate image.Image,
	wg *sync.WaitGroup,
	direction entities.Direction,
	rotated chan image.Image) {

	defer (*wg).Done()
	fmt.Println("entered rotateWorker.")
	if direction == entities.LEFT {
		fmt.Println("rotating left...")
		rotated <- rotateLeft(imageToRotate, direction)
		fmt.Println("image sent.")
	} else if direction == entities.RIGHT {
		rotated <- rotateRight(imageToRotate, direction)
	} else if direction == entities.ALL {
		rotated <- rotateLeft(imageToRotate, direction)
		rotated <- rotateRight(imageToRotate, direction)
	} else {
		panic("The argument provided in the rotate field is not valid. Exiting.")
	}

	fmt.Println("done rotating")
}

func rotateLeft(imageToRotate image.Image,
	direction entities.Direction) image.Image {
	dstImage := imaging.Rotate90(imageToRotate)
	return dstImage
}

func rotateRight(imageToRotate image.Image,
	direction entities.Direction) image.Image {
	dstImage := imaging.Rotate270(imageToRotate)
	return dstImage
}
