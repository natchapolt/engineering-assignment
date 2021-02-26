package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"syscall"
	"text/template"
	"time"

	env "github.com/joho/godotenv"
)

const envFile = ".env"
const dataFile = "data/forms.json"

var loadEnv = env.Load

type formInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (f formInput) validate() error {
	if f.FirstName == "" || f.LastName == "" || f.Email == "" || f.PhoneNumber == "" {
		return errors.New("invalid input")
	}
	return nil
}

func (f formInput) save() error {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}
	var forms []formInput
	err = json.Unmarshal(file, &forms)
	if err != nil {
		return err
	}
	forms = append(forms, f)
	toSave, err := json.Marshal(forms)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dataFile, toSave, os.ModeAppend)
	return err
}

func handleFunc(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(resp, err.Error())
			return
		}
		f := formInput{
			FirstName:   req.PostFormValue("first_name"),
			LastName:    req.PostFormValue("last_name"),
			Email:       req.PostFormValue("email"),
			PhoneNumber: req.PostFormValue("phone_number"),
		}
		err = f.validate()
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(resp, err.Error())
			return
		}
		err = f.save()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(resp, err.Error())
			return
		}
		resp.WriteHeader(http.StatusOK)
		fmt.Fprintln(resp, "<p>form saved</p>")
		fmt.Fprintln(resp, "<a href=\"http://localhost:8080\">back home</a>")
	case http.MethodGet:
		filename := fmt.Sprint(req.URL)
		if filename == "/" {
			filename = "index.html"
		} else if filename[0] == '/' {
			filename = filename[1:]
		}

		// get templated-html page path
		pg := filepath.Join("static", filename)
		tmpl, _ := template.ParseFiles(pg)

		// Return a 404 if the template doesn't exist
		_, err := os.Stat(pg)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(resp, req)
				return
			}
		}

		// build table header for index.html
		if filename == "index.html" {
			file, err := ioutil.ReadFile(dataFile)
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(resp, err.Error())
				return
			}
			var forms []formInput
			err = json.Unmarshal(file, &forms)
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(resp, err.Error())
				return
			}

			theader, trows := "", ""
			for idx, form := range forms {
				v := reflect.ValueOf(form)
				typeOfForm := v.Type()

				trows += "<tr>"
				for i := 0; i < v.NumField(); i++ {
					if idx == 0 {
						theader += "<th>" + fmt.Sprintf("%s", typeOfForm.Field(i).Name) + "</th>"
					}
					trows += "<td>" + fmt.Sprintf("%v", v.Field(i).Interface()) + "</td>"
				}
				trows += "</tr>"
			}
			theader = "<tr>" + theader + "</tr>"

			if len(forms) > 0 {
				tmpl.Parse("{{define \"table-header\"}}" + theader + "{{end}}")
				tmpl.Parse("{{define \"table-rows\"}}" + trows + "{{end}}")
			}
		}
		resp.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(resp, "page", nil)
	default:
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}
}

func run() (s *http.Server) {
	err := loadEnv(envFile)
	if err != nil {
		log.Fatal(err)
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("no port specified")
	}
	port = fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleFunc)

	s = &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return
}

func main() {
	s := run()
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}
	log.Println("Server exiting")
}
