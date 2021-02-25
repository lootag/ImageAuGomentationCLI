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

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func (this Application) runAugmentationJob(job augmentationJob) {
	//The application processes the images (and the annotations, if present)
	//asynchronously. However, if a large number of images is fed to the
	//application, it will not be able to handle them all asynchronously, so
	//the data is split up into batches which are processed sequentially.

	numberOfBatches := this.computeNumberOfBatches(job.Options, job.ImagePaths)
	for index := 0; index < numberOfBatches; index++ {
		batch := batch{index, job.Options, job.ImagePaths, job.ImageNames, job.ClassesToExclude}
		this.processBatch(batch, numberOfBatches)
	}
}

//The reader may wonder why the channels of the resized images and annotations
//are copied: the reason is that the producer needs to pass the data to two
//different consumers, and this cannot be done with a single channel (a channel
//can be consumed by a single consumer).

func (this Application) processBatch(batch batch, numberOfBatches int) {
	pathsToProcess, namesToProcess := this.computePathsAndNamesToProcess(batch.Index, numberOfBatches, batch.Options, batch.ImagePaths, batch.ImageNames)
	resized := make(chan entities.ImageInfo, batch.Options.BatchSize)
	resizedCopy := make(chan entities.ImageInfo, batch.Options.BatchSize)
	resizedAnnotations := make(chan entities.Annotation, batch.Options.BatchSize)
	resizedAnnotationsCopy := make(chan entities.Annotation, batch.Options.BatchSize)
	resizedRotate := []entities.ImageInfo{}
	resizedRotateAnnotations := []entities.Annotation{}
	this.preprocessBatch(pathsToProcess,
		namesToProcess,
		resized,
		resizedCopy,
		resizedAnnotations,
		resizedAnnotationsCopy,
		batch.Options,
		batch.ClassesToExclude)
	for image := range resizedCopy {
		resizedRotate = append(resizedRotate, image)
	}
	if batch.Options.Annotated {
		for annotation := range resizedAnnotationsCopy {
			resizedRotateAnnotations = append(resizedRotateAnnotations, annotation)
		}
	}

	resizedBlur := resizedRotate
	resizedBlurAnnotations := resizedRotateAnnotations
	rotated := make(chan entities.ImageInfo)
	blurred := make(chan entities.ImageInfo)
	rotatedAnnotations := make(chan entities.Annotation, batch.Options.BatchSize)
	blurredAnnotations := make(chan entities.Annotation, batch.Options.BatchSize)
	this.augmentBatch(batch.Options,
		resizedBlur,
		resizedBlurAnnotations,
		blurred,
		blurredAnnotations,
		resizedRotate,
		resizedRotateAnnotations,
		rotated,
		rotatedAnnotations)

	if batch.Options.Annotated {
		(*this.GarbageCollector).CollectGarbage()
	}
}

func (this Application) computeNumberOfBatches(options entities.Options, imagePaths []string) int {
	if options.BatchSize > len(imagePaths) {
		panic("The batch size you've set is larger than the number of elements you intend to process. Exiting.")
	}

	numberOfBatches := int(math.Floor(float64(len(imagePaths))) / float64(options.BatchSize))
	if len(imagePaths)%options.BatchSize != 0 {
		numberOfBatches += 1
	}
	return numberOfBatches

}

func (this Application) computePathsAndNamesToProcess(index int,
	numberOfBatches int,
	options entities.Options,
	imagePaths []string,
	imageNames []string) ([]string, []string) {
	pathsToProcess := []string{}
	namesToProcess := []string{}
	//TODO:Make method to get images and paths to process
	fmt.Println("Processing batch " + strconv.Itoa(index+1) + " out of " + strconv.Itoa(numberOfBatches))
	if index == numberOfBatches-1 {
		start := index * options.BatchSize
		pathsToProcess = imagePaths[start:]
		namesToProcess = imageNames[start:]

	} else {
		start := index * options.BatchSize
		end := start + options.BatchSize
		pathsToProcess = imagePaths[start:end]
		namesToProcess = imageNames[start:end]
	}

	return pathsToProcess, namesToProcess
}

