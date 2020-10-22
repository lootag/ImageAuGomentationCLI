package blur

import (
	"sync"

	"github.com/disintegration/imaging"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func blurImageWorker(imageToBlur entities.ImageInfo,
	blurWaitGroup *sync.WaitGroup,
	augmented chan entities.ImageInfo,
	sigma float64) {
	defer (*blurWaitGroup).Done()
	var blurredImage entities.ImageInfo
	blurredImage.ImageSource = imaging.Blur(imageToBlur.ImageSource, sigma)
	blurredImage.OriginalFileName = imageToBlur.OriginalFileName
	blurredImage.NewName = "blur" + imageToBlur.OriginalFileName
	augmented <- blurredImage
}
