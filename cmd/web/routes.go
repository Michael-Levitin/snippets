package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux() //маршрутизатор HTTP запросов
	// Регистрируем обработчики и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	// Используем методы из структуры в качестве обработчиков маршрутов.
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	// HandleFunc() - адаптер для превращения функции в обработчик
	// - добавляет медот ServeHTTP() в функцию
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
