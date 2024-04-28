package desktop

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// Editor is a GTK component based on layouts.StackTabs.
// It provides test editor capabilities.
type Editor struct {
	*layouts.StackTabs

	i18n *i18n.GtkEditorI18n
	tabs uint
}

// NewEditor returns a new instance of Editor.
func (b Builder) NewEditor(
	tests []string,
	loadTest func(testName string) (*test.Test, error),
	onSubmit func(t *test.Test) error,
) *Editor {
	editor := &Editor{
		StackTabs: Must(layouts.NewStackTabs(gtk.ORIENTATION_HORIZONTAL)),

		i18n: b.app.I18n.Gtk.Editor,
		tabs: 1,
	}

	editorMenu := b.NewEditorMenu(
		tests,
		func(testName string) {
			if testName == "" {
				return
			}

			t, err := loadTest(testName)
			if err != nil {
				return
			}

			page := Must(gtk.ScrolledWindowNew(nil, nil))
			testForm := b.NewTestForm(onSubmit, page.Destroy)

			page.Add(testForm)
			editor.AddTab(page, testName, testName)
			page.ShowAll()

			testForm.SetTest(t)
		},
		func() {
			tabName := fmt.Sprintf(editor.i18n.NewTestLabel, editor.tabs)
			editor.tabs++

			page := Must(gtk.ScrolledWindowNew(nil, nil))

			testForm := b.NewTestForm(
				onSubmit,
				page.Destroy,
			)

			page.Add(testForm)
			editor.AddTab(page, tabName, tabName)
			page.ShowAll()
		},
	)

	editor.AddTab(
		editorMenu,
		editor.i18n.Menu.SidebarLabel,
		editor.i18n.Menu.SidebarLabel,
	)

	return editor
}
