package commands

import (
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/urfave/cli"
)

// TestLoadDashboard test function for loadDashboard()
func TestLoadDashboard(t *testing.T) {

	notExistingFilePath := "/notexistingpath/myfile.json"
	existingFilePath := "/tmp/dogsitter_coverage_test_file1.txt"

	pa := NewPushAction()
	_, err := pa.loadDashboard(notExistingFilePath)

	if err == nil {
		t.Errorf("loadDashboard() did not return an error")
	}

	_ = ioutil.WriteFile(existingFilePath, []byte(expectedContent), 0644)

	content, err := pa.loadDashboard(existingFilePath)

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

	pa := NewPushAction()

	err := pa.uploadDashboard(ts.URL, []byte(expectedContent), apiKey, appKey)

	if err == nil {
		t.Errorf("uploadDashboard() should have returned an error")
	}
}

// TestUploadDashboardAssertRequest test function to ensure that uploadDashboard() send the correct request
func TestUploadDashboardAssertRequest(t *testing.T) {

	// Test when response is succesfull
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Did not get expected HEADER, got %s", r.Header)
		}

		body, _ := ioutil.ReadAll(r.Body)

		if string(body) != dashboardPayload {
			t.Errorf("Did not get expected body,expected '%s' got %s", expectedPrettyJSON, string(body))
		}

		w.Write([]byte(datadogSuccessfullResponse))

	}))

	defer ts.Close()

	pa := NewPushAction()

	err := pa.uploadDashboard(ts.URL, []byte(dashboardPayload), apiKey, appKey)

	if err != nil {
		t.Errorf("uploadDashboard() should not have returned an error: %v", err)
	}

	// Test when response is failling
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("badJson"))

	}))

	defer ts.Close()

	err = pa.uploadDashboard(ts.URL, []byte(dashboardPayload), apiKey, appKey)

	if err == nil {
		t.Errorf("uploadDashboard() should  have returned an error")
	}
}

// TestGetDashboardInfo function to test getDashboardInfo()
func TestGetDashboardInfo(t *testing.T) {

	pa := NewPushAction()

	id, url, err := pa.getDashboardInfo(datadogSuccessfullResponse)

	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	if id != "dnq-s5w-h5j" {
		t.Errorf("Did not get expected id, got %s", id)
	}

	if url != "/dashboard/dnq-s5w-h5j/sre---consul-overview" {
		t.Errorf("Did not get expected url, got %s", url)
	}

	id, url, err = pa.getDashboardInfo("badJson")

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if len(id) != 0 {
		t.Errorf("id should be empty, got %s", id)
	}

	if len(url) != 0 {
		t.Errorf("url should be empty, got %s", url)
	}
}

// TestPushOK test for pull() with proper config
func TestPushOK(t *testing.T) {

	_ = ioutil.WriteFile("/tmp/test-push", []byte(expectedContent), 0644)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(datadogSuccessfullResponse))

	}))

	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	set.String("dh", ts.URL, "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")
	set.String("f", "/tmp/test-push", "doc")

	context := cli.NewContext(nil, set, nil)

	pa := NewPushAction()

	err := pa.push(context)

	if err != nil {
		t.Errorf("push() should not have returned an error")
	}

	// Cleaning up
	err = os.Remove("/tmp/test-push")
	if err != nil {
		t.Errorf("Error while cleaning up %s", "/tmp/test-push")
	}
}

// TestPushKO test for pull() with wrong config
func TestPushKO(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	set.String("dh", "wronghost", "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	pa := NewPushAction()

	err := pa.push(context)

	if err == nil {
		t.Errorf("push() should have returned an error")
	}
}
