package utils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	dasboardID         = "666"
	apiKey             = "1234"
	appKey             = "5678"
	dumpFilePermission = "-rw-------"
	expectedContent    = "aaaa"
	expectedPrettyJSON = `{
	"a": "b"
}`
)

// TestGetDashboard test function and expect an error
func TestGetDashboardReturnError(t *testing.T) {

	_, _, err := getDashboard("https://myddenpoint", dasboardID, apiKey, appKey)

	// We should have an error
	if err == nil {
		t.Errorf("getDashboard() didn't return an error")
	}
}

// TestGetDashboardWrongResponseStatus test function of correct management of return code
func TestGetDashboardWrongResponseStatus(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	_, statusCode, err := getDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboard() should not have return an error")
	}

	if statusCode != 503 {
		t.Errorf("getDashboard() should have 503 code")
	}
}

// TestGetDashboardAssertRequest test function to ensure that the request send is correct
func TestGetDashboardAssertRequest(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard/"+dasboardID+"?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did got expected uri, got '%s'", r.RequestURI)
		}

	}))

	defer ts.Close()

	_, _, err := getDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboard() should not have return an error")
	}
}

func TestDumpDashboard(t *testing.T) {

	expectedContentByte := []byte(expectedContent)
	notExistingFilePath := "/notexistingpath/myfile.json"
	existingFilePath := "/tmp/dogsitter_coverage_test_file1.txt"

	err := dumpDashboard(expectedContentByte, notExistingFilePath)

	if err == nil {
		t.Errorf("dumpDashboard() didn't return an error")
	}

	err = dumpDashboard(expectedContentByte, existingFilePath)

	if err != nil {
		t.Errorf("dumpDashboard() returned an error")
	}

	info, err := os.Stat(existingFilePath)

	if err == nil {
		mode := info.Mode().Perm().String()

		if mode != dumpFilePermission {
			t.Errorf("File has been created with wrong permissions expected '%s' got '%s'.", dumpFilePermission, mode)
		}

		content, _ := ioutil.ReadFile(existingFilePath)

		if string(content) != expectedContent {
			t.Errorf("dumpDashboard() content is not as expected. Expected '%s' got '%s'", string(content), expectedContent)
		}

	} else if os.IsNotExist(err) {
		t.Errorf("File %s has not been created", existingFilePath)
	} else {
		t.Errorf("Unknown error while testing file %s", existingFilePath)
	}
}

func TestBeautify(t *testing.T) {

	payload := `{"a":"b"}`

	prettyPayload := beautify(payload)

	if string(prettyPayload) != expectedPrettyJSON {
		t.Errorf("beautify() return is not correct. Expected '%s', got '%s'", expectedPrettyJSON, string(prettyPayload))
	}
}
