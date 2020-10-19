package preprocess

import (
	"fmt"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
)

func checkAllFilesAreImagesWorker(file string,
	fileName string,
	wg *sync.WaitGroup,
	checked chan entities.ImageInfo) {
	defer (*wg).Done()
	lastCharacter := len(file)
	thirdToLastCharacter := len(file) - 3
	imageFormat := file[thirdToLastCharacter:lastCharacter]
	imageFile, err := os.Open(file)
	if err != nil {
		panic("There was a problem opening the file " + file)
	}
	defer imageFile.Close()
	decodedImage, _, err := image.Decode(imageFile)

	if err != nil {
		panic("Couldn't decode " + file)
	}

	var imageInfo entities.ImageInfo
	imageInfo.OriginalFileName = fileName
	imageInfo.ImageSource = decodedImage

	if imageFormat == "jpg" || imageFormat == "png" {
		checked <- imageInfo
	} else {
		fmt.Println("The file " + file + " is not an image, or its format is not supported. Ignoring.")
	}

}
