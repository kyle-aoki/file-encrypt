package path

import "strings"

type Path struct {
	file string
	dir  string
}

func GetPath(rawPath string) {
	strings.Split(rawPath, "/")
	
}
