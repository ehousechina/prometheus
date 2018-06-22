package rules

import (
	"fmt"
	"testing"
	"time"

	"github.com/prometheus/common/promlog"
)

func TestLoadFile(t *testing.T) {

	cfg, _ := LoadFile(TestMgrConfigFile)

	if cfg.SyncURL != "http://10.122.138.18:8500/v1/kv/BU/jishubaozhangbu/Project/prometheus/Service/prometheus/AlertRuleTests/" {
		t.Fail()
	}

}

func TestLoadAlertRuleTests(t *testing.T) {

	level := promlog.AllowedLevel{}
	level.Set("debug")

	logger = promlog.New(level)

	initSync()
	startSync(logger)

	arts, _ := loadAlertRuleTests()

	for _, art := range arts {
		fmt.Println(art.AlertName)
	}

	if len(arts) != 1 {
		t.Fail()
	}

	time.Sleep(2 * time.Second)
}
