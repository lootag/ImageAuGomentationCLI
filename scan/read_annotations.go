package scan

import (

	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this ScanningService) readAnnotations(annotationType entities.AnnotationType,
	annotationsToRead []string) []entities.Annotation {
	annotationsToGroup := []entities.Annotation{}
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotationPathIndex := range annotationsToRead {
		annotation := annotationReader.ReadSync(annotationsToRead[annotationPathIndex])
		annotationsToGroup = append(annotationsToGroup, annotation)
	}
	return annotationsToGroup
}