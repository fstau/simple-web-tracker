package util_test

import (
	"local/tracker/util"
	"testing"
)

func TestGetTimeUnixMicro(t *testing.T) {
	res := util.GetTimeUnixMicro()
	t.Log(res)
}
