package tests

import (
	"github.com/lootag/ImageAuGomentationCLI/blur"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"sync"
	"testing"
)

func TestBlurServiceReturnsABlurredImage(t *testing.T) {

	var blurService blur.BlurService
	filePath := "testImages/instagram_copy.png"
	sourceBytes, err := os.Open(filePath)
	if err != nil {
		t.Errorf("There was a problem opening the file " + filePath)
	}
	defer sourceBytes.Close()
	sourceImage, _, err := image.Decode(sourceBytes)
	var sourceImageInfo entities.ImageInfo
	sourceImageInfo.ImageSource = sourceImage
	sourceImageInfo.OriginalFileName = filePath
	if err != nil {
		t.Errorf("There was a problem decodeing the image")
	}
	toAugment := make(chan entities.ImageInfo, 1000)
	augmented := make(chan entities.ImageInfo, 1000)
	toAugment <- sourceImageInfo
	close(toAugment)
	var toAugmentSlice []entities.ImageInfo
	for image := range toAugment {
		toAugmentSlice = append(toAugmentSlice, image)
	}
	var wg sync.WaitGroup
	var options entities.Options
	options.Sigma = 20
	wg.Add(1)
	go blurService.Augment(&toAugmentSlice, &wg, augmented, options)
	wg.Wait()
	index := 0
	for image := range augmented {
		index += 1
		outputFile, err := os.Create("outImage" + strconv.Itoa(index) + ".jpg")
		defer outputFile.Close()
		if err != nil {
			t.Errorf("There was a problem opening the file outimage" + strconv.Itoa(index) + ".jpg")
		}
		err = jpeg.Encode(outputFile, image.ImageSource, nil)
		if err != nil {
			t.Errorf("There was an error saving the images to fs")
		}
	}

	err = os.Remove("outImage1.jpg")
	if err != nil {
		t.Errorf("couldn't remove outImage1")
	}
}
