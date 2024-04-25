package layouts

import "github.com/gotk3/gotk3/gtk"

// StackTabs is a GTK layout that wraps Gtk.Paned.
// It has a sidebar and a stack with tabs.
// Sidebar can be used to change tabs.
// Sizes of the sidebar and stack can be changed dynamically.
type StackTabs struct {
	*gtk.Paned

	stack   *gtk.Stack
	sidebar *gtk.StackSidebar
}

// NewStackTabs returns a new instance of StackTabs.
func NewStackTabs(orientation gtk.Orientation) (*StackTabs, error) {
	var err error

	st := &StackTabs{}

	st.Paned, err = gtk.PanedNew(orientation)
	if err != nil {
		return nil, err
	}

	st.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}

	st.sidebar, err = gtk.StackSidebarNew()
	if err != nil {
		return nil, err
	}
	st.sidebar.SetStack(st.stack)

	st.Pack1(st.sidebar, true, true)
	st.Pack2(st.stack, true, false)

	return st, nil
}

// AddTab adds a tab to the stack.
func (st *StackTabs) AddTab(w gtk.IWidget, name, title string) {
	st.stack.AddTitled(w, name, title)
}

// SetActiveTab sets the active tab by name.
func (st *StackTabs) SetActiveTab(name string) {
	st.stack.SetVisibleChildName(name)
}
