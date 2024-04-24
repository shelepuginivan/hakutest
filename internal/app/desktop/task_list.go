package desktop

import (
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// TaskList is a GTK component based on components.ComponentList.
// It allows to edit multiple test tasks.
type TaskList struct {
	*components.ComponentList[*TaskInput]

	b *Builder
}

// NewTaskList returns a new instance of TaskList.
func (b Builder) NewTaskList() *TaskList {
	tl := &TaskList{
		ComponentList: Must(components.NewComponentList("Add task", "-", func() (*TaskInput, error) {
			return b.NewTaskInput(), nil
		})),
		b: &b,
	}

	return tl
}

// GetTasks returns tasks.
func (tl TaskList) GetTasks() []*test.Task {
	var tasks []*test.Task

	tl.ForEach(func(w *TaskInput) {
		task, err := w.GetTask()
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	})

	return tasks
}

// SetTasks sets tasks.
func (tl *TaskList) SetTasks(tasks []*test.Task) {
	for _, task := range tasks {
		ti := tl.b.NewTaskInput()
		ti.SetTask(task)
		tl.AddComponent(ti)
	}
}
