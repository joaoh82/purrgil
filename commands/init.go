package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/guidiego/purrgil/file"
)

func Init(projectName string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	wd, _ := os.Getwd()
	wdpath := wd + projectName

	fmt.Printf("%s starting your app....\n", magenta("PURRGIL"))

	os.Mkdir(projectName, 0777)

	dockercompose := file.NewDockerCompose(wdpath)
	purrgil := file.NewPurrgil(wdpath, projectName)

	dockercompose.SaveFile()
	purrgil.SaveFile()

	fmt.Printf("%s app created! To enter in project give 'cd %s' \n", magenta("PURRGIL"), projectName)

}
