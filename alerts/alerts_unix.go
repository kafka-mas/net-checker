//go:build amd64

package alerts

import (
	"fmt"
)

func ExecSend(num string, msg string) error {
	fmt.Println(msg)

	return nil
}
