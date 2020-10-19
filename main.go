package main

import (
	"fmt"
	"flag"
	"github.com/golobby/container"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/exclusion"
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/scan"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) > 7 {
		panic("augoment requires at most 6 arguments.")
	}

	
	var preprocessor preprocess.Preprocessor
	var converter convert.Converter
	var excluder exclusion.Excluder
	var garbageCollector collectGarbage.GarbageCollector
	var scanner scan.Scanner
	resolve(&preprocessor, &converter, &excluder, &garbageCollector, &scanner)

	folderPtr := flag.String("folder", ".", "The directory where you want to run the utility.")
	rotatePtr := flag.String("rotate", "all", "Values: 'all', 'left', 'right', 'flip'.")
	userDefinedExclusionsPtr := flag.String("user_defined_exclusions", "", "The classes you decide to exclude manually, separated by ';'.")
	inputannotationTypePtr := flag.String("in_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	outputannotationTypePtr := flag.String("out_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	sigmaPtr := flag.Int("blur", 20, "The intensity of the blur. Set to 0 for no blur.")
	batchPtr := flag.Int("batch_size", 100, "The size of the batches you intend to process asynchronously.")
	sizePtr := flag.Int("size", 464, "The height and width to which you intend to resize your images.")
	exclusionThresholdPtr := flag.Int("exclusion_threshold", 10, "The minimum number of instances of a class in order for it not to be excluded.")
	xmlPtr := flag.Bool("annotations", true, "Whether the images are annotated or not")
	flag.Parse()

	

	var options entities.Options
	side, err := convertStringToDirection(*rotatePtr)
	if err != nil {
		panic(err)
	}
	inAnnotationType, err := convertStringToAnnotationType(*inputannotationTypePtr)
	if err != nil {
		panic(err)
	}
	outAnnotationType, err := convertStringToAnnotationType(*outputannotationTypePtr)
	if err != nil {
		panic(err)
	}
	options.BatchSize = *batchPtr
	options.Side = side
	options.Sigma = float64(*sigmaPtr)
	options.Xml = *xmlPtr
	options.Size = *sizePtr
	options.ExclusionThreshold = *exclusionThresholdPtr
	options.InAnnotationType = inAnnotationType
	options.OutAnnotationType = outAnnotationType
	options.UserDefinedExclusions = getUserDefinedExclusions(*userDefinedExclusionsPtr);

	if os.Args[1] == "scan"{
		scanner.Scan(options.InAnnotationType)
		os.Exit(0);
	}

	imagePaths, imageNames := getAllPaths(*folderPtr)
	classesToExclude := excluder.GetClassesToExclude(options.ExclusionThreshold, options.UserDefinedExclusions, imageNames, options.InAnnotationType);
	fmt.Println("All images containing the following classes will be excluded: ")
	fmt.Println(classesToExclude)
	batchProcess(&options,
		&imagePaths,
		&imageNames,
		&preprocessor,
		&converter,
		&garbageCollector,
		classesToExclude)
	
}

func resolve(preprocessor *preprocess.Preprocessor,
	converter *convert.Converter,
	excluder *exclusion.Excluder,
	garbageCollector *collectGarbage.GarbageCollector,
	scanner *scan.Scanner) {
	container.Transient(func() preprocess.Preprocessor {
		return &preprocess.PreprocessingService{}
	})
	container.Transient(func() convert.Converter {
		return &convert.ConvertService{}
	})
	container.Transient(func() exclusion.Excluder {
		return &exclusion.ExclusionService{}
	})
	container.Transient(func() collectGarbage.GarbageCollector {
		return &collectGarbage.GarbageCollectionService{}
	})
	container.Transient(func() scan.Scanner{
		return &scan.ScanningService{}
	})
	container.Make(preprocessor)
	container.Make(converter)
	container.Make(excluder)
	container.Make(garbageCollector)
	container.Make(scanner)

}
