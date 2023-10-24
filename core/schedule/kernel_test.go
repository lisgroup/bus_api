package schedule

import (
	"fmt"
	"testing"
	"time"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
	fmt.Println(j.Name, time.Now())
}

func TestNewInstance(t *testing.T) {
	ins := NewInstance()
	err := ins.BindTask("test1", MyJob{Name: "test1"})
	if err != nil {
		t.Error(err)
	}
	err = ins.BindTask("test2", MyJob{Name: "test2"})
	if err != nil {
		t.Error(err)
	}
	for _, task := range ins.Tasks() {
		t.Log(task)
	}
	entryID, err := ins.UpdateSchedule("test1", "@every 60s")
	if err != nil {
		t.Error(err)
	}
	entryID2, err := ins.UpdateSchedule("test1", "@every 10s")
	if err != nil {
		t.Error(err)
	}
	entryID3, err := ins.UpdateSchedule("test2", "@every 20s")
	if err != nil {
		t.Error(err)
	}
	t.Logf("entryId1: %d, id2: %d id3: %d", entryID, entryID2, entryID3)
	for _, entry := range ins.cron.Entries() {
		t.Log(entry)
	}
	time.Sleep(5 * time.Minute)
}
