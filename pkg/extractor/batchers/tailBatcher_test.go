package batchers

import (
	"rare/pkg/extractor"
	"testing"

	"github.com/hpcloud/tail"
	"github.com/stretchr/testify/assert"
)

func TestTailLineToChan(t *testing.T) {
	tailchan := make(chan *tail.Line)
	ret := make(chan extractor.InputBatch, 10)
	go tailLineToChan("test", tailchan, 1, ret)

	tailchan <- &tail.Line{
		Text: "Hello",
	}

	val := <-ret
	assert.Equal(t, "test", val.Source)
	assert.Equal(t, extractor.BString("Hello"), val.Batch[0])
	assert.Equal(t, uint64(1), val.BatchStart)

	close(tailchan)
}
