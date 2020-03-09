package main

import "os/exec"

func main() {
	cmd := exec.Command("stress", "--verbose", "--cpu", cpu, "--timeout", timeout)

	// TODO: 优化输出
	cmd.Run()

	cmd.Start()

	cmd.Output()
}
