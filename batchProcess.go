package main

import (
	"fmt"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"math"
	"strconv"
	"sync"
)

func batchProcess(options *entities.Options,
	imagePaths *[]string,
	imageNames *[]string,
	preprocessor *preprocess.Preprocessor,
	converter *convert.Converter,
	garbageCollector *collectGarbage.GarbageCollector,
	classesToExclude []string) {
	pathsToProcess := []string{}
	namesToProcess := []string{}

	if (*options).BatchSize > len(*imagePaths) {
		panic("The batch size you've set is larger than the number of elements you intend to process. Exiting.")
	}

	numberOfBatches := int(math.Floor(float64(len(*imagePaths))) / float64((*options).BatchSize))
	if len(*imagePaths)%(*options).BatchSize != 0 {
		numberOfBatches += 1
	}
	for index := 0; index < numberOfBatches; index++ {
		if index == numberOfBatches-1 {
			fmt.Println("Processing batch " + strconv.Itoa(index+1) + " out of " + strconv.Itoa(numberOfBatches))
			start := index * (*options).BatchSize
			pathsToProcess = (*imagePaths)[start:]
			namesToProcess = (*imageNames)[start:]
		} else {
			fmt.Println("Processing batch " + strconv.Itoa(index+1) + " out of " + strconv.Itoa(numberOfBatches))
			start := index * (*options).BatchSize
			end := start + (*options).BatchSize
			pathsToProcess = (*imagePaths)[start:end]
			namesToProcess = (*imageNames)[start:end]
		}

		resized := make(chan entities.ImageInfo, (*options).BatchSize)
		resizedCopy := make(chan entities.ImageInfo, (*options).BatchSize)
		resizedAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		resizedAnnotationsCopy := make(chan entities.Annotation, (*options).BatchSize)
		resizedRotate := []entities.ImageInfo{}
		resizedRotateAnnotations := []entities.Annotation{}
		var wg sync.WaitGroup
		wg.Add(2)
		go (*preprocessor).Preprocess(&pathsToProcess,
			&namesToProcess,
			resized,
			resizedCopy,
			resizedAnnotations,
			resizedAnnotationsCopy,
			(*options).InAnnotationType,
			(*options).Size,
			(*options).Xml,
			classesToExclude,
			&wg)
		go (*converter).ConvertToJPG(resized, &wg, "resize", &pathsToProcess)
		if (*options).Xml {
			wg.Add(1)
			go (*converter).ConvertToText(resizedAnnotations, &wg, (*options).OutAnnotationType)
		}
		wg.Wait()
		for image := range resizedCopy {
			resizedRotate = append(resizedRotate, image)
		}
		if (*options).Xml{
			for annotation := range resizedAnnotationsCopy {
				resizedRotateAnnotations = append(resizedRotateAnnotations, annotation)
			}
		}
		
		resizedBlur := resizedRotate
		resizedBlurAnnotations := resizedRotateAnnotations
		actions := []string{"rotate", "blur"}
		rotated := make(chan entities.ImageInfo, (*options).BatchSize)
		blurred := make(chan entities.ImageInfo, (*options).BatchSize)
		rotatedAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		blurredAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		augmentation, err := AugmentationsFactory(actions[0])
		if err != nil {
			panic(err)
		}
		if (*options).Side != entities.NIL_DIRECTION {
			wg.Add(2)
			go augmentation.Augment(&resizedRotate,
				&resizedRotateAnnotations,
				&wg,
				rotated,
				rotatedAnnotations,
				*options)
			go (*converter).ConvertToJPG(rotated, &wg, actions[0], &pathsToProcess)
			if (*options).Xml {
				wg.Add(1)
				go (*converter).ConvertToText(rotatedAnnotations, &wg, (*options).OutAnnotationType)
			}
		}

		augmentation, err = AugmentationsFactory(actions[1])
		if err != nil {
			panic(err)
		}

		if (*options).Sigma != 0 {
			wg.Add(2)
			go augmentation.Augment(&resizedBlur,
				&resizedBlurAnnotations,
				&wg,
				blurred,
				blurredAnnotations,
				*options)
			go (*converter).ConvertToJPG(blurred, &wg, actions[1], &pathsToProcess)
			if (*options).Xml {
				wg.Add(1)
				go (*converter).ConvertToText(blurredAnnotations, &wg, (*options).OutAnnotationType)
			}
		}

		wg.Wait()
		if (*options).Xml {
			(*garbageCollector).CollectGarbage()
		}

	}

}
