package hello

import (
	"fmt"

	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

func newSubAction() {
	fmt.Print("Hello, {{ .Project.Name }}")

	cmdutils.IsVersionCmd()
}
