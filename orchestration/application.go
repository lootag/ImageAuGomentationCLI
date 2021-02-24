package orchestration

import (
	"fmt"
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/exclusion"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/scan"
	"os"
)

type Application struct {
	Preprocessor     *preprocess.Preprocessor
	Converter        *convert.Converter
	Excluder         *exclusion.Excluder
	GarbageCollector *collectGarbage.GarbageCollector
	Scanner          *scan.Scanner
}

func (this Application) Run(options entities.Options) {
	//There are two paths that the application can take: it either scans the
	//annotations (in case the user inputs "true" as the value of the scan
	//parameter), or it augments the images and/or annotations.
	if options.Scan {
		(*this.Scanner).Scan(options.InAnnotationType, options.Folder)
		os.Exit(0)
	}

	imagePaths, imageNames := commons.GetAllImagePathsAndNames(options.Folder)
	classesToExclude := []string{}
	excludeRequest := entities.ExcludeRequest{options.ExclusionThreshold, options.UserDefinedExclusions, imageNames, options.Folder, options.InAnnotationType}
	if options.Annotated {
		classesToExclude = (*this.Excluder).GetClassesToExclude(excludeRequest).ClassesToExclude
	}
	fmt.Println("All images containing the following classes will be excluded: ")
	fmt.Println(classesToExclude)
	job := augmentationJob{options, imagePaths, imageNames, classesToExclude}
	this.runAugmentationJob(job)

}
