package tests;

import(
	"testing";
	"image/jpeg";
	"sync";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/preprocess";
	"image";
	"strconv";
)

func TestPreprocessingReturnsAResizedJPG(t *testing.T){
	
	var preprocessingService preprocess.PreprocessingService;
	filePaths := []string {"testImages/instagram.png", "testImages/instagram_copy.png"};
	resized := make(chan image.Image, 1000);
	size := 480;
	var wg sync.WaitGroup;
	wg.Add(1);
	go preprocessingService.Preprocess(&filePaths, resized, size, &wg);
	wg.Wait();
	index := 0;
	for image := range resized{
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

	err := os.Remove("outImage1.jpg");
	if err != nil{
		t.Errorf("couldn't remove outImage1");
	}
	err = os.Remove("outImage2.jpg");
	if err != nil{
		t.Errorf("couldn't remove outImage2");
	}
}