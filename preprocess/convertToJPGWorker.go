package preprocess;

import(
	"fmt";
	"sync";
	"image";
	"image/jpeg";
	"bytes";
	"bufio";
)


func convertToJPGWorker(decodedImage *image.Image, 
	toAugment chan []byte,
	wg *sync.WaitGroup){
	defer (*wg).Done();

	var jpegBuffer bytes.Buffer;
	writer := bufio.NewWriter(&jpegBuffer);

	err := jpeg.Encode(writer, *decodedImage, nil);
	fmt.Println("Just encoded the image");

	if err != nil{
		panic("There was a problem encoding an image to jpg");
	}

	toSend := jpegBuffer.Bytes();
	toAugment <- toSend;
	fmt.Println("the bytes were passed");
}