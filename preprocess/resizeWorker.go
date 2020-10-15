package preprocess

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/nfnt/resize"
	"sync"
)

func resizeWorker(decodedImage entities.ImageInfo,
	resized chan entities.ImageInfo,
	resizedCopy chan entities.ImageInfo,
	wg *sync.WaitGroup,
	size int) {
	defer (*wg).Done()
	var imageInfo entities.ImageInfo
	imageInfo.OriginalFileName = decodedImage.OriginalFileName
	imageInfo.NewName = "resized" + decodedImage.OriginalFileName
	imageInfo.ImageSource = resize.Resize(uint(size), uint(size), decodedImage.ImageSource, resize.Lanczos3)
	resized <- imageInfo
	resizedCopy <- imageInfo
}
