package convert

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	_ "image"
	"image/jpeg"
	"os"
	"sync"
)

func convertToJPGWorker(imageToConvert entities.ImageInfo,
	wg *sync.WaitGroup) {
	defer (*wg).Done()
	outputFile, err := os.Create("./AugmentedImages/" + imageToConvert.NewName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, imageToConvert.ImageSource, nil)
	if err != nil {
		panic(err)
	}

}
