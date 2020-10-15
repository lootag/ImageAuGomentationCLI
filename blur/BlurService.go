package blur

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type BlurService struct {
}

func (blurService BlurService) Augment(imagesToAugment *[]entities.ImageInfo,
	annotationsToAugment *[]entities.Annotation,
	mainWaitGroup *sync.WaitGroup,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	options entities.Options) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup;
	wg.Add(1);
	go blur(imagesToAugment,
		annotationsToAugment, 
		&wg, 
		augmentedImages,
		augmentedAnnotations, 
		options.Sigma, 
		options.Xml);
	wg.Wait();
}

func blur(imagesToBlur *[]entities.ImageInfo,
	annotationsToBlur *[]entities.Annotation,
	blurWaitGroup *sync.WaitGroup,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	sigma float64,
	annotated bool) {
	defer (*blurWaitGroup).Done();
	var wg sync.WaitGroup;
	for imageIndex := range *imagesToBlur{
		wg.Add(1);
		go blurImageWorker((*imagesToBlur)[imageIndex], &wg, augmentedImages, sigma);
	}
	if annotated{
		for annotationIndex := range *annotationsToBlur{
			wg.Add(1);
			go blurAnnotationWorker((*annotationsToBlur)[annotationIndex], augmentedAnnotations, &wg);
		}	
	}

	wg.Wait();
	close(augmentedImages);
	close(augmentedAnnotations);
}
