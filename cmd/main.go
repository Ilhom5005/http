package main

import (
	"net"
	"net/http"
	"os"

	"github.com/Ilhom5005/http/cmd/app"
	"github.com/Ilhom5005/http/pkg/banners"
)

	// file, err := os.Create("web/banners/1.")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// file.Close()
// func execute(host string, port string) (err error) {
// 	srv := &http.Server{
// 		Addr: net.JoinHostPort(host, port),
// 		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			if r.Method != "POST" {
// 				return
// 			}
// 			log.Print("tags:", r.URL.Query().Get("tags"))
// 			log.Println("full URL:", r.RequestURI)                        // full URL
// 			log.Println("Method", r.Method)                               // method
// 			log.Println("all Headers:\n", r.Header)                       // all headers
// 			log.Println("specific header:", r.Header.Get("Content-Type")) // specific header
// 			log.Println("FormValue(\"tags\"): ", r.FormValue("tags"))     // только первое значение Query + POST
// 			log.Println("FormValue(\"tags\"): ", r.PostFormValue("tags")) // только первое значение POST

// 			body, err := ioutil.ReadAll(r.Body)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			log.Printf("Body:\n%s", body)

// 			err = r.ParseMultipartForm(10 * 1024 * 1024) // 10MB
// 			if err != nil {
// 				log.Println(err)
// 			}

// 			err = r.ParseForm()
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			src, err := os.Create("1.png")
// 			defer src.Close()

// 			file, _, err := r.FormFile("image")
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			defer file.Close()
// 			io.Copy(src, file)

// 			// can use only after ParseForm (or FormValue, PostFormValue)
// 			log.Println("r.Form:", r.Form)         // all value of form(excepcion file)
// 			log.Println("r.PostForm:", r.PostForm) // all value of form(excepcion file)

// 			// can use olne after ParseMultipart(FormValue, PostFromValue, auto call ParseMultipart)
// 			log.Println(r.FormFile("image"))
// 			// r.MultipartForm.Value - only "обычные поля"
// 			// r.MultipartForm.File - only files
// 		}),
// 	}
// 	return srv.ListenAndServe()
// }

// // i may repeat it :)
// func execute(host string, port string) (err error) {
// 	mux := http.NewServeMux()
// 	bannersSvc := banners.NewService()
// 	server := app.NewServer(mux, bannersSvc)
// 	srv := &http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: server,
// 	}
// 	server.Init()
// 	return srv.ListenAndServe()
// }


func main() {
	//обьявляем порт и хост
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(h, p string) error {
	mux := http.NewServeMux()

	bannerSvc := banners.NewService()

	server := app.NewServer(mux, bannerSvc)

	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(h, p),
		Handler: server,
	}
	return srv.ListenAndServe()
}

