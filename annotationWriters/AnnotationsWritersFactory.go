package annotationWriters

import (
	"errors"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type AnnotationsWritersFactory struct {
}

func (annotationsWritersFactory AnnotationsWritersFactory) Create(annotationType entities.AnnotationType) (AnnotationWriter, error) {
	if annotationType == entities.PASCAL_VOC {
		return PascalVocWriter{}, nil
	}
	return nil, errors.New("The output annotation you've specified is not supported. Exiting.")
}
