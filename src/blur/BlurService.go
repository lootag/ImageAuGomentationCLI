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
package blur

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type BlurService struct {
}

func (this BlurService) Augment(imagesToAugment []entities.ImageInfo,
	annotationsToAugment []entities.Annotation,
	mainWaitGroup *sync.WaitGroup,
	augmentedImages chan entities.ImageInfo,
	augmentedAnnotations chan entities.Annotation,
	options entities.Options) {
	defer (*mainWaitGroup).Done()
	var wg sync.WaitGroup
	wg.Add(1)
	go this.blur(imagesToAugment,
		annotationsToAugment,
		&wg,
		augmentedImages,
		augmentedAnnotations,
		options.Sigma,
		options.Annotated)
	wg.Wait()
}


