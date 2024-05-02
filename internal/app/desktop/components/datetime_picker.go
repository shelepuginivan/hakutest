package components

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
)

// DatetimePicker is a GTK component based on Gtk.Calendar that allows to pick
// both date and time.
type DatetimePicker struct {
	*gtk.Box

	hourSpin   *gtk.SpinButton
	minuteSpin *gtk.SpinButton
	secondSpin *gtk.SpinButton
	calendar   *gtk.Calendar
}

// NewDatetimePicker returns a new instance of DatetimePicker.
func NewDatetimePicker() (*DatetimePicker, error) {
	var err error

	dp := DatetimePicker{}

	dp.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	dp.hourSpin, err = gtk.SpinButtonNewWithRange(0, 23, 1)
	if err != nil {
		return nil, err
	}

	dp.minuteSpin, err = gtk.SpinButtonNewWithRange(0, 59, 1)
	if err != nil {
		return nil, err
	}

	dp.secondSpin, err = gtk.SpinButtonNewWithRange(0, 59, 1)
	if err != nil {
		return nil, err
	}

	timeBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return nil, err
	}

	delimHM, err := gtk.LabelNew(":")
	if err != nil {
		return nil, err
	}

	delimMS, err := gtk.LabelNew(":")
	if err != nil {
		return nil, err
	}

	dp.calendar, err = gtk.CalendarNew()
	if err != nil {
		return nil, err
	}

	timeBox.PackStart(dp.hourSpin, false, false, 2)
	timeBox.PackStart(delimHM, false, false, 2)
	timeBox.PackStart(dp.minuteSpin, false, false, 2)
	timeBox.PackStart(delimMS, false, false, 2)
	timeBox.PackStart(dp.secondSpin, false, false, 2)

	dp.PackStart(timeBox, false, false, 4)
	dp.PackStart(dp.calendar, false, false, 4)

	return &dp, nil
}

// GetDate returns a selected datetime.
func (dp DatetimePicker) GetDate() time.Time {
	year, month, day := dp.calendar.GetDate()

	return time.Date(
		int(year),
		time.Month(month+1), // GTK month is zero-indexed.
		int(day),
		dp.hourSpin.GetValueAsInt(),
		dp.minuteSpin.GetValueAsInt(),
		dp.secondSpin.GetValueAsInt(),
		0,
		time.Local,
	)
}

// SetDate sets the selected datetime.
func (dp *DatetimePicker) SetDate(date time.Time) {
	dp.calendar.SelectDay(uint(date.Day()))
	dp.calendar.SelectMonth(uint(date.Month()-1), uint(date.Year()))
	dp.hourSpin.SetValue(float64(date.Hour()))
	dp.minuteSpin.SetValue(float64(date.Minute()))
	dp.secondSpin.SetValue(float64(date.Second()))
}
