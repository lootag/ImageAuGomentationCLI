package main

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type Augmenter interface {
	Augment(imagesToAugment []entities.ImageInfo,
		annotationsToAugment []entities.Annotation,
		wg *sync.WaitGroup,
		augmentedImages chan entities.ImageInfo,
		augmentedAnnotation chan entities.Annotation,
		options entities.Options)
}
