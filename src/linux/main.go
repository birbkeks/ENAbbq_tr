package main

import (
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
	w := a.NewWindow("Установщик русификатора для ENA: Dream BBQ")
	w.Resize(fyne.NewSize(800, 600))
	a.Settings().SetTheme(theme.DarkTheme())

	mainLabel := widget.NewLabel("Добро пожаловать в установщик русификатора для ENA: Dream BBQ")

	teamLabel := canvas.NewText("  by BARBEQUE TEAM", color.RGBA{169, 169, 169, 255})
	teamLabel.TextSize = 12

	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButton("Продолжить", func() {
		w.SetContent(page2(w))
	})

	page1 := container.New(layout.NewVBoxLayout(),
		mainLabel,
		teamLabel,
		btnContinue,
		errorLabel,
	)

	btnContinue.Disable()

	page1 = container.New(layout.NewCenterLayout(), page1)

	checkIntegrity(btnContinue, errorLabel)

	w.SetContent(page1)
	w.ShowAndRun()
}

func page2(w fyne.Window) *fyne.Container {
	var path string
	appDir, _ := os.Getwd()
	steamIcon, _ := fyne.LoadResourceFromPath(filepath.Join(appDir, "resources", "steamIcon.png"))

	labelPath := widget.NewLabel("")
	errorLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})

	btnContinue := widget.NewButtonWithIcon("Установить", theme.DownloadIcon(), func() {
		w.SetContent(page3(w, path))
	})
	btnContinue.Disable()

	buttonsContainer := container.New(layout.NewVBoxLayout())

	btnSteam := widget.NewButtonWithIcon("Steam", steamIcon, func() {
		homeDir := os.Getenv("HOME")
		path = filepath.Join(homeDir, ".steam", "root", "steamapps", "common", "ENA Dream BBQ")
		checkExecutable(path, btnContinue, errorLabel)
		labelPath.SetText("Выбранный путь: " + path)
	})

	btnBrowse := widget.NewButtonWithIcon("Открыть", theme.SearchIcon(), func() {
		browseFile(w, func(selectedPath string) {
			path = selectedPath
			checkExecutable(path, btnContinue, errorLabel)
			labelPath.SetText("Выбранный путь: " + path)
		})
	})

	buttonsContainer.Add(btnSteam)
	buttonsContainer.Add(btnBrowse)

	page2 := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Выберите путь до игры. Если она установлена через Steam нажмите на кнопку Steam."),
		buttonsContainer,
		labelPath,
		errorLabel,
		btnContinue,
	)

	page2 = container.New(layout.NewCenterLayout(), page2)

	return page2
}

func checkIntegrity(btnContinue *widget.Button, errorLabel *canvas.Text) {
	appDir, _ := os.Getwd()
	resourcesPath := filepath.Join(appDir, "resources", "yarnmeta")
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		btnContinue.Disable()
		errorLabel.Text = "[ERROR]: \"resources\" не найдено, пожалуйста распакуйте архив с установщиком."
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
		errorLabel.Text = "[ERROR]: \"ENA-4-DreamBBQ.exe\" не найден, выберите папку с исполняемым файлом игры"
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

func page3(w fyne.Window, path string) *fyne.Container {
	installer(path)
	page3Container := container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Спасибо за установку"),
			widget.NewButtonWithIcon("Закрыть", theme.WindowCloseIcon(), func() {
				fyne.CurrentApp().Quit()
			}),
		),
	)
	return page3Container
}
