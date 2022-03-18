package drone_git_push

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
)

type Plugin struct {
	LocalDir      string
	CommitMessage string
	Branch        string
}

func (p *Plugin) check() error {
	if p.LocalDir == "" {
		return errors.New("local_dir can not nil")
	}

	return nil
}

func (p *Plugin) run(c *exec.Cmd) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cacheDir := path.Join(cwd, p.LocalDir)

	c.Env = os.Environ()
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Dir = cacheDir

	return c.Run()
}

func (p *Plugin) Exec() error {
	if err := p.check(); err != nil {
		return err
	}

	// gitAdd := exec.Command("git", "add", ".")
	// gitCommit := exec.Command("git", "commit", "-a", "-m", p.CommitMessage)
	gitPush := exec.Command("git", "push", "-u", "origin", p.Branch)

	// log.Println("Add file contents to the index")
	// if err := p.run(gitAdd); err != nil {
	// 	return err
	// }

	// log.Println("Record changes to the repository")
	// if err := p.run(gitCommit); err != nil {
	// 	return err
	// }

	log.Println("Update remote refs along with associated objects")
	if err := p.run(gitPush); err != nil {
		return err
	}

	return nil
}
