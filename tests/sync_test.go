package tests

import (
	"log"
	"time"

	"github.com/deepglint/streamtools/st/blocks"
	"github.com/deepglint/streamtools/st/loghub"
	"github.com/deepglint/streamtools/test_utils"
	. "launchpad.net/gocheck"
)

type SyncSuite struct{}

var syncSuite = Suite(&SyncSuite{})

func (s *SyncSuite) TestSync(c *C) {
	loghub.Start()
	log.Println("testing Sync")
	b, ch := test_utils.NewBlock("testingSync", "sync")
	go blocks.BlockRoutine(b)
	time.AfterFunc(time.Duration(5)*time.Second, func() {
		ch.QuitChan <- true
	})
	err := <-ch.ErrChan
	if err != nil {
		c.Errorf(err.Error())
	}
}
