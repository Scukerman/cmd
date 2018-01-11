// +build windows

package cmd

func (c *Cmd) killProcess() error {
	return c.cmd.Process.Kill()
}

func (c *Cmd) setpgid() {}
