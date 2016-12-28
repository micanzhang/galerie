package galerie

import (
	"log"
	"os"
	"os/exec"
)

type Command interface {
	Run(arg ...string)
	Help()
}

type Cmd struct {
	Name    string `json:"name"`
	Alias   string `json:"alias"`
	Command string `json:"command"`
}

func (c *Cmd) Run(arg ...string) {
	cmd := exec.Command(c.Command, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
