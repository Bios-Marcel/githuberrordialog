package githuberrordialog

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

//ShowErrorDialog shows an error dialog to the user, offering him to report the issue on Github.
//The error messages will contain information about the users runtime and the error message.
func ShowErrorDialog(err error) {
	ShowErrorDialogWithMessage("An error occurred.", err)
}

//ShowErrorDialogWithMessage shows an error dialog to the user, offering him to report the issue on Github.
//The error messages will contain information about the users runtime and the error message.
func ShowErrorDialogWithMessage(message string, err error) {
	title, _ := gtk.LabelNew(message)
	title.SetHAlign(gtk.ALIGN_START)

	errorMessage := fmt.Sprintf("The following error occurred during execution: \n\t%s", err.Error())
	errorMessageLabel, _ := gtk.LabelNew(errorMessage)
	errorMessageLabel.SetHAlign(gtk.ALIGN_START)

	layout, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	layout.SetMarginTop(10)
	layout.SetMarginEnd(10)
	layout.SetMarginBottom(10)
	layout.SetMarginStart(10)
	layout.Add(title)
	layout.Add(errorMessageLabel)

	//Repository might be empty, in that case the dialog won't offer a report link.
	if Repository != "" {
		reportIssueLink, _ := gtk.LinkButtonNew("Report this issue on Github")
		reportIssueLink.SetUri(CreateIssueURL(Repository, message))
		layout.Add(reportIssueLink)
	}

	errorDialog, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	errorDialog.SetTitle("Error")
	errorDialog.Add(layout)
	errorDialog.SetResizable(false)
	errorDialog.SetPosition(gtk.WIN_POS_CENTER)

	errorDialog.ShowAll()
}

//ShowErrorDialogOnError only shows an error dialog if the passed error is not equal to nil
func ShowErrorDialogOnError(err error) {
	if err != nil {
		ShowErrorDialog(err)
	}
}
