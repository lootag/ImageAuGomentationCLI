package tests

import (
	"github.com/lootag/ImageAuGomentationCLI/annotationWriters"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"os"
	"sync"
	"testing"
)

func TestPascalVocWriter(t *testing.T) {
	var annotation entities.Annotation
	var writer annotationWriters.PascalVocWriter
	err := os.Mkdir("./AugmentedAnnotations", 0755)
	if err != nil {

	}
	annotation.FileName = "TestAnnotation.xml"
	annotation.Width = 123
	annotation.Height = 145
	annotation.Depth = 3
	annotation.BoundingBox.Xmin = 100
	annotation.BoundingBox.Ymin = 80
	annotation.BoundingBox.Xmax = 120
	annotation.BoundingBox.Ymax = 90
	annotation.Class = "adidas_1"
	var wg sync.WaitGroup
	wg.Add(1)
	go writer.Write(annotation, &wg)
	wg.Wait()
	//TODO Assert fields.

}
