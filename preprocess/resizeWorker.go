package preprocess;

import(
	"fmt";
	"image";
	"sync";
	"github.com/nfnt/resize";
)

func resizeWorker(decodedImage *image.Image,
				  resized chan image.Image,
				  wg *sync.WaitGroup,
				  size int){
	defer (*wg).Done();
	fmt.Println("entered resizeWorker")
	newImage := resize.Resize(uint(size), uint(size), *decodedImage,resize.Lanczos3);
	fmt.Println("New image created")
	resized <- newImage;
	fmt.Println("resized image sent");

}