/*
This file is part of ImageAuGomentationCLI.

ImageAuGomentationCLI is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

ImageAuGomentationCLI is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ImageAuGomentationCLI.  If not, see <https://www.gnu.org/licenses/>.
*/
package orchestration

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
)

func BatchProcess(options *entities.Options,
	imagePaths []string,
	imageNames []string,
	preprocessor *preprocess.Preprocessor,
	converter *convert.Converter,
	garbageCollector *collectGarbage.GarbageCollector,
	classesToExclude []string) {
	pathsToProcess := []string{}
	namesToProcess := []string{}

	if (*options).BatchSize > len(imagePaths) {
		panic("The batch size you've set is larger than the number of elements you intend to process. Exiting.")
	}

	//TODO:Make method to get number of batches
	numberOfBatches := int(math.Floor(float64(len(imagePaths))) / float64((*options).BatchSize))
	if len(imagePaths)%(*options).BatchSize != 0 {
		numberOfBatches += 1
	}
	for index := 0; index < numberOfBatches; index++ {
		//TODO:Make method to get images and paths to process
		if index == numberOfBatches-1 {
			fmt.Println("Processing batch " + strconv.Itoa(index+1) + " out of " + strconv.Itoa(numberOfBatches))
			start := index * (*options).BatchSize
			pathsToProcess = imagePaths[start:]
			namesToProcess = imageNames[start:]
		} else {
			fmt.Println("Processing batch " + strconv.Itoa(index+1) + " out of " + strconv.Itoa(numberOfBatches))
			start := index * (*options).BatchSize
			end := start + (*options).BatchSize
			pathsToProcess = imagePaths[start:end]
			namesToProcess = imageNames[start:end]
		}

		resized := make(chan entities.ImageInfo, (*options).BatchSize)
		resizedCopy := make(chan entities.ImageInfo, (*options).BatchSize)
		resizedAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		resizedAnnotationsCopy := make(chan entities.Annotation, (*options).BatchSize)
		resizedRotate := []entities.ImageInfo{}
		resizedRotateAnnotations := []entities.Annotation{}
		var wg sync.WaitGroup
		wg.Add(2)
		go (*preprocessor).Preprocess(pathsToProcess,
			namesToProcess,
			resized,
			resizedCopy,
			resizedAnnotations,
			resizedAnnotationsCopy,
			(*options).InAnnotationType,
			(*options).Folder,
			(*options).Size,
			(*options).Annotated,
			classesToExclude,
			&wg)
		go (*converter).ConvertToJPG(resized, &wg)
		if (*options).Annotated {
			wg.Add(1)
			go (*converter).ConvertToText(resizedAnnotations, &wg, (*options).OutAnnotationType)
		}
		wg.Wait()
		for image := range resizedCopy {
			resizedRotate = append(resizedRotate, image)
		}
		if (*options).Annotated {
			for annotation := range resizedAnnotationsCopy {
				resizedRotateAnnotations = append(resizedRotateAnnotations, annotation)
			}
		}

		resizedBlur := resizedRotate
		resizedBlurAnnotations := resizedRotateAnnotations
		actions := []string{"rotate", "blur"}
		rotated := make(chan entities.ImageInfo)
		blurred := make(chan entities.ImageInfo)
		rotatedAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		blurredAnnotations := make(chan entities.Annotation, (*options).BatchSize)
		//TODO: Try to make a dispatcher of some sort here...
		augmentation, err := AugmentationsFactory(actions[0])
		if err != nil {
			panic(err)
		}
		if (*options).Direction != entities.NIL_DIRECTION {
			wg.Add(2)
			go augmentation.Augment(resizedRotate,
				resizedRotateAnnotations,
				&wg,
				rotated,
				rotatedAnnotations,
				*options)
			go (*converter).ConvertToJPG(rotated, &wg)
			if (*options).Annotated {
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
			go augmentation.Augment(resizedBlur,
				resizedBlurAnnotations,
				&wg,
				blurred,
				blurredAnnotations,
				*options)
			go (*converter).ConvertToJPG(blurred, &wg)
			if (*options).Annotated {
				wg.Add(1)
				go (*converter).ConvertToText(blurredAnnotations, &wg, (*options).OutAnnotationType)
			}
		}

		wg.Wait()
		if (*options).Annotated {
			(*garbageCollector).CollectGarbage()
		}

	}

}
