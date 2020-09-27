package main;

import(
	"image";
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type Augmenter interface{
	Augment(toAugment chan image.Image, 
			wg *sync.WaitGroup, 
			augmented chan image.Image, 
			options entities.Options);
}

