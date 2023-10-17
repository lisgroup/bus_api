package push

import (
	"testing"
)

func TestServerJPush(t *testing.T) {
	server := NewServerJ(ServerJParam{Key: "SCT204585TYnW8nBPYBPoYoGeWaG6kap6j",
		Title: "test", Desp: "我是一个测试"})
	err := server.Push()
	if err != nil {
		t.Error(err)
	}
}
