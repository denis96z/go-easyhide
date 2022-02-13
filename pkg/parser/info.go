package parser

type Info struct {
	Pkg   PkgInfo
	Types []TypeInfo

	GeneratedFilePath string
}

type PkgInfo struct {
	Name string
	Path string
}

type TypeInfo struct {
	Name string
}
