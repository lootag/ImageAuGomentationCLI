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
package rotate

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type RotateService struct {
}

func (rotateService RotateService) Augment(imagesToAugment []entities.ImageInfo,
	annotationsToAugment []entities.Annotation,
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
		options.Direction,
		augmentedImages,
		augmentedAnnotations,
		options.Annotated)
	wg.Wait()
}

func rotate(imagesToRotate []entities.ImageInfo,
	annotationsToRotate []entities.Annotation,
	augmentWaitGroup *sync.WaitGroup,
	direction entities.Direction,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	annotated bool) {
	defer (*augmentWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := range imagesToRotate {
		wg.Add(1)
		go rotateImageWorker(imagesToRotate[imageIndex], &wg, direction, augmentedImages)
	}

	if annotated {
		for annotationIndex := range annotationsToRotate {
			wg.Add(1)
			go rotateAnnotationWorker(annotationsToRotate[annotationIndex],
				augmentedAnnotations,
				direction,
				&wg)
		}
	}
	wg.Wait()
	close(augmentedImages)
	close(augmentedAnnotations)
}
