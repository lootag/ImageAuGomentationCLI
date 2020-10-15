package main;

import(
	"github.com/lootag/ImageAuGomentationCLI/entities";
	"errors";
)

func convertStringToAnnotationType(argument string) (entities.AnnotationType, error){
	switch argument {
	case "pascalvoc":
		return entities.PASCAL_VOC, nil;
	}
	return entities.NIL_ANNOTATION, errors.New("The rotation you've specified is not supported.");
}