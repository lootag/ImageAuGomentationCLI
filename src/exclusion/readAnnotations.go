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
package exclusion

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
)

func(this ExclusionService) readAnnotations(annotationType entities.AnnotationType,
	excludeWaitGroup *sync.WaitGroup,
	annotationPaths chan string,
	annotationsToGroup chan entities.Annotation) {
	defer (*excludeWaitGroup).Done()
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotationPath := range annotationPaths {
		annotation := annotationReader.ReadSync(annotationPath)
		annotationsToGroup <- annotation
	}
	close(annotationsToGroup)
}