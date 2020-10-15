package convert;

import(
	"sync";
	_"image";
	"image/jpeg";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

func convertToJPGWorker(imageToConvert entities.ImageInfo,
	wg *sync.WaitGroup){
		defer (*wg).Done();
		outputFile, err := os.Create("./AugmentedImages/" + imageToConvert.NewName);
		if err != nil{
			panic(err);
		}
		defer outputFile.Close();
		err = jpeg.Encode(outputFile, imageToConvert.ImageSource, nil);
		if err != nil{
			panic(err);
		}

}
