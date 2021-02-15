/*
This file is part of ImageAuGomentationCLI.

ImageAuGomentationCLI is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

ImageAuGomentationCLI is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ImageAuGomentationCLI.  If not, see <https://www.gnu.org/licenses/>.
*/

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