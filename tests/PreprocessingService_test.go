package tests;

import(
	"testing";
	_"image/jpeg";
	_"sync";
	_"os";
	_"github.com/lootag/ImageAuGomentationCLI/preprocess";
	_"image";
	_"strconv";
)

func TestPreprocessingReturnsAResizedJPG(t *testing.T){
	/*
	var preprocessingService preprocess.PreprocessingService;
	filePaths := []string {"../Images/instagram_copy.png", "../Images/38295588_481089599031169_6763125026364325888_n.jpg", "../Images/43408237_2228882777185642_761152503110039906_n.jpg"};
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
	*/
	/*
	err := os.Remove("outImage1.jpg");
	if err != nil{
		t.Errorf("couldn't remove outImage1");
	}
	err = os.Remove("outImage2.jpg");
	if err != nil{
		t.Errorf("couldn't remove outImage2");
	}
	*/
}