package annotationReaders

import (
	"errors"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type AnnotationReadersFactory struct {
}

func (annotationReadersFactory AnnotationReadersFactory) Create(annotationType entities.AnnotationType) (AnnotationReader, error) {

	if annotationType == entities.PASCAL_VOC {
		return PascalVocReader{}, nil
	} else if annotationType == entities.GOOGLE {
		return GoogleReader{}, nil
	}

	return nil, errors.New("The annotation type you've specified is not supported. Exiting.")
}
