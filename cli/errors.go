package cli

import (
	"fmt"
	"os"
)

func ExitWithError(err error, code int) {
	fmt.Println(err.Error())
	os.Exit(code)
}
