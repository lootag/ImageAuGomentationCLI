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

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this RotateService) rotateBoundingBoxRight(boundingBoxesToRotate []entities.BoundingBox,
	height int,
	width int) []entities.BoundingBox {
	rotatedBoundingBoxes := []entities.BoundingBox{}
	for bndBoxIndex := 0; bndBoxIndex < len(boundingBoxesToRotate); bndBoxIndex++ {
		var rotatedBoundingBox entities.BoundingBox
		xminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmin) - (float64(width) / 2)
		yminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymin) - (float64(height) / 2)
		xmaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmax) - (float64(width) / 2)
		ymaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymax) - (float64(height) / 2)
		newXminTranslated := -ymaxTranslated
		newYminTranslated := xminTranslated
		newXmaxTranslated := -yminTranslated
		newYmaxTranslated := xmaxTranslated
		rotatedBoundingBox.Xmin = int(newXminTranslated + float64(width)/2)
		rotatedBoundingBox.Ymin = int(newYminTranslated + float64(height)/2)
		rotatedBoundingBox.Xmax = int(newXmaxTranslated + float64(width)/2)
		rotatedBoundingBox.Ymax = int(newYmaxTranslated + float64(height)/2)
		rotatedBoundingBoxes = append(rotatedBoundingBoxes, rotatedBoundingBox)
	}
	return rotatedBoundingBoxes
}