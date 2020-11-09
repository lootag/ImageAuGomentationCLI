package orchestration

import (
	"github.com/golobby/container"
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/exclusion"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/scan"
)

func ResolveDependencies(preprocessor *preprocess.Preprocessor,
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
	container.Transient(func() scan.Scanner {
		return &scan.ScanningService{}
	})
	container.Make(preprocessor)
	container.Make(converter)
	container.Make(excluder)
	container.Make(garbageCollector)
	container.Make(scanner)

}