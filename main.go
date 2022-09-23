package main

import (
	"fyne.io/fyne/v2"
	fapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	app        fyne.App
	baseWindow fyne.Window
	infoWindow fyne.Window
}

func (a *App) BaseScreen() {
	a.baseWindow = a.app.NewWindow("Лабораторная работа") // Задаём название для главного окна
	a.baseWindow.Resize(fyne.NewSize(900, 600))           // Устанавливаем размеры главного окна
	a.baseWindow.CenterOnScreen()
	a.baseWindow.SetFixedSize(true)
}
func (a *App) BaseScreen1() {
	a.SetBackgroundImage(a.baseWindow)
	// Записываем текст в переменную, устанавливаем шрифт и устанавливаем позиционирование текста по ценрту экрана
	const title = "Защита компьютерной информации"
	label := widget.NewLabel(title)
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	//Записываем текст в переменную, устанавливаем шрифт и прикрепляем к правому краю окна
	const college = "УО 'Минский Государственный Колледж Электроники'"
	txtInfo := widget.NewLabel(college)
	txtInfo.TextStyle = fyne.TextStyle{Bold: true}
	txtInfo.Alignment = fyne.TextAlignTrailing

	//txtInfo1 := container.NewWithoutLayout(txtInfo)

	// Кнопка открытия справки
	const titleInfo = "Справка"
	info := widget.NewButton(titleInfo, func() {
		a.InfoWindow()
	})

	btnFullScreen := container.NewVBox(
		widget.NewButton("123", func() {
			a.SetFullScreen()
		}))

	//Кнопка для закрытия всей программы
	const exit = "Выход"
	btnExit := container.NewHBox(
		info,
		btnFullScreen,
		layout.NewSpacer(),
		widget.NewButton(exit, func() {
			a.app.Quit()
		}),
	)

	const workLabel = "Просмотр работ"
	workView := container.NewGridWithColumns(3, //Создаём контейнер с 3 колонками
		layout.NewSpacer(), // В 1 колонку устанавливаем пустое пространство
		container.NewVBox( // Создаём вертикальный контейнер во 2 колонку
			widget.NewButton(workLabel, func() {
				a.LabsScreen()
				// Создаём новую кноку
			})),
		layout.NewSpacer()) // Устанавливаем пустое пространство в 3 колонку

	//Создаём кнопку для главного меню
	btnMenu := container.NewVBox(
		label,
		layout.NewSpacer(),
		workView,
		layout.NewSpacer(),
		txtInfo,
		btnExit)

	//В главное окно загружаем фоновую картинку и кнопку
	a.baseWindow.SetContent(container.New(layout.NewMaxLayout(),
		canvas.NewImageFromFile("123.jpg"),
		btnMenu))

	a.baseWindow.Show()      // Показываем главное окно
	a.baseWindow.SetMaster() // Устанавливаем для главного окна тег мастер, без него, после закрытия дочернего окна, закроется и главное
	//a.app.Run()
}

func (a *App) LabsScreen() {
	left := container.NewVBox(
		widget.NewButton("Лабораторная работа №1", func() {}),
		widget.NewButton("Лабораторная работа №3", func() {}),
		widget.NewButton("Лабораторная работа №5", func() {}),
		widget.NewButton("Лабораторная работа №7", func() {}),
		widget.NewButton("Лабораторная работа №9", func() {}),
		widget.NewButton("Лабораторная работа №11", func() {}),
	)
	right := container.NewVBox(
		widget.NewButton("Лабораторная работа №2", func() {}),
		widget.NewButton("Лабораторная работа №4", func() {}),
		widget.NewButton("Лабораторная работа №6", func() {}),
		widget.NewButton("Лабораторная работа №8", func() {}),
		widget.NewButton("Лабораторная работа №10", func() {}),
		widget.NewButton("Лабораторная работа №12", func() {}),
	)
	//Группируем кнопки с заданиями в 2 колонки

	work := container.NewGridWithColumns(2,
		left,
		right,
		layout.NewSpacer(), // Пустое пространство
	)

	const labelMenu = "Меню"
	btnMenu := container.NewGridWithColumns(3,
		layout.NewSpacer(),
		widget.NewButton(labelMenu, func() {
			a.BaseScreen1()
		}),
		layout.NewSpacer())

	//Создаём кнопку для главного меню
	btnCont := container.NewVBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		work,
		btnMenu)

	//В главное окно загружаем фоновую картинку и кнопку
	a.baseWindow.SetContent(container.New(layout.NewMaxLayout(),
		canvas.NewImageFromFile("123.jpg"),
		btnCont))

	a.baseWindow.Show() // Показываем главное окно

}

