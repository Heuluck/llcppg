package thirddep2

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

type TypeThirdDep2 struct {
	A c.Int
	B c.Int
}
