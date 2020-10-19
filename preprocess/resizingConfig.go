package preprocess;

import(
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type resizingConfig struct{
	Xmls bool;
	CheckedImages chan entities.Annotation;
	
}