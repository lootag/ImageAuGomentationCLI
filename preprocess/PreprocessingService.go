package preprocess

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type PreprocessingService struct {
}

//Implements Preprocessor
func (preprocessingService PreprocessingService) Preprocess(imagePaths []string,
	imageNames []string,
	resizedImages chan entities.ImageInfo,
	resizedImagesCopy chan entities.ImageInfo,
	resizedAnnotations chan entities.Annotation,
	resizedAnnotationsCopy chan entities.Annotation,
	annotationType entities.AnnotationType,
	size int,
	annotated bool,
	classesToExclude []string,
	mainWaitGroup *sync.WaitGroup) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	validatedImages := make(chan entities.ImageInfo)
	validatedAnnotations := make(chan string)
	annotationsToResize := make(chan entities.Annotation)

	wg.Add(2)
	go checkAllFilesAreImages(imagePaths, imageNames, validatedImages, &wg)
	go resizeImages(validatedImages, resizedImages, resizedImagesCopy, size, &wg)
	if annotated {
		wg.Add(3)
		go checkAllImagesAreAnnotated(imageNames,
			&wg,
			validatedAnnotations)
		go readAnnotations(annotationType,
			validatedAnnotations,
			&wg,
			annotationsToResize)
		go resizeAnnotations(annotationsToResize,
			resizedAnnotations,
			resizedAnnotationsCopy,
			size,
			classesToExclude,
			&wg)
	}
	wg.Wait()
}

func checkAllFilesAreImages(imagePaths []string,
	fileNames []string,
	validatedImages chan entities.ImageInfo,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := 0; imageIndex < len(imagePaths); imageIndex++ {
		wg.Add(1)
		go checkAllFilesAreImagesWorker(imagePaths[imageIndex],
			fileNames[imageIndex],
			&wg,
			validatedImages)
	}
	wg.Wait()
	close(validatedImages)

}

func checkAllImagesAreAnnotated(fileNames []string,
	preprocessWaitGroup *sync.WaitGroup,
	validatedAnnotations chan string) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := 0; imageIndex < len(fileNames); imageIndex++ {
		wg.Add(1)
		go checkAllImagesAreAnnotatedWorker(fileNames[imageIndex],
			&wg,
			validatedAnnotations)
	}
	wg.Wait()
	close(validatedAnnotations)

}

func readAnnotations(annotationType entities.AnnotationType,
	validatedAnnotations chan string,
	preprocessWaitGroup *sync.WaitGroup,
	annotationsToResize chan entities.Annotation) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotation := range validatedAnnotations {
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
	classesToExclude []string,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for annotation := range annotationsToResize {
		intersection := intersectStringArrays(annotation.Classes, classesToExclude)
		if len(intersection) == 0 {
			wg.Add(1)
			go resizeAnnotationWorker(annotation, resizedAnnotations, resizeAnnotationsCopy, newSize, &wg)
		}
	}
	wg.Wait()
	close(resizedAnnotations)
	close(resizeAnnotationsCopy)
}

func resizeImages(validatedImages chan entities.ImageInfo,
	resizedImages chan entities.ImageInfo,
	resizedImagesCopy chan entities.ImageInfo,
	size int,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for image := range validatedImages {
		wg.Add(1)
		go resizeImageWorker(image, resizedImages, resizedImages, &wg, size)
	}
	wg.Wait()
	close(resizedImages)
	close(resizedImagesCopy)

}

func intersectStringArrays(stringArray1 []string, stringArray2 []string) []string {
	intersection := []string{}
	for stringArrayIndex1 := range stringArray1 {
		if stringArraycontains(stringArray2, stringArray1[stringArrayIndex1]) {
			intersection = append(intersection, stringArray1[stringArrayIndex1])
		}
	}
	return intersection
}

func stringArraycontains(stringArray1 []string, toCheck string) bool {
	for stringIndex := range stringArray1 {
		if stringArray1[stringIndex] == toCheck {
			return true
		}
	}
	return false
}
