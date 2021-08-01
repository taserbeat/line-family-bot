package functions_test

import (
	"testing"

	"github.com/taserbeat/line-family-bot/modules/functions"
)

func TestGetToday(t *testing.T) {
	_, err := functions.GetToday()
	if err != nil {
		t.Errorf("「今日は何の日」の取得エラー: %v", err)
	}
}
