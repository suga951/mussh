package player

import (
	"os/exec"
	"syscall"
)

type Player struct {
	cmd *exec.Cmd
}

func (p *Player) Play(url string) error {
	p.Stop()

	cmd := exec.Command("mpv", "--no-video", "--quiet", url)
	if err := cmd.Start(); err != nil {
		return err
	}
	p.cmd = cmd
	return nil
}

func (p *Player) Stop() {
	if p.cmd != nil && p.cmd.Process != nil {
		_ = p.cmd.Process.Signal(syscall.SIGTERM)
		go func(cmd *exec.Cmd) {
			cmd.Wait()
		}(p.cmd)
		p.cmd = nil
	}
}
