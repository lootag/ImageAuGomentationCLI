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
package blur

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)
func(this BlurService) blur(imagesToBlur []entities.ImageInfo,
	annotationsToBlur []entities.Annotation,
	blurWaitGroup *sync.WaitGroup,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	sigma float64,
	annotated bool) {
	defer (*blurWaitGroup).Done()
	var wg sync.WaitGroup
	for imageIndex := range imagesToBlur {
		wg.Add(1)
		go this.blurImageWorker(imagesToBlur[imageIndex], &wg, augmentedImages, sigma)
	}
	if annotated {
		for annotationIndex := range annotationsToBlur {
			wg.Add(1)
			go this.blurAnnotationWorker(annotationsToBlur[annotationIndex], augmentedAnnotations, &wg)
		}
	}

	wg.Wait()
	close(augmentedImages)
	close(augmentedAnnotations)
}