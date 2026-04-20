//go:build arm64

package alerts

import (
	"fmt"
	"os/exec"
)

func ExecSend(num string, msg string) error {
	cmd := exec.Command("termux-sms-send", "-n", num, msg)

	// Выполнение и получение вывода
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}
