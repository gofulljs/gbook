package sync

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
}

func (s *Sync2TestSuite) SetupSuite() {
	s.bookVersionPath = filepath.Join(s.bookHome, global.BOOK_VERSION)
}

func (s *Sync2TestSuite) TestAll() {
	s.RunGitbookInstall()
}

func (s *Sync2TestSuite) RunGitbookInstall() {
	os.RemoveAll(s.bookHome)
	s.False(checkGitbookIsExist(s.bookVersionPath))
	err := MustDownloadGitbook(s.bookHome, "", "", s.sourceURI)
	s.True(checkGitbookIsExist(s.bookVersionPath))
	s.NoError(err)
}

func TestAll(t *testing.T) {
	suite.Run(t, &Sync2TestSuite{
		bookHome:  "tmp",
		sourceURI: "https://github.com/gofulljs/gitbook/releases/download/3.2.3/3.2.3.tar.gz",
	})
}
