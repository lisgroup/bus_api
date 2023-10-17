package schedule

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"sort"
	"sync"
)

type Task struct {
	Name string       `json:"task_name"`
	Id   cron.EntryID `json:"id"`
	Job  cron.Job     `json:"job"`
}

type Instance struct {
	tasks map[string]*Task
	cron  *cron.Cron
	mux   sync.Mutex
}

var instance *Instance

// NewInstance 计划任务初始化
func NewInstance() *Instance {
	if instance != nil {
		return instance
	}
	instance = &Instance{}
	var opts []cron.Option
	opts = append(opts, cron.WithChain(cron.Recover(cron.DefaultLogger)))
	instance.cron = cron.New(opts...)
	instance.tasks = make(map[string]*Task)
	instance.cron.Start()
	return instance
}

// BindTask 绑定任务
func (i *Instance) BindTask(name string, job cron.Job) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	if _, ok := i.tasks[name]; ok {
		return errors.New(fmt.Sprintf("job %s exists", name))
	}
	i.tasks[name] = &Task{
		Name: name,
		Job:  job,
	}
	return nil
}

// RemoveTask 删除任务
func (i *Instance) RemoveTask(name string) {
	i.mux.Lock()
	defer i.mux.Unlock()
	if job, ok := i.tasks[name]; ok {
		if job.Id != 0 {
			i.cron.Remove(job.Id)
		}
		delete(i.tasks, name)
	}
}

// BindSchedule 绑定执行计划
func (i *Instance) BindSchedule(name, spec string) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	if job, ok := i.tasks[name]; ok {
		var err error
		job.Id, err = i.cron.AddJob(spec, job.Job)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("job %s not exists", name))
	}
}

// BindTaskAndSchedule 绑定任务和执行计划
func (i *Instance) BindTaskAndSchedule(name, spec string, job cron.Job) error {
	err := i.BindTask(name, job)
	if err != nil {
		return err
	}
	return i.BindSchedule(name, spec)
}

// UpdateSchedule 更新执行计划
func (i *Instance) UpdateSchedule(name, spec string) (cron.EntryID, error) {
	i.mux.Lock()
	defer i.mux.Unlock()
	if job, ok := i.tasks[name]; ok {
		if job.Id != 0 {
			i.cron.Remove(job.Id)
		}
		var err error
		job.Id, err = i.cron.AddJob(spec, job.Job)
		return job.Id, err
	} else {
		return job.Id, errors.New(fmt.Sprintf("job %s not exists", name))
	}
}

// Tasks 任务列表
func (i *Instance) Tasks() map[string]*Task {
	i.mux.Lock()
	defer i.mux.Unlock()
	return instance.tasks
}

// GetTask 任务列表
func (i *Instance) GetTask(name string) *Task {
	i.mux.Lock()
	defer i.mux.Unlock()
	if job, ok := i.tasks[name]; ok {
		return job
	}
	return nil
}

// GetEntries 获取全部的 cron.Entry
func (i *Instance) GetEntries() []cron.Entry {
	i.mux.Lock()
	defer i.mux.Unlock()
	entries := i.cron.Entries()
	// 排序
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ID < entries[j].ID
	})
	return entries
}

// GetEntry 根据 cron.EntryID 获取对应的 cron.Entry
func (i *Instance) GetEntry(entries []cron.Entry, id cron.EntryID) cron.Entry {
	i.mux.Lock()
	defer i.mux.Unlock()
	// 二分查找
	idx := sort.Search(len(entries), func(i int) bool {
		return entries[i].ID >= id
	})
	if idx < len(entries) && entries[idx].ID == id {
		return entries[idx]
	}
	return cron.Entry{}
}
