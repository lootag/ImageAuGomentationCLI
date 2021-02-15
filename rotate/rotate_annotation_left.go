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

func(this RotateService) rotateAnnotationLeft(annotationToRotate entities.Annotation,
	rotatedAnnotations chan entities.Annotation) {
	var left entities.Annotation
	left.BoundingBoxes = this.rotateBoundingBoxLeft(annotationToRotate.BoundingBoxes,
		annotationToRotate.Height,
		annotationToRotate.Width)
	left.Height = annotationToRotate.Width
	left.Width = annotationToRotate.Height
	left.Depth = annotationToRotate.Depth
	left.FileName = annotationToRotate.FileName
	left.NewName = commons.RenameAnnotation(annotationToRotate.FileName, "rotateleft")
	left.Classes = annotationToRotate.Classes
	rotatedAnnotations <- left
}