package punchclock

import (
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SetMainMenu(metadata fyne.AppMetadata, w fyne.Window) *fyne.MainMenu {
	version := metadata.Version + " build " + strconv.Itoa(metadata.Build)

	versionLabel := widget.NewLabel("Punchclock v" + version)
	versionLabel.Alignment = fyne.TextAlignCenter
	description := widget.NewLabel("Software to help the user record working time.")
	description.Alignment = fyne.TextAlignCenter
	copyright := widget.NewLabel("Copyright (C) 2023 Martin Olausson")
	copyright.Alignment = fyne.TextAlignCenter
	linkText := "https://github.com/simmarn/punchclock/"
	url, _ := url.Parse(linkText)
	link := widget.NewHyperlink(linkText, url)
	link.Alignment = fyne.TextAlignCenter

	aboutContainer := container.New(layout.NewVBoxLayout(), versionLabel, description, copyright, link)
	about := fyne.NewMenuItem("About", func() {
		dialog.ShowCustom("About", "Close", aboutContainer, w)
	})
	licenseLabel := widget.NewLabel(string(resourceLicenseTxt.Content()))
	licenseLinkText := "https://github.com/simmarn/punchclock/blob/master/LICENSE"
	liceseUrl, _ := url.Parse(licenseLinkText)
	licenseUrlLabel := widget.NewHyperlink(licenseLinkText, liceseUrl)
	licenseBox := container.New(layout.NewVBoxLayout(), licenseLabel, licenseUrlLabel)
	licenseScroll := container.NewScroll(licenseBox)
	licenseScroll.SetMinSize(licenseBox.MinSize())
	license := fyne.NewMenuItem("License", func() {
		dialog.ShowCustom("License", "OK", licenseScroll, w)
	})
	mainmenu := fyne.NewMainMenu(fyne.NewMenu("...", about, license))
	return mainmenu
}
