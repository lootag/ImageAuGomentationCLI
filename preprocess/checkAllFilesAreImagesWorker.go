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
package preprocess

import (
	"fmt"
	"image"
	"os"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func checkAllFilesAreImagesWorker(imagePath string,
	imageName string,
	wg *sync.WaitGroup,
	checked chan entities.ImageInfo) {
	defer (*wg).Done()
	lastCharacter := len(imagePath)
	thirdToLastCharacter := len(imagePath) - 3
	imageFormat := imagePath[thirdToLastCharacter:lastCharacter]
	imageFile, err := os.Open(imagePath)
	if err != nil {
		panic("There was a problem opening the file " + imagePath)
	}
	defer imageFile.Close()
	decodedImage, _, err := image.Decode(imageFile)

	if err != nil {
		panic("Couldn't decode " + imagePath)
	}

	var imageInfo entities.ImageInfo
	imageInfo.OriginalFileName = imageName
	imageInfo.ImageSource = decodedImage

	if imageFormat == "jpg" || imageFormat == "png" {
		checked <- imageInfo
	} else {
		fmt.Println("The file " + imagePath + " is not an image, or its format is not supported. Ignoring.")
	}

}
