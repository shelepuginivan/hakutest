package layouts

import (
	"testing"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/stretchr/testify/assert"
)

var (
	application *gtk.Application
	window      *gtk.Window
)

func setup() func() {
	gtk.Init(nil)

	var err error

	application, err = gtk.ApplicationNew("com.test", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		panic(err)
	}

	window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(err)
	}

	return func() {
		window.Destroy()
		application.Quit()
		gtk.MainQuit()
	}

}

func TestMain(m *testing.M) {
	defer setup()()
	m.Run()
}

func TestContainer(t *testing.T) {
	c, err := NewContainer()

	assert.NoError(t, err)
	assert.IsType(t, c, &gtk.Box{})
	assert.Equal(t, gtk.ORIENTATION_VERTICAL, c.GetOrientation())
	assert.Equal(t, 28, c.GetSpacing())
	assert.Equal(t, gtk.ALIGN_FILL, c.GetHAlign())
	assert.Equal(t, gtk.ALIGN_CENTER, c.GetVAlign())

	window.Add(c)
	c.Destroy()
}

func TestForm(t *testing.T) {
	f, err := NewForm()

	assert.NoError(t, err)
	assert.IsType(t, f, &gtk.Box{})
	assert.Equal(t, gtk.ORIENTATION_VERTICAL, f.GetOrientation())
	assert.Equal(t, 80, f.GetSpacing())
	assert.Equal(t, gtk.ALIGN_FILL, f.GetHAlign())
	assert.Equal(t, gtk.ALIGN_CENTER, f.GetVAlign())
	assert.Equal(t, 80, f.GetMarginTop())
	assert.Equal(t, 80, f.GetMarginBottom())
	assert.Equal(t, 24, f.GetMarginStart())
	assert.Equal(t, 24, f.GetMarginEnd())

	window.Add(f)
	f.Destroy()
}

func TestNotebook(t *testing.T) {
	label, err := gtk.LabelNew("label")
	if err != nil {
		t.Fatal(err)
	}

	child, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		t.Fatal(err)
	}

	nb, err := NewNotebook(
		NotebookPage{
			Label: label,
			Child: child,
		},
	)

	assert.NoError(t, err)
	assert.IsType(t, &gtk.Notebook{}, nb)

	window.Add(nb)
	nb.Destroy()
}

func TestScrolled(t *testing.T) {
	w, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewScrolled(w)

	assert.NoError(t, err)
	assert.IsType(t, &gtk.ScrolledWindow{}, s)

	window.Add(s)
	s.Destroy()
}

func TestStackTabs(t *testing.T) {
	st, err := NewStackTabs(gtk.ORIENTATION_HORIZONTAL)

	assert.NoError(t, err)
	assert.IsType(t, &StackTabs{}, st)
	assert.IsType(t, &gtk.Stack{}, st.stack)
	assert.IsType(t, &gtk.StackSidebar{}, st.sidebar)

	window.Add(st)

	w1, err := gtk.LabelNew("1")
	if err != nil {
		t.Fatal(err)
	}

	w2, err := gtk.LabelNew("2")
	if err != nil {
		t.Fatal(err)
	}

	st.AddTab(w1, "1", "1")
	st.AddTab(w2, "2", "2")
	st.SetActiveTab("2")

	st.Destroy()
}
