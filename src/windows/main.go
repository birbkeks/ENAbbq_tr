package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Türkçeleştirilmiş ENA: Dream BBQ")
	w.Resize(fyne.NewSize(800, 600))
	a.Settings().SetTheme(theme.DarkTheme())

	mainLabel := widget.NewLabel("Lütfen kaynaklar yüklenirken bekleyin...")
	init := container.New(layout.NewCenterLayout(), mainLabel)
	w.SetContent(init)
	w.Show()

	installFunc := func() {
		err := install()
		if err != 0 {
			w.SetContent(pageERR(w, err))
		} else {
			w.SetContent(page0(w))
		}
	}

	go installFunc()
	fyne.CurrentApp().Run()
}

func page0(w fyne.Window) *fyne.Container {
	mainLabel := widget.NewLabel("Türkçeleştirilmiş ENA: Dream BBQ'ya hoşgeldiniz")

	teamLabel := canvas.NewText("  BARBEQUE TEAM katkılarıyla", color.RGBA{169, 169, 169, 255})
	teamLabel.TextSize = 12

	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButton("İlerle", func() {
		w.SetContent(pageInstall(w))
	})

	page0 := container.New(layout.NewVBoxLayout(),
		mainLabel,
		teamLabel,
		btnContinue,
		errorLabel,
	)

	page0 = container.New(layout.NewCenterLayout(), page0)

	checkIntegrity(btnContinue, errorLabel)

	return page0
}

func pageERR(w fyne.Window, err int) *fyne.Container {
	pageERRContainer := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			canvas.NewText("[FATL]: Dosyalar yüklenirken kritik hata meydana geldi", color.RGBA{255, 0, 0, 255}),
			canvas.NewText("[FATL]: Hata "+fmt.Sprint(err), color.RGBA{255, 0, 0, 255}),
			widget.NewButtonWithIcon("Kapat", theme.WindowCloseIcon(), func() {
				fyne.CurrentApp().Quit()
			}),
		),
	)
	return pageERRContainer
}

func pageInstall(w fyne.Window) *fyne.Container {
	var path string
	appDir, _ := os.Getwd()
	steamIcon, _ := fyne.LoadResourceFromPath(filepath.Join(appDir, "resources", "steamIcon.png"))

	labelPath := widget.NewLabel("")
	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButtonWithIcon("Kur", theme.DownloadIcon(), func() {
		w.SetContent(pageEnd(path))
	})
	btnContinue.Disable()

	buttonsContainer := container.New(layout.NewVBoxLayout())

	btnSteam := widget.NewButtonWithIcon("Steam", steamIcon, func() {
		selectedPath := filepath.Join("C:\\", "Program Files (x86)", "Steam", "steamapps", "common", "ENA Dream BBQ")
		path = selectedPath
		checkExecutable(path, btnContinue, errorLabel)
		labelPath.SetText("Seçilen dizin: " + path)
	})

	btnBrowse := widget.NewButtonWithIcon("Aç", theme.SearchIcon(), func() {
		browseFile(w, func(selectedPath string) {
			path = selectedPath
			checkExecutable(path, btnContinue, errorLabel)
			labelPath.SetText("Seçilen dizin: " + path)
		})
	})

	buttonsContainer.Add(btnSteam)
	buttonsContainer.Add(btnBrowse)

	pageInstall := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Oyunun olduğu dizini seçin. Eğer standart dizinde kurulu ise Steam tuşuna basın."),
		buttonsContainer,
		labelPath,
		errorLabel,
		btnContinue,
	)

	pageInstall = container.New(layout.NewCenterLayout(), pageInstall)

	return pageInstall
}

func checkIntegrity(btnContinue *widget.Button, errorLabel *canvas.Text) {
	appDir, _ := os.Getwd()
	resourcesPath := filepath.Join(appDir, "resources", "meta.json")
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		btnContinue.Disable()
		errorLabel.Text = "[FATL]: \"resources\" bulunamadı."
		errorLabel.Refresh()
	} else {
		btnContinue.Enable()
		errorLabel.Text = ""
		errorLabel.Refresh()
	}
}

func checkExecutable(selectedPath string, btnContinue *widget.Button, errorLabel *canvas.Text) {
	executablePath := filepath.Join(selectedPath, "ENA-4-DreamBBQ.exe")
	if _, err := os.Stat(executablePath); os.IsNotExist(err) {
		btnContinue.Disable()
		errorLabel.Text = "[ERROR]: \"ENA-4-DreamBBQ.exe\" bulunamadı, oyun çalıştırıcısının olduğu dizini seçin"
		errorLabel.Refresh()
	} else {
		btnContinue.Enable()
		errorLabel.Text = ""
		errorLabel.Refresh()
	}
}

func browseFile(w fyne.Window, onPathSelected func(string)) {
	dialog.ShowFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if folder != nil {
			onPathSelected(folder.Path())
		}
	}, w)
}

func pageEnd(path string) *fyne.Container {
	appDir, _ := os.Getwd()
	err := inject(path)
	os.RemoveAll(filepath.Join(appDir, "resources"))
	if err != nil {
		pageEndContainer := container.New(layout.NewCenterLayout(),
			container.New(layout.NewVBoxLayout(),
				canvas.NewText("[FATL]: Dosyalar yerleştirilirken kritik hata meydana geldi.", color.RGBA{255, 0, 0, 255}),
				canvas.NewText("[FATL]: Hata"+fmt.Sprint(err), color.RGBA{255, 0, 0, 255}),
				widget.NewButtonWithIcon("Kapat", theme.WindowCloseIcon(), func() {
					fyne.CurrentApp().Quit()
				}),
			),
		)
		return pageEndContainer
	} else {
		pageEndContainer := container.New(layout.NewCenterLayout(),
			container.New(layout.NewVBoxLayout(),
				widget.NewLabel("Kurulum tamamlandı"),
				widget.NewButtonWithIcon("Kapat", theme.WindowCloseIcon(), func() {
					fyne.CurrentApp().Quit()
				}),
			),
		)
		return pageEndContainer
	}
}
