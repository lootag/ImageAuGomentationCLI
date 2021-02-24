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

	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type ExclusionService struct {
}

func (this ExclusionService) GetClassesToExclude(request entities.ExcludeRequest) entities.ExcludeResponse {
	var wg sync.WaitGroup
	annotationPaths := make(chan string, len(request.ImageNames))
	annotationsToGroup := make(chan entities.Annotation, len(request.ImageNames))
	wg.Add(2)
	go this.getAnnotationPathsFromImageNames(request.ImageNames, request.Folder, annotationPaths, &wg)
	go this.readAnnotations(request.AnnotationType, &wg, annotationPaths, annotationsToGroup)
	wg.Wait()
	countMap := this.getCountMap(annotationsToGroup)
	classesToExclude := this.excludeClassesWithCountBelowThreshold(request.ExclusionThreshold, countMap)
	for _, userDefinedExclusion := range request.UserDefinedExclusions {
		if !commons.StringArrayContains(classesToExclude, userDefinedExclusion) {
			classesToExclude = append(classesToExclude, userDefinedExclusion)
		}
	}
	return entities.ExcludeResponse{classesToExclude}
}
