package rotate

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

func rotateAnnotationWorker(annotationToRotate entities.Annotation,
	rotatedAnnotations chan entities.Annotation,
	direction entities.Direction,
	rotateWaitGroup *sync.WaitGroup) {
	defer (*rotateWaitGroup).Done()
	if direction == entities.LEFT {
		var left entities.Annotation
		left.BoundingBoxes = rotateAnnotationLeft(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		left.Height = annotationToRotate.Width
		left.Width = annotationToRotate.Height
		left.FileName = annotationToRotate.FileName
		left.NewName = "rotateleft" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		left.Classes = annotationToRotate.Classes
		rotatedAnnotations <- left
	} else if direction == entities.RIGHT {
		var right entities.Annotation
		right.BoundingBoxes = rotateAnnotationRight(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		right.Height = annotationToRotate.Width
		right.Width = annotationToRotate.Height
		right.FileName = annotationToRotate.FileName
		right.NewName = "rotateright" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		right.Classes = annotationToRotate.Classes
		rotatedAnnotations <- right
	} else if direction == entities.FLIP {
		var flipped entities.Annotation
		flipped.BoundingBoxes = flipAnnotation(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		flipped.FileName = annotationToRotate.FileName
		flipped.NewName = "flipped" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		flipped.Classes = annotationToRotate.Classes
		rotatedAnnotations <- flipped
	} else if direction == entities.ALL {
		var left entities.Annotation
		left.BoundingBoxes = rotateAnnotationLeft(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		left.Height = annotationToRotate.Width
		left.Width = annotationToRotate.Height
		left.FileName = annotationToRotate.FileName
		left.NewName = "rotateleft" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		left.Classes = annotationToRotate.Classes
		rotatedAnnotations <- left
		var right entities.Annotation
		right.BoundingBoxes = rotateAnnotationRight(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		right.Height = annotationToRotate.Width
		right.Width = annotationToRotate.Height
		right.FileName = annotationToRotate.FileName
		right.NewName = "rotateright" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		right.Classes = annotationToRotate.Classes
		rotatedAnnotations <- right
		var flipped entities.Annotation
		flipped.BoundingBoxes = flipAnnotation(annotationToRotate.BoundingBoxes,
			annotationToRotate.Height,
			annotationToRotate.Width)
		flipped.FileName = annotationToRotate.FileName
		flipped.NewName = "flipped" + annotationToRotate.FileName[:len(annotationToRotate.FileName)-3] + "xml"
		flipped.Classes = annotationToRotate.Classes
		rotatedAnnotations <- flipped
	} else {
		panic("The argument provided in the rotate field is not valid. Exiting.")
	}

}

func rotateAnnotationLeft(boundingBoxesToRotate []entities.BoundingBox,
	height int,
	width int) []entities.BoundingBox {
	rotatedBoundingBoxes := []entities.BoundingBox{}
	for bndBoxIndex := 0; bndBoxIndex < len(boundingBoxesToRotate); bndBoxIndex++ {
		var rotatedBoundingBox entities.BoundingBox
		xminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmin) - (float64(width) / 2)
		yminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymin) - (float64(height) / 2)
		xmaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmax) - (float64(width) / 2)
		ymaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymax) - (float64(height) / 2)
		newXminTranslated := yminTranslated
		newYminTranslated := -xmaxTranslated
		newXmaxTranslated := ymaxTranslated
		newYmaxTranslated := -xminTranslated
		rotatedBoundingBox.Xmin = int(newXminTranslated + float64(width)/2)
		rotatedBoundingBox.Ymin = int(newYminTranslated + float64(height)/2)
		rotatedBoundingBox.Xmax = int(newXmaxTranslated + float64(width)/2)
		rotatedBoundingBox.Ymax = int(newYmaxTranslated + float64(height)/2)
		rotatedBoundingBoxes = append(rotatedBoundingBoxes, rotatedBoundingBox)
	}
	return rotatedBoundingBoxes
}

func rotateAnnotationRight(boundingBoxesToRotate []entities.BoundingBox,
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

func flipAnnotation(boundingBoxesToRotate []entities.BoundingBox,
	height int,
	width int) []entities.BoundingBox {
	rotatedBoundingBoxes := []entities.BoundingBox{}
	for bndBoxIndex := 0; bndBoxIndex < len(boundingBoxesToRotate); bndBoxIndex++ {
		var rotatedBoundingBox entities.BoundingBox
		xminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmin) - (float64(width) / 2)
		yminTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymin) - (float64(height) / 2)
		xmaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Xmax) - (float64(width) / 2)
		ymaxTranslated := float64(boundingBoxesToRotate[bndBoxIndex].Ymax) - (float64(height) / 2)
		newXminTranslated := -xmaxTranslated
		newYminTranslated := -ymaxTranslated
		newXmaxTranslated := -xminTranslated
		newYmaxTranslated := -yminTranslated
		rotatedBoundingBox.Xmin = int(newXminTranslated + float64(width)/2)
		rotatedBoundingBox.Ymin = int(newYminTranslated + float64(height)/2)
		rotatedBoundingBox.Xmax = int(newXmaxTranslated + float64(width)/2)
		rotatedBoundingBox.Ymax = int(newYmaxTranslated + float64(height)/2)
		rotatedBoundingBoxes = append(rotatedBoundingBoxes, rotatedBoundingBox)
	}
	return rotatedBoundingBoxes
}
