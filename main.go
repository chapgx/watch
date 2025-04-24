package main

import (
	"fmt"
	"github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
	"github.com/racg0092/rhombifer/pkg/models"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	if e := rhombifer.Start(); e != nil {
		panic(e)
	}

}

func init() {
	root := rhombifer.Root()
	help := builtin.HelpCommand(nil, nil)
	root.AddSub(&help)
	root.Run = RootRun
	root.AddFlags(&models.Flag{Name: "interval", Short: "watcher inteval in seconds", SingleValue: false, ShortFormat: "i"})
	config := rhombifer.GetConfig()
	config.RunHelpIfNoInput = true
}

func RootRun(args ...string) error {
	interval := 1 * time.Second
	f, _ := rhombifer.FindFlag("interval", "i")
	if f != nil {
		v, e := f.GetSingleValue()
		if e != nil {
			panic(e)
		}
		v = strings.ReplaceAll(v, "\n", "")
		i, e := strconv.Atoi(v)
		if e != nil {
			panic(e)
		}
		interval = time.Second * time.Duration(i)
	}

	params := args[2:]
	if len(params) == 0 {
		fmt.Println("expected command and args got 0")
		return nil
	}

	cmd := params[0]
	for {

		fmt.Print("\x1b[2j\x1b[H")

		var exe *exec.Cmd
		if len(params) > 1 {
			p := params[1:]
			exe = exec.Command(cmd, p...)
		} else {
			exe = exec.Command(cmd)
		}

		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr

		e := exe.Run()
		if e != nil {
			return e
		}

		exe.Process.Kill()

		time.Sleep(interval)
	}

}
