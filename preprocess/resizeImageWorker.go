package preprocess

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/nfnt/resize"
)

func resizeImageWorker(decodedImage entities.ImageInfo,
	resizedImage chan entities.ImageInfo,
	resizedImageCopy chan entities.ImageInfo,
	wg *sync.WaitGroup,
	size int) {
	defer (*wg).Done()
	var imageInfo entities.ImageInfo
	imageInfo.OriginalFileName = decodedImage.OriginalFileName
	imageInfo.NewName = "resized" + decodedImage.OriginalFileName
	imageInfo.ImageSource = resize.Resize(uint(size), uint(size), decodedImage.ImageSource, resize.Lanczos3)
	resizedImage <- imageInfo
	resizedImageCopy <- imageInfo
}
