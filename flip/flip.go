package flip

import (
	"fmt"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"sync"
)

type FlipService struct {
}

func (flipService FlipService) Augment(toAugment chan image.Image,
	wg *sync.WaitGroup,
	augmented chan image.Image,
	options entities.Options) {
	defer wg.Done()
	flip()
}

func flip() {
	fmt.Println("Flip")
}
