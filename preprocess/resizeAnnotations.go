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
package preprocess

import (
	"sync"
	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this PreprocessingService) resizeAnnotations(annotationsToResize chan entities.Annotation,
	resizedAnnotations chan entities.Annotation,
	resizeAnnotationsCopy chan entities.Annotation,
	newSize int,
	classesToExclude []string,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for annotation := range annotationsToResize {
		intersection := commons.IntersectStringArrays(annotation.Classes, classesToExclude)
		if len(intersection) == 0 {
			wg.Add(1)
			go this.resizeAnnotationWorker(annotation, resizedAnnotations, resizeAnnotationsCopy, newSize, &wg)
		}
	}
	wg.Wait()
	close(resizedAnnotations)
	close(resizeAnnotationsCopy)
}