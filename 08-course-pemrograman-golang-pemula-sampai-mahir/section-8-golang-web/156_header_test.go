package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")

	writer.Header().Add("Created-By", "Novalita Kusuma Wardhana")
	fmt.Fprintf(writer, "Header: %s", contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", nil)
	request.Header.Add("Content-Type", "Test add content type")
	record := httptest.NewRecorder()
	RequestHeader(record, request)

	response := record.Result()
	responseByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response: ", string(responseByte))
	fmt.Println("Response Header: ", record.Header().Get("Created-By"))
}

func RequestHeaderAdvance(writer http.ResponseWriter, request *http.Request) {
	auth := request.Header.Get("Authorization")
	connection := request.Header.Get("Connection")
	contentType := request.Header.Get("Content-Type")

	fmt.Fprintf(writer, " auth: %s \n", auth)
	fmt.Fprintf(writer, " Connection: %s \n", connection)
	fmt.Fprintf(writer, " ContentType: %s \n", contentType)

	writer.Header().Add("Approve-by", "NovalWardhana")
}

func TestRequestHeaderAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000?name=Noval&name=kusuma&name=wardhana", nil)
	request.Header.Add("authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImRhdGEiOnsiaWQiOjEsInVzZXJuYW1lIjoicm9vdCIsInJvbGVzIjpbInJvb3QiXSwiaXNzdWVyX2NvZGUiOltdfSwiaWF0IjoxNjcwNDI4Mzg3LCJleHAiOjE2NzA1MDAzODd9.oiMJJLRs7KDhuy8KyUfbfMk06HcR_pvCvy749uuB5ND-VbXFKvhZfngfHMjg9N7-fSWMUENRX46g_MFbA2f2234mGXbnHzJyDRIBVZDHU3pbh1QSGyDP3SQBFZ84Y8CRrqplmsttCChOFnX7S275_6TecS4P9-GSEel21raWZSA")
	request.Header.Add("connection", "linux")
	request.Header.Add("content-type", "application/json")
	record := httptest.NewRecorder()
	RequestHeaderAdvance(record, request)

	result := record.Result()
	resultBody, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(resultBody))
	fmt.Println("Response header: ", record.Header().Get("approve-by"))
}
