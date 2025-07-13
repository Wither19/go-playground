package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// func pkgInstall() {
// 	log.Fatalln("That feature has not been implemented yet")
// }

func getPkgList() *gtk.ScrolledWindow {
	pkgSlice := sliceFlatpakList(flatpakList("name"))
	pkgIDSlice := sliceFlatpakList(flatpakList("application"))

	pkgListScrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatalln("Scrolled Window creation error:", err)
	}

	pkgListScrolledWindow.SetMinContentHeight(160)

	pkgListBox, err := gtk.FlowBoxNew()
	if err != nil {
		log.Fatalln("List box creation error:", err)
	}
	pkgListBox.SetMarginStart(16)
	pkgListBox.SetMarginEnd(16)


	for i, pkg := range pkgSlice {
		pkgID := pkgIDSlice[i]

		pkgBtn, err := gtk.ButtonNewWithLabel(pkg)
		if err != nil {
			log.Fatalln("List entry creation error:", err)
		}

		if (pkg != "" && pkgID != "") {
			pkgBtn.Connect("clicked", func() {
				pkgRemoveWarn(pkg, pkgID)
			})

			pkgListBox.Add(pkgBtn)
		}

	}

	pkgListScrolledWindow.Add(pkgListBox)

	return pkgListScrolledWindow
}

func pkgRemoveWarn(pkgName string, pkgID string) *gtk.Dialog {
	pkgRemoveDialog, err := gtk.DialogNew()
	if err != nil {
		log.Fatalln("Dialog creation error:", err)
	}

	pkgRemoveDialog.SetTitle("Remove Package")

	appIconImg, err := gtk.ImageNewFromIconName(pkgID, gtk.ICON_SIZE_DIALOG)
	if err != nil {
		log.Fatalln("Image creation error:", err)
	}


	dialogLabelText := fmt.Sprintf("Are you sure you would like to remove %v? (%v)", pkgName, pkgID)

	dialogLabel, err := gtk.LabelNew(dialogLabelText)
	if err != nil {
		log.Fatalln("Dialog label creation error:", err)
	}

	dialogLabel.SetJustify(gtk.JUSTIFY_CENTER)

	dialogLabel.SetMarginTop(12)

	dialogLabel.SetMarginStart(24)
	dialogLabel.SetMarginEnd(24)

	dialogContents, err := pkgRemoveDialog.GetContentArea()
	if err != nil {
		log.Fatalln("Dialog contents fetch error:", err)
	}

	dialogContents.Add(appIconImg)
	dialogContents.Add(dialogLabel)

	pkgRemoveDialog.AddButton("No", gtk.RESPONSE_NO)
	pkgRemoveDialog.AddButton("Yes", gtk.RESPONSE_YES)

	pkgRemoveDialog.SetDefaultSize(400, 200)

	pkgRemoveDialog.ShowAll()

	pkgRemoveDialog.Run()
	pkgRemoveDialog.Connect("response", func(_ *gtk.Dialog, responseID gtk.ResponseType) {
		if responseID == gtk.RESPONSE_YES {
			pkgRemove(pkgID)
		}
	})

	pkgRemoveDialog.Destroy()

	return pkgRemoveDialog
}
