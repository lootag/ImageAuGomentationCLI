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

package main

import (
	"fmt"
	"os"
	"sync"
	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/orchestration"
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/exclusion"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/scan"
)

var wg sync.WaitGroup

func main() {
	numberOfCLIArguments := 12
	if len(os.Args) > numberOfCLIArguments {
		panic("augoment requires at most 11 arguments.")
	}

	var preprocessor preprocess.Preprocessor
	var converter convert.Converter
	var excluder exclusion.Excluder
	var garbageCollector collectGarbage.GarbageCollector
	var scanner scan.Scanner
	orchestration.ResolveDependencies(&preprocessor, &converter, &excluder, &garbageCollector, &scanner)
	options := orchestration.GetCommandLineOptions()

	if options.Scan {
		scanner.Scan(options.InAnnotationType, options.Folder)
		os.Exit(0)
	}

	imagePaths, imageNames := commons.GetAllImagePathsAndNames(options.Folder)
	classesToExclude := []string{};
	if options.Annotated {
		classesToExclude = excluder.GetClassesToExclude(options.ExclusionThreshold, 
														options.UserDefinedExclusions, 
														imageNames, 
														options.Folder, 
														options.InAnnotationType)
	}
	fmt.Println("All images containing the following classes will be excluded: ")
	fmt.Println(classesToExclude)
	orchestration.BatchProcess(&options,
		imagePaths,
		imageNames,
		&preprocessor,
		&converter,
		&garbageCollector,
		classesToExclude)

}

