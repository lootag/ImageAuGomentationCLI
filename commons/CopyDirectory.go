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

package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)
func CopyDirectory(scrDir, dest string) error {
    entries, err := ioutil.ReadDir(scrDir)
    if err != nil {
        return err
    }
    for _, entry := range entries {
        sourcePath := filepath.Join(scrDir, entry.Name())
        destPath := filepath.Join(dest, entry.Name())

        fileInfo, err := os.Stat(sourcePath)
        if err != nil {
            return err
        }

        stat, ok := fileInfo.Sys().(*syscall.Stat_t)
        if !ok {
            return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
        }

        switch fileInfo.Mode() & os.ModeType{
        case os.ModeDir:
            if err := CreateIfNotExists(destPath, 0755); err != nil {
                return err
            }
            if err := CopyDirectory(sourcePath, destPath); err != nil {
                return err
            }
        case os.ModeSymlink:
            if err := CopySymLink(sourcePath, destPath); err != nil {
                return err
            }
        default:
            if err := Copy(sourcePath, destPath); err != nil {
                return err
            }
        }

        if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
            return err
        }

        isSymlink := entry.Mode()&os.ModeSymlink != 0
        if !isSymlink {
            if err := os.Chmod(destPath, entry.Mode()); err != nil {
                return err
            }
        }
    }
    return nil
}