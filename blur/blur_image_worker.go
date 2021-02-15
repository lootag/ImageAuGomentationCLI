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

	"github.com/disintegration/imaging"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this BlurService) blurImageWorker(imageToBlur entities.ImageInfo,
	blurWaitGroup *sync.WaitGroup,
	augmented chan entities.ImageInfo,
	sigma float64) {
	defer (*blurWaitGroup).Done()
	var blurredImage entities.ImageInfo
	blurredImage.ImageSource = imaging.Blur(imageToBlur.ImageSource, sigma)
	blurredImage.OriginalFileName = imageToBlur.OriginalFileName
	blurredImage.NewName = "blur" + imageToBlur.OriginalFileName
	augmented <- blurredImage
}
