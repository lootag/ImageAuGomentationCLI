package exclusion;

import (
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type Excluder interface {
	GetClassesToExclude(exclusionThreshold int, 
		userDefinedExclusions []string,
		imageNames []string,
		annotationType entities.AnnotationType) []string;
}