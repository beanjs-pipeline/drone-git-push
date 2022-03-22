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
	log.Printf("cache dir: %s", cacheDir)

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

	log.Printf("git config")
	if err := p.run(exec.Command("git", "config", "--global", "user.email", "drone-ci@gmail.com")); err != nil {
		return err
	}
	if err := p.run(exec.Command("git", "config", "--global", "user.name", "drone-ci")); err != nil {
		return err
	}

	gitAdd := exec.Command("git", "add", ".")
	gitCommit := exec.Command("git", "commit", "-a", "-m", p.CommitMessage)
	gitPush := exec.Command("git", "push", "-u", "origin", p.Branch)

	log.Println("Add file contents to the index")
	if err := p.run(gitAdd); err != nil {
		return err
	}

	log.Println("Record changes to the repository")
	if err := p.run(gitCommit); err != nil {
		log.Println("There are no changes to commit.")
		return nil
	}

	log.Println("Update remote refs along with associated objects")
	if err := p.run(gitPush); err != nil {
		return err
	}

	return nil
}
