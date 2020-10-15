package rotate

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type RotateService struct {
}

func (rotateService RotateService) Augment(imagesToAugment *[]entities.ImageInfo,
	annotationsToAugment *[]entities.Annotation,
	mainWaitGroup *sync.WaitGroup,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	options entities.Options) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	wg.Add(1)
	go rotate(imagesToAugment, 
		annotationsToAugment,
		&wg, 
		options.Side, 
		augmentedImages,
		augmentedAnnotations,
		options.Xml)
	wg.Wait()
}

func rotate(imagesToRotate *[]entities.ImageInfo,
	annotationsToRotate *[]entities.Annotation,
	augmentWaitGroup *sync.WaitGroup,
	direction entities.Direction,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	annotated bool) {
	defer (*augmentWaitGroup).Done()
	var wg sync.WaitGroup;
	for imageIndex := range *imagesToRotate {
		wg.Add(1)
		go rotateImageWorker((*imagesToRotate)[imageIndex], &wg, direction, augmentedImages)
	}

	if annotated{
		for annotationIndex := range *annotationsToRotate{
			wg.Add(1)
			go rotateAnnotationWorker((*annotationsToRotate)[annotationIndex], 
			augmentedAnnotations,
			direction, 
			&wg);
		}
	}
	wg.Wait()
	close(augmentedImages);
	close(augmentedAnnotations);
}
