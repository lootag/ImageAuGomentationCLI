package preprocess

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
)

func checkAllFilesAreImagesWorker(file string,
	wg *sync.WaitGroup,
	checked chan image.Image) {
	defer (*wg).Done()
	lastCharacter := len(file)
	thirdToLastCharacter := len(file) - 3
	imageFormat := file[thirdToLastCharacter:lastCharacter]
	imageFile, err := os.Open(file)
	if err != nil {
		panic("There was a problem opening the file " + file)
	}

	decodedImage, _, err := image.Decode(imageFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("The format is " + imageFormat)

	if imageFormat == "jpg" || imageFormat == "png" {
		checked <- decodedImage
		fmt.Println("checked image sent")
	} else {
		panic("The file " + file + " is not an image, or its format is not supported")
	}

}
