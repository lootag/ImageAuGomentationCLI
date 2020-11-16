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
package preprocess

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this PreprocessingService) resizeImages(validatedImages chan entities.ImageInfo,
	resizedImages chan entities.ImageInfo,
	resizedImagesCopy chan entities.ImageInfo,
	size int,
	preprocessWaitGroup *sync.WaitGroup) {
	defer (*preprocessWaitGroup).Done()
	var wg sync.WaitGroup
	for image := range validatedImages {
		wg.Add(1)
		go this.resizeImageWorker(image, resizedImages, resizedImagesCopy, &wg, size)
	}
	wg.Wait()
	close(resizedImages)
	close(resizedImagesCopy)

}