package tests

import (
	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
	"testing"
)

func TestPascalVocReader(t *testing.T) {
	var expectedAnnotation entities.Annotation
	expectedAnnotation.Class = "adidas_1"
	expectedAnnotation.BoundingBox.Xmin = 408
	expectedAnnotation.BoundingBox.Ymin = 279
	expectedAnnotation.BoundingBox.Xmax = 477
	expectedAnnotation.BoundingBox.Ymax = 359
	var wg sync.WaitGroup
	fileName := "49901162003.xml"
	inputAnnotations := make(chan entities.Annotation, 1)
	var reader annotationReaders.PascalVocReader
	wg.Add(1)
	go reader.Read(fileName, inputAnnotations, &wg)
	wg.Wait()
	close(inputAnnotations)
	outputArray := []entities.Annotation{}
	for annotation := range inputAnnotations {
		outputArray = append(outputArray, annotation)
	}
	actualAnnotation := outputArray[0]
	if actualAnnotation.Class != expectedAnnotation.Class {
		t.Errorf("Not the class I was expecting")
	}
	if actualAnnotation.BoundingBox.Xmin != expectedAnnotation.BoundingBox.Xmin {
		t.Errorf("Not the Xmin I was expecting")
	}
	if actualAnnotation.BoundingBox.Ymin != expectedAnnotation.BoundingBox.Ymin {
		t.Errorf("Not the Ymin I was expecting")
	}
	if actualAnnotation.BoundingBox.Xmax != expectedAnnotation.BoundingBox.Xmax {
		t.Errorf("Not the Xmax I was expecting")
	}
	if actualAnnotation.BoundingBox.Ymax != expectedAnnotation.BoundingBox.Ymax {
		t.Errorf("Not the Ymax I was expecting")
	}

}
