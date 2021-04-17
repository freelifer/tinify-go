package tinify

import (
	"testing"
	"time"
)

const Key = "rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"

func TestCompressFromFile(t *testing.T) {
	SetKey(Key)
	t.Log("FromFile start...", time.Now() )
	source, err := FromFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("FromFile end...", time.Now() )
	err = source.ToFile("./test_output/CompressFromFile.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Compress successful", time.Now() )
}
