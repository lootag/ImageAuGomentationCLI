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
	"github.com/lootag/ImageAuGomentationCLI/collectGarbage"
	"github.com/lootag/ImageAuGomentationCLI/convert"
	"github.com/lootag/ImageAuGomentationCLI/exclusion"
	"github.com/lootag/ImageAuGomentationCLI/orchestration"
	"github.com/lootag/ImageAuGomentationCLI/preprocess"
	"github.com/lootag/ImageAuGomentationCLI/scan"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	validateNumberOfCLIArguments()
	//These are all the services which the application needs
	var preprocessor preprocess.Preprocessor
	var converter convert.Converter
	var excluder exclusion.Excluder
	var garbageCollector collectGarbage.GarbageCollector
	var scanner scan.Scanner

	//Here the dependencies are resolved
	orchestration.ResolveDependencies(&preprocessor, &converter, &excluder, &garbageCollector, &scanner)
	options := orchestration.GetCommandLineOptions()

	//The dependencies are injected in the application class
	app := orchestration.Application{&preprocessor, &converter, &excluder, &garbageCollector, &scanner}
	app.Run(options)
}

func validateNumberOfCLIArguments() {
	numberOfCLIArguments := 12
	if len(os.Args) > numberOfCLIArguments {
		panic("augoment requires at most 11 arguments.")
	}
}