func (this Application) augmentBatch(options entities.Options,
	resizedBlur []entities.ImageInfo,
	resizedBlurAnnotations []entities.Annotation,
	blurred chan entities.ImageInfo,
	blurredAnnotations chan entities.Annotation,
	resizedRotate []entities.ImageInfo,
	resizedRotateAnnotations []entities.Annotation,
	rotated chan entities.ImageInfo,
	rotatedAnnotations chan entities.Annotation) {
	var wg sync.WaitGroup
	wg.Add(2)
	go this.rotateBatch(options,
		resizedRotate,
		resizedRotateAnnotations,
		rotated,
		rotatedAnnotations,
		&wg)

	go this.blurBatch(options,
		resizedBlur,
		resizedBlurAnnotations,
		blurred,
		blurredAnnotations,
		&wg)
	wg.Wait()

}

func (this Application) rotateBatch(options entities.Options,
	resizedRotateImages []entities.ImageInfo,
	resizedRotateAnnotations []entities.Annotation,
	rotatedImages chan entities.ImageInfo,
	rotatedAnnotations chan entities.Annotation,
	clientWg *sync.WaitGroup) {
	defer (*clientWg).Done()
	action := "rotate"
	var wg sync.WaitGroup
	augmentation, err := AugmentationsFactory(action)
	if err != nil {
		panic(err)
	}
	if options.Direction != entities.NIL_DIRECTION {
		wg.Add(2)
		go augmentation.Augment(resizedRotateImages,
			resizedRotateAnnotations,
			&wg,
			rotatedImages,
			rotatedAnnotations,
			options)
		go (*this.Converter).ConvertToJPG(rotatedImages, &wg)
		if options.Annotated {
			wg.Add(1)
			go (*this.Converter).ConvertToText(rotatedAnnotations, &wg, options.OutAnnotationType)
		}
	}
	wg.Wait()

}

func (this Application) blurBatch(options entities.Options,
	resizedBlur []entities.ImageInfo,
	resizedBlurAnnotations []entities.Annotation,
	blurred chan entities.ImageInfo,
	blurredAnnotations chan entities.Annotation,
	clientWg *sync.WaitGroup) {
	defer (*clientWg).Done()
	action := "blur"
	var wg sync.WaitGroup
	augmentation, err := AugmentationsFactory(action)
	if err != nil {
		panic(err)
	}

	if options.Sigma != 0 {
		wg.Add(2)
		go augmentation.Augment(resizedBlur,
			resizedBlurAnnotations,
			&wg,
			blurred,
			blurredAnnotations,
			options)
		go (*this.Converter).ConvertToJPG(blurred, &wg)
		if options.Annotated {
			wg.Add(1)
			go (*this.Converter).ConvertToText(blurredAnnotations, &wg, options.OutAnnotationType)
		}
	}
	wg.Wait()

}

func (this Application) preprocessBatch(pathsToProcess []string,
	namesToProcess []string,
	resized chan entities.ImageInfo,
	resizedCopy chan entities.ImageInfo,
	resizedAnnotations chan entities.Annotation,
	resizedAnnotationsCopy chan entities.Annotation,
	options entities.Options,
	classesToExclude []string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go (*this.Preprocessor).Preprocess(pathsToProcess,
		namesToProcess,
		resized,
		resizedCopy,
		resizedAnnotations,
		resizedAnnotationsCopy,
		options.InAnnotationType,
		options.Folder,
		options.Size,
		options.Annotated,
		classesToExclude,
		&wg)
	go (*this.Converter).ConvertToJPG(resized, &wg)
	if options.Annotated {
		wg.Add(1)
		go (*this.Converter).ConvertToText(resizedAnnotations, &wg, options.OutAnnotationType)
	}
	wg.Wait()
}
