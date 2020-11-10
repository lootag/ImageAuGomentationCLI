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
	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this RotateService) rotateAnnotationRight(annotationToRotate entities.Annotation,
	rotatedAnnotations chan entities.Annotation) {
	var right entities.Annotation
	right.BoundingBoxes = this.rotateBoundingBoxRight(annotationToRotate.BoundingBoxes,
		annotationToRotate.Height,
		annotationToRotate.Width)
	right.Height = annotationToRotate.Width
	right.Width = annotationToRotate.Height
	right.Depth = annotationToRotate.Depth
	right.FileName = annotationToRotate.FileName
	right.NewName = commons.RenameAnnotation(annotationToRotate.FileName, "rotateright")
	right.Classes = annotationToRotate.Classes
	rotatedAnnotations <- right
}