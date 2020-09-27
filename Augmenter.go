package main

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"image"
	"sync"
)

type Augmenter interface {
	Augment(toAugment chan image.Image,
		wg *sync.WaitGroup,
		augmented chan image.Image,
		options entities.Options)
}
