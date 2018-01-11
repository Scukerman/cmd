// +build !windows

package cmd

import "syscall"

func (c *Cmd) killProcess() error {
	return syscall.Kill(-c.status.PID, syscall.SIGTERM)
}

func (c *Cmd) setpgid() {
	c.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}
