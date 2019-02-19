package resource_test

import (
	"github.com/binsix/transfer-data/src/main/utils"
	"testing"
)

func TestA(t *testing.T) {
	t.Log("A")
	t.Log(utils.GetConf())
}