func (a *App) InfoWindow() {
	a.infoWindow = a.app.NewWindow("Справка")   // Создаём и задаём название для дочернего окна
	a.infoWindow.Resize(fyne.NewSize(480, 360)) // Устанавливаем размеры дочернего окна
	a.infoWindow.CenterOnScreen()

	//Записываем текст в переменную и устанавливаем перенос слов в зависимости от размера окна
	text := widget.NewLabel("Приложение разработано учащимися Минского Государственного Колледжа Электроники, группы 53ТП Пожарицким Максимом и Старовойтовым Степаном")
	text.Wrapping = fyne.TextWrapWord

	a.infoWindow.SetContent(
		text)

	a.infoWindow.Show()

}

func (a *App) SetFullScreen() {
	if a.baseWindow.FullScreen() {
		a.baseWindow.SetFullScreen(false)
	} else {
		a.baseWindow.SetFullScreen(true)
	}
	return
}

func (a *App) SetBackgroundImage(window fyne.Window) {
	background := canvas.NewImageFromFile("proj/123.jpg")
	window.SetContent(
		container.New(layout.NewMaxLayout(),
			background))
}

func newApp() App {
	return App{
		app: fapp.New(),
	}
}

func main() {
	// Создаём приложение
	var a = newApp()
	a.BaseScreen()
	a.BaseScreen1()
	a.app.Run()

	/*
		btnClose := widget.NewButton("Закрыть", func() { // Создаём кнопку закрытия справки
			w1.Hide() // Обязательно используем ф-ю Hide, она позволяет скрыть окно и открывать его много раз
		})

		//Создаём кнопки с заданиями
		left := container.NewVBox(
			widget.NewButton("Лабораторная работа №1", func() {}),
			widget.NewButton("Лабораторная работа №3", func() {}),
			widget.NewButton("Лабораторная работа №5", func() {}),
			widget.NewButton("Лабораторная работа №7", func() {}),
			widget.NewButton("Лабораторная работа №9", func() {}),
			widget.NewButton("Лабораторная работа №11", func() {}),
		)
		right := container.NewVBox(
			widget.NewButton("Лабораторная работа №2", func() {}),
			widget.NewButton("Лабораторная работа №4", func() {}),
			widget.NewButton("Лабораторная работа №6", func() {}),
			widget.NewButton("Лабораторная работа №8", func() {}),
			widget.NewButton("Лабораторная работа №10", func() {}),
			widget.NewButton("Лабораторная работа №12", func() {}),
		)
		//Группируем кнопки с заданиями в 2 колонки
		work := container.NewGridWithColumns(2,
			left,
			right,
			layout.NewSpacer(), // Пустое пространство
		)

		workView := container.NewGridWithColumns(3, //Создаём контейнер с 3 колонками
			layout.NewSpacer(), // В 1 колонку устанавливаем пустое пространство
			container.NewVBox( // Создаём вертикальный контейнер во 2 колонку
				widget.NewButton("Просмотр работ", func() { // Создаём новую кноку
					w.SetContent( // По нажатию на кнопку устанавливаем содержимое главного окна
						container.New(layout.NewMaxLayout(), // Создаём новый пустой контейнер и устанавливаем слой во всю ширину окна
							canvas.NewImageFromFile("proj/123.jpg"), // Устанавливаем изображение на фон
							container.NewVBox( //Создаём вертикальный контейнер
								layout.NewSpacer(), // Создаём пустое пространство
								work,               // Устанавливаем кнопку с заданиями
								btnExit)))          // Устанавливаем кнопку выхода и справки
				})),
			layout.NewSpacer()) // Устанавливаем пустое пространство в 3 колонку

		//Создаём кнопку для справки
		btnText := container.NewVBox(
			text,
			layout.NewSpacer(),
			btnClose) // Кнопка закрытия окна справки

		w1.SetContent(container.New(layout.NewMaxLayout(), // Устанавливаем в окно справки новый слой с шириной окна
			canvas.NewImageFromFile("123.jpg"),
			btnText))

		//Создаём кнопку для главного меню
		btnMenu := container.NewVBox(
			label,
			layout.NewSpacer(),
			workView,
			layout.NewSpacer(),
			txtInfo,
			btnExit)

		//В главное окно загружаем фоновую картинку и кнопку
		w.SetContent(container.New(layout.NewMaxLayout(),
			canvas.NewImageFromFile("123.jpg"),
			btnMenu))

		w.Show()      // Показываем главное окно
		w.SetMaster() // Устанавливаем для главного окна тег мастер, без него, после закрытия дочернего окна, закроется и главное
		a.Run()       // Запускаем приложение


	*/
}
