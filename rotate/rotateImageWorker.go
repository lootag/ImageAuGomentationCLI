/*
This file is part of ImageAuGomentationCLI.

ImageAuGomentationCLI is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

ImageAuGomentationCLI is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ImageAuGomentationCLI.  If not, see <https://www.gnu.org/licenses/>.
*/
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
