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

	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
)

func(this PreprocessingService) readAnnotations(annotationType entities.AnnotationType,
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