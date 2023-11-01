package sync2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gofulljs/gbook/global"
	"github.com/stretchr/testify/suite"
)

type Sync2TestSuite struct {
	suite.Suite
	bookHome        string
	sourceURI       string
	bookVersionPath string
	onlyStep        int
}

func (s *Sync2TestSuite) SetupSuite() {
	s.bookVersionPath = filepath.Join(s.bookHome, global.BOOK_VERSION)
}

func (s *Sync2TestSuite) TestAll() {
	if s.onlyStep == 1 || s.onlyStep == 0 {
		s.RunGitbookInstall()
	}
	if s.onlyStep == 2 || s.onlyStep == 0 {
		s.RunNodeInstall()
	}
}

func (s *Sync2TestSuite) RunGitbookInstall() {
	os.RemoveAll(s.bookHome)
	s.False(checkGitbookIsExist(s.bookVersionPath))
	err := MustDownloadGitbook(s.bookHome, "", "", s.sourceURI)
	s.NoError(err)
}

func (s *Sync2TestSuite) RunNodeInstall() {
	s.False(checkGitbookIsExist(s.bookVersionPath))
	err := nodeInstall(s.bookVersionPath)
	s.NoError(err, "%+v", err)
	s.True(checkGitbookIsExist(s.bookVersionPath))
}

func TestAll(t *testing.T) {
	suite.Run(t, &Sync2TestSuite{
		bookHome:  "tmp",
		sourceURI: "https://github.com/gofulljs/gitbook/archive/refs/tags/3.2.3.tar.gz",
		onlyStep:  0,
	})
}
