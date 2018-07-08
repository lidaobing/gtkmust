package gtkmust

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func AboutDialogNew() *gtk.AboutDialog {
	res, err := gtk.AboutDialogNew()
	if err != nil {
		panic(err)
	}
	return res
}

func ApplicationNew(appId string, flags glib.ApplicationFlags) *gtk.Application {
	res, err := gtk.ApplicationNew(appId, flags)
	if err != nil {
		panic(err)
	}
	return res
}

func ApplicationWindowNew(app *gtk.Application) *gtk.ApplicationWindow {
	res, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		panic(err)
	}
	return res
}

func BoxNew(o gtk.Orientation, spacing int) *gtk.Box {
	res, err := gtk.BoxNew(o, spacing)
	if err != nil {
		panic(err)
	}
	return res
}

func ButtonNew() *gtk.Button {
	res, err := gtk.ButtonNew()
	if err != nil {
		panic(err)
	}
	return res
}

func ButtonNewWithLabel(label string) *gtk.Button {
	res, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		panic(err)
	}
	return res
}

func CellRendererPixbufNew() *gtk.CellRendererPixbuf {
	res, err := gtk.CellRendererPixbufNew()
	if err != nil {
		panic(err)
	}
	return res
}

func CellRendererTextNew() *gtk.CellRendererText {
	res, err := gtk.CellRendererTextNew()
	if err != nil {
		panic(err)
	}
	return res
}

func DialogNew() *gtk.Dialog {
	res, err := gtk.DialogNew()
	if err != nil {
		panic(err)
	}
	return res
}

func FileChooserDialogNewWith1Button(title string,
	parent *gtk.Window,
	action gtk.FileChooserAction,
	first_button_text string,
	first_button_id gtk.ResponseType) *gtk.FileChooserDialog {
	res, err := gtk.FileChooserDialogNewWith1Button(title, parent, action, first_button_text, first_button_id)
	if err != nil {
		panic(err)
	}
	return res
}

func FileChooserDialogNewWith2Buttons(title string,
	parent *gtk.Window,
	action gtk.FileChooserAction,
	first_button_text string,
	first_button_id gtk.ResponseType,
	second_button_text string,
	second_button_id gtk.ResponseType) *gtk.FileChooserDialog {
	res, err := gtk.FileChooserDialogNewWith2Buttons(title, parent, action,
		first_button_text, first_button_id,
		second_button_text, second_button_id)
	if err != nil {
		panic(err)
	}
	return res
}

func ImageNew() *gtk.Image {
	res, err := gtk.ImageNew()
	if err != nil {
		panic(err)
	}
	return res
}

func LabelNew(label string) *gtk.Label {
	res, err := gtk.LabelNew(label)
	if err != nil {
		panic(err)
	}
	return res
}

func ListStoreNew(types ...glib.Type) *gtk.ListStore {
	res, err := gtk.ListStoreNew(types...)
	if err != nil {
		panic(err)
	}
	return res
}

func MenuNew() *gtk.Menu {
	res, err := gtk.MenuNew()
	if err != nil {
		panic(err)
	}
	return res
}

func MenuBarNew() *gtk.MenuBar {
	res, err := gtk.MenuBarNew()
	if err != nil {
		panic(err)
	}
	return res
}

func MenuItemNew() *gtk.MenuItem {
	res, err := gtk.MenuItemNew()
	if err != nil {
		panic(err)
	}
	return res
}

func MenuItemNewWithMnemonic(label string) *gtk.MenuItem {
	res, err := gtk.MenuItemNewWithMnemonic(label)
	if err != nil {
		panic(err)
	}
	return res
}

func ScrolledWindowNew(h *gtk.Adjustment, v *gtk.Adjustment) *gtk.ScrolledWindow {
	res, err := gtk.ScrolledWindowNew(h, v)
	if err != nil {
		panic(err)
	}
	return res
}

func TargetEntryNew(target string, flags gtk.TargetFlags, info uint) *gtk.TargetEntry {
	res, err := gtk.TargetEntryNew(target, flags, info)
	if err != nil {
		panic(err)
	}
	return res
}

func TreeViewNew() *gtk.TreeView {
	res, err := gtk.TreeViewNew()
	if err != nil {
		panic(err)
	}
	return res
}

func TreeViewColumnNew() *gtk.TreeViewColumn {
	res, err := gtk.TreeViewColumnNew()
	if err != nil {
		panic(err)
	}
	return res
}

func TreeViewColumnNewWithAttribute(title string, renderer gtk.ICellRenderer, attribute string, column int) *gtk.TreeViewColumn {
	res, err := gtk.TreeViewColumnNewWithAttribute(title, renderer, attribute, column)
	if err != nil {
		panic(err)
	}
	return res
}

func WindowNew(t gtk.WindowType) *gtk.Window {
	res, err := gtk.WindowNew(t)
	if err != nil {
		panic(err)
	}
	return res
}

func MustSList(res *glib.SList, err error) *glib.SList {
	if err != nil {
		panic(err)
	}
	return res
}

func MustBox(res *gtk.Box, err error) *gtk.Box {
	if err != nil {
		panic(err)
	}
	return res
}

func MustTreeSelection(res *gtk.TreeSelection, err error) *gtk.TreeSelection {
	if err != nil {
		panic(err)
	}
	return res
}