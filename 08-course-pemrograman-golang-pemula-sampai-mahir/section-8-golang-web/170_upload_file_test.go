package belajar_golang_web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

//go:embed templates/*.gohtml
var uploadFileTemplate embed.FS
var loadUploadFileTemplate = template.Must(template.ParseFS(uploadFileTemplate, "templates/*.gohtml"))

// uploadForm:
func uploadForm(writer http.ResponseWriter, request *http.Request) {
	if err := loadUploadFileTemplate.ExecuteTemplate(writer, "upload_file_form.gohtml", nil); err != nil {
		panic(err)
	}
}

// uploadProcess
func uploadProcess(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create(filepath.Join("./resources", fileHeader.Filename))
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(fileDestination, file); err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	loadUploadFileTemplate.ExecuteTemplate(writer, "upload_file_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

// TestUploadFormServer:
func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", uploadForm)
	mux.HandleFunc("/process", uploadProcess)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/macbook.jpeg
var macbook []byte

// TestUploadFile:
func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Noval Wardhana")
	file, err := writer.CreateFormFile("file", fmt.Sprintf("macbook_%s.jpeg", time.Now().Format("20060102150405")))
	if err != nil {
		panic(err)
	}
	file.Write(macbook)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()
	uploadProcess(recorder, request)

	response := recorder.Result()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}

//go:embed templates/*.gohtml
var uploadMultipleFile embed.FS
var loadUploadMultipleFile = template.Must(template.ParseFS(uploadMultipleFile, "templates/*.gohtml"))

// TestUploadMultipleFileServer:
func TestUploadMultipleFileServer(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		loadUploadMultipleFile.ExecuteTemplate(writer, "upload_multiple_file_form.gohtml", map[string]interface{}{
			"Title": "Upload Multiple File",
		})
	})

	mux.HandleFunc("/process", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseMultipartForm(12); err != nil {
			panic(err)
		}

		file1, fileHeader1, err := request.FormFile("file1")
		if err != nil {
			panic(err)
		}
		fileDestination1, _ := os.Create(filepath.Join("./resources", fileHeader1.Filename))
		io.Copy(fileDestination1, file1)

		file2, fileHeader2, err := request.FormFile("file2")
		if err != nil {
			panic(err)
		}
		fileDestination2, _ := os.Create(filepath.Join("./resources", fileHeader2.Filename))
		io.Copy(fileDestination2, file2)

		loadUploadMultipleFile.ExecuteTemplate(writer, "upload_multiple_file_success.gohtml", map[string]interface{}{
			"Title": "Upload Multiple File",
			"Name":  request.FormValue("name"),
			"File1": fileHeader1.Filename,
			"File2": fileHeader2.Filename,
		})
	})

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/arsenal.png
var arsenal []byte

//go:embed resources/real_madrid.png
var realMadrid []byte

// TestUploadMultipleFile:
func TestUploadMultipleFileNoServer(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* File 1 */
	go func() {
		defer wg.Done()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		name, err := writer.CreateFormField("name")
		if err != nil {
			panic(err)
		}
		name.Write([]byte("Arsenal FC"))
		file, err := writer.CreateFormFile("file", "arsenal2.png")
		if err != nil {
			panic(err)
		}
		file.Write(arsenal)
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "localhost:8080", body)
		request.Header.Add("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			if err := request.ParseMultipartForm(12); err != nil {
				panic(err)
			}
			file, fileHeader, err := request.FormFile("file")
			if err != nil {
				panic(err)
			}
			fileDestincation, err := os.Create(filepath.Join("./resources", fileHeader.Filename))
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(fileDestincation, file); err != nil {
				panic(err)
			}
			fmt.Fprintf(writer, fmt.Sprintf("Success upload %s", request.FormValue("name")))
		}
		handler(recorder, request)

		response := recorder.Result()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(responseBody))
	}()

	/* File 2 */
	go func() {
		defer wg.Done()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		name, err := writer.CreateFormField("name")
		if err != nil {
			panic(err)
		}
		name.Write([]byte("Real Madrid 2"))
		file, err := writer.CreateFormFile("file", "real_madrid2.png")
		if err != nil {
			panic(err)
		}
		file.Write(realMadrid)
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", body)
		request.Header.Add("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			if err := request.ParseMultipartForm(12); err != nil {
				panic(err)
			}
			file, fileHeader, err := request.FormFile("file")
			if err != nil {
				panic(err)
			}
			fileDestination, err := os.Create(filepath.Join("./resources", fileHeader.Filename))
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(fileDestination, file); err != nil {
				panic(err)
			}
			fmt.Fprintf(writer, fmt.Sprintf("Success upload %s", request.FormValue("name")))
		}
		handler(recorder, request)

		response := recorder.Result()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(responseBody))
	}()

	wg.Wait()

	fmt.Println("Finish")

}
