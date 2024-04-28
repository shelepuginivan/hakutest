package desktop

import (
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// TaskList is a GTK component based on components.ComponentList.
// It allows to edit multiple test tasks.
type TaskList struct {
	*components.ComponentList[*TaskInput]

	i18n *i18n.GtkEditorTaskListI18n
	b    *Builder
}

// NewTaskList returns a new instance of TaskList.
func (b Builder) NewTaskList() *TaskList {
	tl := &TaskList{
		i18n: b.app.I18n.Gtk.Editor.Task.List,
		b:    &b,
	}

	tl.ComponentList = Must(components.NewComponentList(tl.i18n.ButtonAdd, "-", func() (*TaskInput, error) {
		return b.NewTaskInput(), nil
	}))

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
