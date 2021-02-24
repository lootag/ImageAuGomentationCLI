package orchestration

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type augmentationJob struct {
	Options          entities.Options
	ImagePaths       []string
	ImageNames       []string
	ClassesToExclude []string
}
