package main

import (
	"io/ioutil"
)

func getAllPaths(folder string) ([]string, []string) {
	root := folder + "/Images"
	fileInfos, err := ioutil.ReadDir(root)
	var paths []string
	var names []string
	if err != nil {
		panic(err)
	}
	for infoIndex := range fileInfos {
		names = append(names, fileInfos[infoIndex].Name())
		paths = append(paths, root+"/"+fileInfos[infoIndex].Name())
	}
	return paths, names
}
