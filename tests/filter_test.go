package tests

import (
	"log"
	"time"

	"github.com/deepglint/streamtools/st/blocks"
	"github.com/deepglint/streamtools/test_utils"
	. "launchpad.net/gocheck"
)

type FilterSuite struct{}

var filterSuite = Suite(&FilterSuite{})

func (s *FilterSuite) TestFilter(c *C) {
	log.Println("testing Filter")
	b, ch := test_utils.NewBlock("testingFilter", "filter")
	go blocks.BlockRoutine(b)

	ruleMsg := map[string]interface{}{"Filter": ".device == 'iPhone'"}
	toRule := &blocks.Msg{Msg: ruleMsg, Route: "rule"}
	ch.InChan <- toRule

	outChan := make(chan *blocks.Msg)
	ch.AddChan <- &blocks.AddChanMsg{Route: "1", Channel: outChan}

	queryOutChan := make(blocks.MsgChan)
	ch.QueryChan <- &blocks.QueryMsg{MsgChan: queryOutChan, Route: "rule"}

	time.AfterFunc(time.Duration(5)*time.Second, func() {
		ch.QuitChan <- true
	})

	for {
		select {
		case messageI := <-queryOutChan:
			c.Assert(messageI, DeepEquals, ruleMsg)
		case message := <-outChan:
			log.Println(message)

		case err := <-ch.ErrChan:
			if err != nil {
				c.Errorf(err.Error())
			} else {
				return
			}
		}
	}
}
