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
package convert

import (
	"os"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/annotationWriters"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type ConvertService struct {
}

func (this ConvertService) ConvertToJPG(toConvert chan entities.ImageInfo,
	mainWaitGroup *sync.WaitGroup) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	err := os.Mkdir("./AugmentedImages", 0755)
	if err != nil {
	}
	for image := range toConvert {
		wg.Add(1)
		go this.convertToJPGWorker(image, &wg)
	}
	wg.Wait()
}

func (this ConvertService) ConvertToText(annotationToConvert chan entities.Annotation,
	mainWaitGroup *sync.WaitGroup,
	annotationType entities.AnnotationType) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	var factory annotationWriters.AnnotationsWritersFactory
	err := os.Mkdir("./AugmentedAnnotations", 0755)
	writer, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotation := range annotationToConvert {
		wg.Add(1)
		go writer.Write(annotation, &wg)
	}
	wg.Wait()
}
