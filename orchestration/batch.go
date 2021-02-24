package orchestration

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type batch struct {
	Index            int
	Options          entities.Options
	ImagePaths       []string
	ImageNames       []string
	ClassesToExclude []string
}
