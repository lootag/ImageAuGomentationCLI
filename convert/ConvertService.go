package convert

import (
	"os"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/annotationWriters"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type ConvertService struct {
}

func (convertSerivce ConvertService) ConvertToJPG(toConvert chan entities.ImageInfo,
	mainWaitGroup *sync.WaitGroup) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	err := os.Mkdir("./AugmentedImages", 0755)
	if err != nil {
	}
	for image := range toConvert {
		wg.Add(1)
		go convertToJPGWorker(image, &wg)
	}
	wg.Wait()
}

func (convertSerivce ConvertService) ConvertToText(annotationToConvert chan entities.Annotation,
	mainWaitGroup *sync.WaitGroup,
	annotationType entities.AnnotationType) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	var factory annotationWriters.AnnotationsWritersFactory
	err := os.Mkdir("./AugmentedAnnotations", 0755)
	writer, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotation := range annotationToConvert {
		wg.Add(1)
		go writer.Write(annotation, &wg)
	}
	wg.Wait()
}
