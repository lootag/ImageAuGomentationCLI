package preprocess;

import(
	"fmt";
	"sync";
	"image";
)

type PreprocessingService struct{

}

//Implements IPreprocessingService
func (preprocessingService PreprocessingService) Preprocess(images *[]string, 
															toAugment chan []byte, 
															size int, 
															mainWaitGroup *sync.WaitGroup){
	defer (*mainWaitGroup).Done();
	var wg sync.WaitGroup;
	checked := make(chan image.Image);
	resized := make(chan image.Image)
	
	wg.Add(3);
	go checkAllFilesAreImages(images, checked, &wg);
	go resizing(checked, resized, size, &wg);
	go convertToJPG(resized, toAugment, &wg);
	wg.Wait();
}

func checkAllFilesAreImages(images *[]string, 
							checked chan image.Image,
							preprocessWaitGroup *sync.WaitGroup){
	defer (*preprocessWaitGroup).Done();
	var wg sync.WaitGroup;
	for imageIndex := 0; imageIndex < len(*images); imageIndex++{
		wg.Add(1);
		go checkAllFilesAreImagesWorker((*images)[imageIndex],
										&wg,
										checked)
	}
	wg.Wait();
	close(checked);
	fmt.Println("checked closed");

}

func convertToJPG(resized chan image.Image, 
				  toAugment chan []byte,
				  preprocessWaitGroup *sync.WaitGroup){
	
	defer (*preprocessWaitGroup).Done();
	var wg sync.WaitGroup;
	for image := range resized{
		fmt.Println("here's an image to encode")
		wg.Add(1);
		go convertToJPGWorker(&image, toAugment, &wg);
	}
	wg.Wait();
	close(toAugment);
	fmt.Println("image converted");
}

func resizing(checked chan image.Image, 
			  resized chan image.Image, 
			  size int,
			  preprocessWaitGroup *sync.WaitGroup){
	defer (*preprocessWaitGroup).Done();
	var wg sync.WaitGroup;
	for image := range checked{
		wg.Add(1);
		go resizeWorker(&image, resized, &wg, size);
	}
	wg.Wait();
	close(resized);
	fmt.Println("resized closed")
	
}
