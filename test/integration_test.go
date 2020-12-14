package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"rentauto/controllers"
	"rentauto/models"
	"rentauto/repositories"
	"strings"
	"sync"
	"testing"
)

type TestArray struct {
	Data []TestData `json:"data"`
}

type TestData struct {
	Body   string `json:"body"`
	RType  string `json:"type"`
	URL    string `json:"url"`
	Expect string `json:"expect"`
}

var wg sync.WaitGroup

func TestService(t *testing.T) {
	cmdExec(t, "docker-compose", "down")
	cmdExec(t, "docker-compose", "up", "--build", "-d", "app")
	models.Init()
	wg.Add(3)
	defer func() {
		wg.Wait()
		cmdExec(t, "docker-compose", "down")
	}()

	testAccount(t)
	testCar(t)
	testInvoice(t)
}

func cmdExec(t *testing.T, args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		t.Fatalf("cannot %s due to %s", args, err.Error())
	}
	t.Logf("success %s for the integration test\n", args)
}

func loadJSON(t *testing.T, path string, entity interface{}) {
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(byteValue))
		err = json.Unmarshal(byteValue, entity)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func getResult(t *testing.T, w *httptest.ResponseRecorder, data TestData) {
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Results")
	t.Log(resp.StatusCode)
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(string(body))

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid status code %d. Expected: %d", resp.StatusCode, http.StatusOK)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Fatalf("Invalid content-type %s. Expected: %s", contentType, expectedContentType)
	}

	if !strings.Contains(string(body), data.Expect) {
		t.Fatalf("URL: %s, RequestBody: %s\nInvalid body: %s.  Expected: %s", data.URL, data.Body, string(body), data.Expect)
	}
}

func testAccount(t *testing.T) {
	fmt.Println("Collect test Account data")
	ta := TestArray{}
	loadJSON(t, "test_account_data.json", &ta)

	fmt.Println("Make controller")
	repAccount := repositories.NewAccountRepository()
	repSession := repositories.NewSessionRepository(repAccount)
	authController := controllers.NewAuthController(repAccount, repSession)

	var testRequest = func(data TestData) {
		requestBody := strings.NewReader(data.Body)

		r := httptest.NewRequest(data.RType, "http://localhost/"+data.URL, requestBody)
		w := httptest.NewRecorder()

		fmt.Println("Make handler")
		switch data.URL {
		case "account/login":
			authController.Authenticate(w, r)
		case "account/create":
			authController.CreateAccount(w, r)
		case "account/get":
			authController.GetAccount(w, r)
		case "account/update":
			authController.UpdateAccount(w, r)
		case "account/delete":
			authController.DeleteAccount(w, r)
		default:
			t.Fatal("Lost handler")
		}

		getResult(t, w, data)
	}

	fmt.Println("Make request")
	t.Logf("Len data: %d", len(ta.Data))
	for i := 0; i < len(ta.Data); i++ {
		testRequest(ta.Data[i])
	}

	fmt.Println("Finish testAccount")
	wg.Done()
}

func testCar(t *testing.T) {
	fmt.Println("Collect test Car data")
	ta := TestArray{}
	loadJSON(t, "test_car_data.json", &ta)

	fmt.Println("Make controller")
	var testRequest = func(data TestData) {
		requestBody := strings.NewReader(data.Body)

		r := httptest.NewRequest(data.RType, "http://localhost/"+data.URL, requestBody)
		w := httptest.NewRecorder()

		fmt.Println("Make handler")
		switch data.URL {
		case "carbrand/create":
			controllers.
				NewManageController(models.NewCarBrand(), models.NewCarModel(), models.NewCar()).
				CreateCarBrand(w, r)
		case "carmodel/create":
			controllers.
				NewManageController(models.NewCarBrand(), models.NewCarModel(), models.NewCar()).
				CreateCarModel(w, r)
		case "car/create":
			controllers.
				NewManageController(models.NewCarBrand(), models.NewCarModel(), models.NewCar()).
				CreateCar(w, r)
		case "car/get":
			controllers.
				NewManageController(models.NewCarBrand(), models.NewCarModel(), models.NewCar()).
				GetCar(w, r)
		default:
			t.Fatal("Lost handler")
		}
		getResult(t, w, data)
	}

	fmt.Println("Make request")
	t.Logf("Len data: %d", len(ta.Data))
	for i := 0; i < len(ta.Data); i++ {
		testRequest(ta.Data[i])
	}

	fmt.Println("Finish testCar")
	wg.Done()
}

func testInvoice(t *testing.T) {
	fmt.Println("Collect test Invoice data")
	ta := TestArray{}
	loadJSON(t, "test_invoice_data.json", &ta)

	fmt.Println("Make controller")
	var testRequest = func(data TestData) {
		requestBody := strings.NewReader(data.Body)

		r := httptest.NewRequest(data.RType, "http://localhost/"+data.URL, requestBody)
		w := httptest.NewRecorder()

		fmt.Println("Make handler")
		switch data.URL {
		case "invoice/create":
			controllers.NewInvoiceController(models.NewInvoice()).CreateInvoice(w, r)
		case "invoice/get":
			controllers.NewInvoiceController(models.NewInvoice()).GetInvoice(w, r)
		case "invoice/close":
			controllers.NewInvoiceController(models.NewInvoice()).CloseInvoice(w, r)
		default:
			t.Fatal("Lost handler")
		}
		getResult(t, w, data)
	}

	fmt.Println("Make request")
	t.Logf("Len data: %d", len(ta.Data))
	for i := 0; i < len(ta.Data); i++ {
		testRequest(ta.Data[i])
	}
	fmt.Println("Finish testInvoice")
	wg.Done()
}
