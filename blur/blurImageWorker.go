package blur;

import(
	"sync";
	"github.com/disintegration/imaging";
	"github.com/lootag/ImageAuGomentationCLI/entities";

)


func blurImageWorker(imageToBlur entities.ImageInfo,
	blurWaitGroup *sync.WaitGroup,
	augmented chan entities.ImageInfo,
	sigma float64){
	defer (*blurWaitGroup).Done();
	var dstImage entities.ImageInfo;
	dstImage.ImageSource = imaging.Blur(imageToBlur.ImageSource, sigma);
	dstImage.OriginalFileName = imageToBlur.OriginalFileName;
	dstImage.NewName = "blur" + imageToBlur.OriginalFileName;
	augmented <- dstImage;
}

