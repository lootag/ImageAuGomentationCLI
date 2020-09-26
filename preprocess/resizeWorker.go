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
	newImage := resize.Resize(uint(size), 0, *decodedImage,resize.Lanczos3);
	resized <- newImage;
	fmt.Println("resized image sent");

}