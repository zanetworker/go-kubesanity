package version

import (
	"io/ioutil"
	"os/exec"
	"strings"

	utils "github.com/zanetworker/go-kubesanity/pkg/kubesanityutils"
)

func getVersion() string {
	versionFileLocation := utils.GetDir("root") + "/" + "VERSION"
	versionBytes, err := ioutil.ReadFile(versionFileLocation)
	if err != nil {
		return ""
	}
	return string(versionBytes)
}

func getGitCommit() string {
	cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	outputBytes, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(outputBytes), "\n")
}
