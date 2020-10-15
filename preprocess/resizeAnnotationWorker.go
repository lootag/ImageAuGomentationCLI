package preprocess

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"math"
	"sync"
)

func resizeAnnotationWorker(annotationToResize entities.Annotation,
	resizedAnnotations chan entities.Annotation,
	resizedAnnotationsCopy chan entities.Annotation,
	newSize int,
	resizeAnnotationsWaitGroup *sync.WaitGroup,
) {
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
	resizedAnnotation.NewName = "resized" + annotationToResize.FileName[:len(annotationToResize.FileName)-3] + "xml"
	resizedAnnotations <- resizedAnnotation
	resizedAnnotationsCopy <- resizedAnnotation

}
