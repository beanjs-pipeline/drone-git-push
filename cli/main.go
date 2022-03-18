package main

import (
	"log"
	"os"

	droneGitPush "github.com/beanjs-pipeline/drone-git-push"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "Drone git push"
	app.Usage = "Update remote refs along with associated objects"
	app.Copyright = "@ 2022 beanjs"
	app.Authors = []cli.Author{
		{Name: "beanjs", Email: "502554248@qq.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "branch",
			Usage:  "branch",
			EnvVar: "PLUGIN_BRANCH",
			Value:  "master",
		},
		cli.StringFlag{
			Name:   "local_dir",
			Usage:  "local_dir",
			EnvVar: "PLUGIN_LOCAL_DIR",
		},
		cli.StringFlag{
			Name:   "commit_message",
			Usage:  "commit_message",
			EnvVar: "PLUGIN_COMMIT_MESSAGE",
			Value:  "commit from ci",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Fatal: %v", err)
	}
}

func run(c *cli.Context) {
	p := droneGitPush.Plugin{
		Branch:        c.String("branch"),
		LocalDir:      c.String("local_dir"),
		CommitMessage: c.String("commit_message"),
	}

	if err := p.Exec(); err != nil {
		log.Fatalf("Run Error: %v", err)
		return
	}

	log.Printf("Push Success")
}
