package desktop

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

func (b Builder) NewEditor(
	tests []string,
	loadTest func(testName string) (*test.Test, error),
	onSubmit func(t *test.Test) error,
) *layouts.StackTabs {
	tabCounter := 1
	editor := Must(layouts.NewStackTabs(gtk.ORIENTATION_HORIZONTAL))

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
			tabName := fmt.Sprintf("New %d", tabCounter)
			tabCounter++

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

	editor.AddTab(editorMenu, "Menu", "Menu")

	return editor
}
