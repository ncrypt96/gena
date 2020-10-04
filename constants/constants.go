package constants

import "fmt"

const (
	RepoURLUsage = "A valid git repository url"
	DirPathUsage = "The directory to be used"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	SuccessColor = "\033[0;32m%s\033[0m"
)

const ConfFileName = "__config.json"

const DirNotEmptyErr = "The given Directory: %s is not empty"

var (
	EnterValForVarMsg = fmt.Sprintf("-> Please enter value for variable: %s\n ", InfoColor)
	DefaultValMsg     = fmt.Sprintf("Default: %s\n", InfoColor)
)

const (
	RepoFlag          = "repo"
	CloneDirFLag      = "dir"
	RepoFlagShort     = "r"
	CloneDirFLagShort = "d"
	DefaultCloneDir   = "./"
	DefaultRepo       = ""
)

const (
	LoaderType               = 21
	EmptyGitURL              = "Please enter a valid git repository url"
	PrivateURLUnsupportedErr = "Private repositories are currently not supported"
	ModFilesMsg              = "\nModifying files\n"
	HappyHackingMsg          = "\n\n(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧Happy Hacking!✧ﾟ･: *ヽ(◕ヮ◕ヽ)\n\n"
)

var (
	ClonigFrom           = fmt.Sprintf("Cloning from %s to %s \n", NoticeColor, NoticeColor)
	ClonedRepoSuccessMsg = fmt.Sprintf("%s:\n\tCloned the repository to the given directory\n\n", SuccessColor)
	GenFilesSuccessMsg   = fmt.Sprintf("\n%s:\n\tgenerated your files\n", SuccessColor)
)
