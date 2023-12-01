package rdx2go

import (
	"go/token"
	"os"

	"github.com/goplus/gop/x/build"
)

const (
	rdxpkg = "github.com/realdream-ai/rdx"
)

func init() {
	build.RegisterClassFileType(".rdx", "Game", []*build.Class{{
		Ext:   ".rdx",
		Class: "Sprite"},
	}, rdxpkg)
}

func Main(outputFile, dir string) (err error) {
	fset := token.NewFileSet()
	impl := NewImporter(fset)
	ctx := build.NewContext(impl, fset)
	data, err := ctx.BuildDir(dir)
	if err != nil {
		return
	}
	return os.WriteFile(outputFile, data, 0644)
}
