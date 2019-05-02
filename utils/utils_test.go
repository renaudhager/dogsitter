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

// TestGetDashboardWrongResponseStatus test function to ensure that getDashboard() handle correctly return code
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

// TestGetDashboardAssertRequest test function to ensure that getDashboard() send the correct request
func TestGetDashboardAssertRequest(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard/"+dasboardID+"?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

	}))

	defer ts.Close()

	_, _, err := getDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboard() should not have return an error")
	}
}

// TestDumpDashboard test function for dumpDashboard()
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

		// Cleaning up
		err = os.Remove(existingFilePath)

		if err != nil {
			t.Errorf("Error while cleaning up %s", existingFilePath)
		}

	} else if os.IsNotExist(err) {
		t.Errorf("File %s has not been created", existingFilePath)
	} else {
		t.Errorf("Unknown error while testing file %s", existingFilePath)
	}
}

// TestDumpDashboard test function for beautify()
func TestBeautify(t *testing.T) {

	payload := `{"a":"b"}`

	prettyPayload := beautify(payload)

	if string(prettyPayload) != expectedPrettyJSON {
		t.Errorf("beautify() return is not correct. Expected '%s', got '%s'", expectedPrettyJSON, string(prettyPayload))
	}
}

// TestLoadDashboard test function for loadDashboard()
func TestLoadDashboard(t *testing.T) {

	notExistingFilePath := "/notexistingpath/myfile.json"
	existingFilePath := "/tmp/dogsitter_coverage_test_file1.txt"

	_, err := loadDashboard(notExistingFilePath)

	if err == nil {
		t.Errorf("loadDashboard() did not return an error")
	}

	_ = ioutil.WriteFile(existingFilePath, []byte(expectedContent), 0644)

	content, err := loadDashboard(existingFilePath)

	if err != nil {
		t.Errorf("loadDashboard() did return an error while reading %s", existingFilePath)
	}
	if string(content) != expectedContent {
		t.Errorf("loadDashboard() did return the expected content, expected '%s' got '%s'", content, expectedContent)
	}

	// Cleaning up
	err = os.Remove(existingFilePath)

	if err != nil {
		t.Errorf("Error while cleaning up %s", existingFilePath)
	}
}

// TestUploadDashboard test function for uploadDashboard() handle correctly return code.
func TestUploadDashboardWrongResponseStatus(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	err := uploadDashboard(ts.URL, []byte(expectedContent), apiKey, appKey)

	if err == nil {
		t.Errorf("uploadDashboard() should have returned an error")
	}
}

// TestUploadDashboardAssertReques test function to ensure that uploadDashboard() send the correct request
func TestUploadDashboardAssertReques(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did got expected uri, got '%s'", r.RequestURI)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Did not get expected HEADER, got %s", r.Header)
		}

		body, _ := ioutil.ReadAll(r.Body)

		if string(body) != expectedPrettyJSON {
			t.Errorf("Did not get expected body,expected '%s' got %s", expectedPrettyJSON, string(body))
		}

	}))

	defer ts.Close()

	err := uploadDashboard(ts.URL, []byte(expectedPrettyJSON), apiKey, appKey)

	if err != nil {
		t.Errorf("uploadDashboard() should not have returned an error")
	}
}
