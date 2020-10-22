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
