package main

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type Augmenter interface {
	Augment(imagesToAugment *[]entities.ImageInfo,
		annotationsToAugment *[]entities.Annotation,
		wg *sync.WaitGroup,
		augmentedImages chan entities.ImageInfo,
		augmentedAnnotation chan entities.Annotation,
		options entities.Options)
}
