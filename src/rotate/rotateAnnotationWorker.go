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

func(this RotateService) rotateAnnotationWorker(annotationToRotate entities.Annotation,
	rotatedAnnotations chan entities.Annotation,
	direction entities.Direction,
	rotateWaitGroup *sync.WaitGroup) {
	defer (*rotateWaitGroup).Done()
	if direction == entities.LEFT {
		//TODO: All of this needs to be put in a method
		this.rotateAnnotationLeft(annotationToRotate, rotatedAnnotations)
	} else if direction == entities.RIGHT {
		this.rotateAnnotationRight(annotationToRotate, rotatedAnnotations)
	} else if direction == entities.FLIP {
		this.flipAnnotation(annotationToRotate, rotatedAnnotations)
	} else if direction == entities.ALL {
		this.rotateAnnotationLeft(annotationToRotate, rotatedAnnotations)
		this.rotateAnnotationRight(annotationToRotate, rotatedAnnotations)
		this.flipAnnotation(annotationToRotate, rotatedAnnotations)

	} else {
		panic("The argument provided in the rotate field is not valid. Exiting.")
	}

}

