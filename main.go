package main

import (
	"flag"
	"github.com/golobby/container"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
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
	resolve(&preprocessor, &converter)

	folderPtr := flag.String("folder", ".", "The directory where you want to run the utility.")
	rotatePtr := flag.String("rotate", "all", "Values: 'all', 'left', 'right', 'flip'.")
	inputannotationTypePtr := flag.String("in_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	outputannotationTypePtr := flag.String("out_annotationtype", "pascalvoc", "Values: 'pascalvoc'")
	sigmaPtr := flag.Int("blur", 20, "The intensity of the blur. Set to 0 for no blur.")
	batchPtr := flag.Int("batch_size", 100, "The size of the batches you intend to process asynchronously.")
	sizePtr := flag.Int("size", 464, "The height and width to which you intend to resize your images.")
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
	options.InAnnotationType = inAnnotationType
	options.OutAnnotationType = outAnnotationType

	imagePaths, imageNames := getAllPaths(*folderPtr)

	batchProcess(&options,
		&imagePaths,
		&imageNames,
		&preprocessor,
		&converter)
}

func resolve(preprocessor *preprocess.Preprocessor,
	converter *convert.Converter) {
	container.Transient(func() preprocess.Preprocessor {
		return &preprocess.PreprocessingService{}
	})
	container.Transient(func() convert.Converter {
		return &convert.ConvertService{}
	})
	container.Make(preprocessor)
	container.Make(converter)

}
