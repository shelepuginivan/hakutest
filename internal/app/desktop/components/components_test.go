package components

import (
	"container/list"
	"fmt"
	"math/rand"
	"testing"
	"time"

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

func TestAttachmentSource(t *testing.T) {
	as, err := NewAttachmentSource("", "", "", "", "", "")

	assert.NoError(t, err)
	assert.IsType(t, &AttachmentSource{}, as)
	assert.IsType(t, &gtk.Stack{}, as.stack)
	assert.IsType(t, &gtk.Label{}, as.loadedLabel)
	assert.IsType(t, &gtk.Entry{}, as.urlEntry)
	assert.IsType(t, &gtk.FileChooserButton{}, as.fileButton)
	assert.IsType(t, &gtk.RadioButton{}, as.baseRadio)
	assert.Equal(t, AttachmentSourceModeURL, as.mode)

	as.SetMode(AttachmentSourceModeFile)
	assert.Equal(t, AttachmentSourceModeFile, as.mode)

	as.SetMode(AttachmentSourceModeLoaded)
	assert.Equal(t, AttachmentSourceModeLoaded, as.mode)

	urlSrc := "https://example.com"
	as.SetSource(urlSrc)
	assert.Equal(t, AttachmentSourceModeURL, as.mode)

	src, err := as.urlEntry.GetText()
	assert.NoError(t, err)
	assert.Equal(t, urlSrc, src)

	src, err = as.GetSource()
	assert.NoError(t, err)
	assert.Equal(t, urlSrc, src)

	base64Src := "data:image/png;base64,somebase64stuff"
	as.SetSource(base64Src)
	assert.Equal(t, AttachmentSourceModeLoaded, as.mode)

	assert.Equal(t, base64Src, as.loadedSource)

	src, err = as.GetSource()
	assert.NoError(t, err)
	assert.Equal(t, base64Src, src)

	window.Add(as)
	as.Destroy()
}

func TestComponentList(t *testing.T) {
	cl, err := NewComponentList("", "-", func() (*gtk.Label, error) {
		return gtk.LabelNew("label")
	})

	assert.NoError(t, err)
	assert.IsType(t, &ComponentList[*gtk.Label]{}, cl)
	assert.IsType(t, &list.List{}, cl.list)
	assert.IsType(t, &gtk.Box{}, cl.componentBox)
	assert.Equal(t, "-", cl.removeButtonLabel)

	for i := range 100 {
		l, err := gtk.LabelNew(fmt.Sprintf("Label %d", i+1))
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, cl.AddComponent(l))
	}

	cnt := 0

	cl.ForEach(func(w *gtk.Label) {
		assert.IsType(t, &gtk.Label{}, w)
		cnt += 1
	})

	assert.Equal(t, 100, cnt)

	window.Add(cl)
	cl.Clear()
	cl.Destroy()
}

func TestDatetimePicker(t *testing.T) {
	dp, err := NewDatetimePicker()

	assert.NoError(t, err)
	assert.IsType(t, &DatetimePicker{}, dp)

	datetime := time.Now()

	dp.SetDate(datetime)
	actual := dp.GetDate()

	assert.Equal(t, datetime.Year(), actual.Year())
	assert.Equal(t, datetime.Month(), actual.Month())
	assert.Equal(t, datetime.Day(), actual.Day())
	assert.Equal(t, datetime.Hour(), actual.Hour())
	assert.Equal(t, datetime.Minute(), actual.Minute())
	assert.Equal(t, datetime.Second(), actual.Second())
	assert.Equal(t, 0, actual.Nanosecond())

	window.Add(dp)
	dp.Destroy()
}

func TestEntryListMultiple(t *testing.T) {
	lst, err := NewEntryListMultiple("", "", "-")

	assert.NoError(t, err)
	assert.IsType(t, &EntryListMultiple{}, lst)
	assert.IsType(t, &ComponentList[*EntryListMultipleItem]{}, lst.cl)
	assert.Equal(t, "-", lst.cl.removeButtonLabel)

	item, err := lst.NewEntryListMultipleItem(&EntryListMultipleValue{
		Text:     "text",
		Selected: false,
	})

	assert.NoError(t, err)
	assert.IsType(t, &EntryListMultipleItem{}, item)

	text, err := item.entry.GetText()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "text", text)
	assert.False(t, item.check.GetActive())

	var values []*EntryListMultipleValue

	for i := range 100 {
		values = append(values, &EntryListMultipleValue{
			Text:     fmt.Sprintf("Entry %d", i+1),
			Selected: rand.Intn(2) == 1,
		})
	}

	lst.SetValues(values)
	assert.EqualValues(t, values, lst.GetValues())

	window.Add(lst)
	lst.Destroy()
}

func TestEntryListSingle(t *testing.T) {
	lst, err := NewEntryListSingle("", "", "-")

	assert.NoError(t, err)
	assert.IsType(t, &EntryListSingle{}, lst)
	assert.IsType(t, &ComponentList[*EntryListSingleItem]{}, lst.cl)
	assert.Equal(t, "-", lst.cl.removeButtonLabel)

	item, err := lst.NewEntryListSingleItem(&EntryListSingleValue{
		Text:     "text",
		Selected: false,
	})

	assert.NoError(t, err)
	assert.IsType(t, &EntryListSingleItem{}, item)

	text, err := item.entry.GetText()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "text", text)
	assert.False(t, item.radio.GetActive())

	var values []*EntryListSingleValue

	for i := range 100 {
		values = append(values, &EntryListSingleValue{
			Text:     fmt.Sprintf("Entry %d", i+1),
			Selected: false,
		})
	}

	values = append(values, &EntryListSingleValue{
		Text:     "Entry 101",
		Selected: true,
	})

	lst.SetValues(values)
	assert.EqualValues(t, values, lst.GetValues())

	window.Add(lst)
	lst.Destroy()
}

func TestHeadings(t *testing.T) {
	h1, err := NewHeadingH1("heading h1")
	assert.NoError(t, err)
	assert.IsType(t, &gtk.Label{}, h1)
	assert.True(t, h1.GetUseMarkup())

	h2, err := NewHeadingH2("heading h2")
	assert.NoError(t, err)
	assert.IsType(t, &gtk.Label{}, h2)
	assert.True(t, h2.GetUseMarkup())

	window.Add(h1)
	h1.Destroy()

	window.Add(h2)
	h2.Destroy()
}

func TestInput(t *testing.T) {
	e, err := gtk.EntryNew()
	if err != nil {
		t.Fatal(err)
	}

	i, err := NewInput("label", e)

	assert.NoError(t, err)
	assert.IsType(t, &Input{}, i)
	assert.Equal(t, gtk.ORIENTATION_VERTICAL, i.GetOrientation())
	assert.Equal(t, 4, i.GetSpacing())
	assert.Equal(t, gtk.ALIGN_FILL, i.GetHAlign())

	window.Add(i)
	i.Destroy()
}
