package orchestration

import (
	"flag"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func GetCommandLineOptions() entities.Options{
	folderPtr := flag.String("folder", ".", "The directory where you want to run the utility.")
	rotatePtr := flag.String("rotate", "all", "The rotations you want to apply to the images. Set to 'skip' in order not to perform any rotations. Values: 'all', 'left', 'right', 'flip', 'skip'.")
	userDefinedExclusionsPtr := flag.String("user_defined_exclusions", "", "The classes you decide to exclude manually, separated by ';'.")
	inputannotationTypePtr := flag.String("in_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	outputannotationTypePtr := flag.String("out_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	sigmaPtr := flag.Int("blur", 20, "The intensity of the blur. Set to 0 for no blur.")
	batchPtr := flag.Int("batch_size", 50, "The size of the batches you intend to process asynchronously.")
	sizePtr := flag.Int("size", 464, "The height and width to which you intend to resize your images.")
	exclusionThresholdPtr := flag.Int("exclusion_threshold", 10, "The minimum number of instances of a class in order for it not to be excluded.")
	xmlPtr := flag.Bool("annotations", true, "Whether the images are annotated or not")
	scanPtr := flag.Bool("scan", false, "Whether you want to scan the data")
	flag.Parse()

	var options entities.Options
	side, err := ConvertStringToDirection(*rotatePtr)
	if err != nil {
		panic(err)
	}
	inAnnotationType, err := ConvertStringToAnnotationType(*inputannotationTypePtr)
	if err != nil {
		panic(err)
	}
	outAnnotationType, err := ConvertStringToAnnotationType(*outputannotationTypePtr)
	if err != nil {
		panic(err)
	}
	options.Folder = *folderPtr
	options.BatchSize = *batchPtr
	options.Direction = side
	options.Sigma = float64(*sigmaPtr)
	options.Annotated = *xmlPtr
	options.Size = *sizePtr
	options.ExclusionThreshold = *exclusionThresholdPtr
	options.InAnnotationType = inAnnotationType
	options.OutAnnotationType = outAnnotationType
	options.UserDefinedExclusions = GetUserDefinedExclusions(*userDefinedExclusionsPtr)
	options.Scan = *scanPtr

	return options
}