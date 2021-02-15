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
	"math"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this PreprocessingService) resizeAnnotationWorker(annotationToResize entities.Annotation,
	resizedAnnotations chan entities.Annotation,
	resizedAnnotationsCopy chan entities.Annotation,
	newSize int,
	resizeAnnotationsWaitGroup *sync.WaitGroup) {
	defer (*resizeAnnotationsWaitGroup).Done()
	resizedAnnotation := annotationToResize
	resizedAnnotation.Height = newSize
	resizedAnnotation.Width = newSize
	widthRatio := float64(annotationToResize.Width) / float64(newSize)
	heightRatio := float64(annotationToResize.Height) / float64(newSize)
	for bndBoxIndex := 0; bndBoxIndex < len(annotationToResize.BoundingBoxes); bndBoxIndex++ {
		xmin := int(math.Floor(float64(annotationToResize.BoundingBoxes[bndBoxIndex].Xmin) / widthRatio))
		ymin := int(math.Floor(float64(annotationToResize.BoundingBoxes[bndBoxIndex].Ymin) / heightRatio))
		xmax := int(math.Floor(float64(annotationToResize.BoundingBoxes[bndBoxIndex].Xmax) / widthRatio))
		ymax := int(math.Floor(float64(annotationToResize.BoundingBoxes[bndBoxIndex].Ymax) / heightRatio))
		resizedAnnotation.BoundingBoxes[bndBoxIndex].Xmin = xmin
		resizedAnnotation.BoundingBoxes[bndBoxIndex].Ymin = ymin
		resizedAnnotation.BoundingBoxes[bndBoxIndex].Xmax = xmax
		resizedAnnotation.BoundingBoxes[bndBoxIndex].Ymax = ymax
	}
	resizedAnnotation.NewName = this.renameAnnotation(annotationToResize.FileName)
	resizedAnnotations <- resizedAnnotation
	resizedAnnotationsCopy <- resizedAnnotation

}
