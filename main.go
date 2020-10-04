package main

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/ncrypt96/gena/conf"
	"github.com/ncrypt96/gena/constants"
	"github.com/ncrypt96/gena/fs"
	"github.com/ncrypt96/gena/githandle"
	"github.com/ncrypt96/gena/helpers"
	"github.com/ncrypt96/gena/templating"
	"github.com/ncrypt96/gena/utils"
	flag "github.com/spf13/pflag"
)

func main() {
	repoURL := flag.StringP(constants.RepoFlag, constants.RepoFlagShort, constants.DefaultRepo, constants.RepoURLUsage)
	cloneDirectory := flag.StringP(constants.CloneDirFLag, constants.CloneDirFLagShort, constants.DefaultCloneDir, constants.DirPathUsage)
	flag.Parse()

	if !utils.IsNotEmptyStr(*repoURL) {
		fmt.Printf(constants.EmptyGitURL)
		os.Exit(1)
	}

	fmt.Printf(constants.ClonigFrom, *repoURL, *cloneDirectory)

	s := spinner.New(spinner.CharSets[constants.LoaderType], 100*time.Millisecond)
	s.Start()

	url := *repoURL
	directory := *cloneDirectory

	dExists, dErr := fs.DirExists(directory)
	if dErr != nil {
		fmt.Printf(constants.ErrorColor, dErr)
		os.Exit(1)
	}

	if dExists {
		_, deErr := fs.DirIsEmpty(directory)
		if deErr != nil {
			fmt.Printf(constants.ErrorColor, deErr)
			os.Exit(1)
		}
	}

	cDir := fs.CreateDirIfNotExists(directory)
	if cDir != nil {
		fmt.Printf(constants.ErrorColor, cDir)
		os.Exit(1)
	}

	err := githandle.PlainClone(directory, url)
	s.Stop()

	if err != nil {
		if err == transport.ErrAuthenticationRequired {
			fmt.Printf(constants.ErrorColor, constants.PrivateURLUnsupportedErr)
		} else {
			fmt.Printf(constants.ErrorColor, err)
		}
		os.Exit(1)
	}

	fmt.Printf(constants.ClonedRepoSuccessMsg, "SUCCESS")

	cf := conf.NewConfig()
	by, err := cf.GetConfigAsBytes(directory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cf.UnmarshalConfig(by)
	helpers.ChangeDefaultVarMap(cf.Variables)
	fmt.Printf(constants.NoticeColor, constants.ModFilesMsg)
	s.Start()
	paths, err := fs.GetPaths(directory, cf.AddGit)

	templating.GenTemplateFiles(paths, cf.Variables)
	s.Stop()
	fmt.Printf(constants.GenFilesSuccessMsg, "SUCCESS")
	fmt.Printf(constants.SuccessColor, constants.HappyHackingMsg)

}
