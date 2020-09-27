package tests;

import(
	"testing";
	"image/jpeg";
	"sync";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/rotate";
	"github.com/lootag/ImageAuGomentationCLI/entities";
	"image";
	"strconv";
)

func TestRotateServiceReturnsARotatedImage(t *testing.T){
	
	var rotateService rotate.RotateService;
	filePath := "testImages/instagram.png";
	sourceBytes, err := os.Open(filePath);
	if err != nil{
		t.Errorf("There was a problem opening the file " + filePath);
	}
	defer sourceBytes.Close();
	sourceImage, _, err := image.Decode(sourceBytes);
	if err != nil{
		t.Errorf("There was a problem decodeing the image");
	}
	toAugment := make(chan image.Image, 1000);
	augmented := make(chan image.Image, 1000);
	toAugment <- sourceImage;
	close(toAugment);
	var wg sync.WaitGroup;
	var options entities.Options;
	options.Side = entities.ALL;
	wg.Add(1);
	go rotateService.Augment(toAugment, &wg, augmented, options);
	wg.Wait();
	close(augmented);
	index := 0;
	for image := range augmented{
		index += 1;
		outputFile, err := os.Create("outImage" + strconv.Itoa(index) + ".jpg");
		defer outputFile.Close();
		if err != nil{
			t.Errorf("There was a problem opening the file outimage" + strconv.Itoa(index) + ".jpg")
		}
		err = jpeg.Encode(outputFile, image, nil);
		if err != nil{
			t.Errorf("There was an error saving the images to fs");
		}
	}

	err = os.Remove("outImage1.jpg");
	if err != nil{
		t.Errorf("couldn't remove outImage1");
	}

	err = os.Remove("outImage2.jpg");

	if err != nil{
		t.Errorf("couldn't remove outImage2");
	}
}