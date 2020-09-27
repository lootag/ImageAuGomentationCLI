package rotate

import (
	"fmt"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"sync"
)

type RotateService struct {
}

func (rotateService RotateService) Augment(toAugment chan image.Image,
	mainWaitGroup *sync.WaitGroup,
	augmented chan image.Image,
	options entities.Options) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	wg.Add(1)
	go rotate(toAugment, &wg, options.Side, augmented)
	wg.Wait()
}

func rotate(toRotate chan image.Image,
	augmentWaitGroup *sync.WaitGroup,
	direction entities.Direction,
	rotated chan image.Image) {
	defer (*augmentWaitGroup).Done()
	var wg sync.WaitGroup
	for image := range toRotate {
		wg.Add(1)
		go rotateWorker(image, &wg, direction, rotated)
	}
	wg.Wait()
	fmt.Println("rotate routine complete.")
}
