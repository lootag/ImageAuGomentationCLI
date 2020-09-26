package tests;

import(
	"fmt";
	"testing";
	"bytes";
	"image/jpeg";
	"sync";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/preprocess";
)

func TestPreprocessingReturnsAResizedJPG(t *testing.T){
	
	var preprocessingService preprocess.PreprocessingService;
	filePaths := []string {"testImages/instagram.png"};
	outputFile, err := os.Create("outImage.jpg");
	if err != nil{
		t.Errorf("There was an error creating the output file");
	}
	toAugment := make(chan []byte, 1000000);
	size := 480;
	var wg sync.WaitGroup;
	wg.Add(1);
	go preprocessingService.Preprocess(&filePaths, toAugment, size, &wg);
	wg.Wait();
	for image := range toAugment{
		fmt.Println(image);
		decodedImage, err := jpeg.Decode(bytes.NewReader(image));
		if err != nil{
			t.Errorf("There was an error decoding the images");
		}
		err = jpeg.Encode(outputFile, decodedImage, nil);

		if err != nil{
			t.Errorf("There was an error saving the images to fs");
		}
	}
}