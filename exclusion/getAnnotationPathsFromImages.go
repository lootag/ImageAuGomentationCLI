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
package exclusion

import (
	"sync"
)

func(this ExclusionService) getAnnotationPathsFromImageNames(imageNames []string,
	folder string,
	annotationPaths chan string,
	excludeWaitGroup *sync.WaitGroup) {
	defer (*excludeWaitGroup).Done()
	var wg sync.WaitGroup
	for imageNameIndex := range imageNames {
		wg.Add(1)
		go this.getAnnotationPathsFromImageNamesWorker(imageNames[imageNameIndex], folder, annotationPaths, &wg)
	}
	wg.Wait()
	close(annotationPaths)
}