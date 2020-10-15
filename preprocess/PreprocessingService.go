package preprocess

import (
	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type PreprocessingService struct {
}

//Implements Preprocessor
func (preprocessingService PreprocessingService) Preprocess(images *[]string,
	fileNames *[]string,
	resized chan entities.ImageInfo,
	resizedCopy chan entities.ImageInfo,
	resizedAnnotations chan entities.Annotation,
	resizedAnnotationsCopy chan entities.Annotation,
	annotationType entities.AnnotationType,
	size int,
	xmls bool,
	mainWaitGroup *sync.WaitGroup) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	checked := make(chan entities.ImageInfo)
	checkedAnnotations := make(chan string)
	annotationsToResize := make(chan entities.Annotation)

	wg.Add(2)
	go checkAllFilesAreImages(images, fileNames, checked, &wg)
	go resizing(checked, resized, resizedCopy, size, &wg)
	if xmls {
		wg.Add(3)
		go checkAllImagesAreAnnotated(fileNames,
			&wg,
			checkedAnnotations)
		go readAnnotations(annotationType,
			checkedAnnotations,
			&wg,
			annotationsToResize)
		go resizeAnnotations(annotationsToResize,
			resizedAnnotations,
			resizedAnnotationsCopy,
			size,
			&wg)
	}
	wg.Wait()
}

func checkAllFilesAreImages(imagePaths *[]string,
	fileNames *[]string,
	checked chan entities.ImageInfo,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := 0; imageIndex < len(*imagePaths); imageIndex++ {
		wg.Add(1)
		go checkAllFilesAreImagesWorker((*imagePaths)[imageIndex],
			(*fileNames)[imageIndex],
			&wg,
			checked)
	}
	wg.Wait()
	close(checked)

}

func checkAllImagesAreAnnotated(fileNames *[]string,
	preprocessWaitGroup *sync.WaitGroup,
	checkedAnnotations chan string) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := 0; imageIndex < len(*fileNames); imageIndex++ {
		wg.Add(1)
		go checkAllImagesAreAnnotatedWorker((*fileNames)[imageIndex],
			&wg,
			checkedAnnotations)
	}
	wg.Wait()
	close(checkedAnnotations)

}

func readAnnotations(annotationType entities.AnnotationType,
	checkedAnnotations chan string,
	preprocessWaitGroup *sync.WaitGroup,
	annotationsToResize chan entities.Annotation) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotation := range checkedAnnotations {
		wg.Add(1)
		go annotationReader.Read(annotation, annotationsToResize, &wg)
	}
	wg.Wait()
	close(annotationsToResize)
}

func resizeAnnotations(annotationsToResize chan entities.Annotation,
	resizedAnnotations chan entities.Annotation,
	resizeAnnotationsCopy chan entities.Annotation,
	newSize int,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for annotation := range annotationsToResize {
		wg.Add(1)
		go resizeAnnotationWorker(annotation, resizedAnnotations, resizeAnnotationsCopy, newSize, &wg)
	}
	wg.Wait()
	close(resizedAnnotations)
	close(resizeAnnotationsCopy)
}

func resizing(checked chan entities.ImageInfo,
	resized chan entities.ImageInfo,
	resizedCopy chan entities.ImageInfo,
	size int,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for image := range checked {
		wg.Add(1)
		go resizeWorker(image, resized, resizedCopy, &wg, size)
	}
	wg.Wait()
	close(resized)
	close(resizedCopy)

}
