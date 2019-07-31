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

// TestGetDashboard test function and expect an error
func TestGetDashboardReturnError(t *testing.T) {

	pa := NewPullAction()
	_, _, err := pa.getDashboard("https://myddenpoint", dasboardID, apiKey, appKey)

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

	pa := NewPullAction()

	_, statusCode, err := pa.getDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboard() should not have returned an error")
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

		w.Write([]byte(datadogSuccessfullGetDashboard))

	}))

	defer ts.Close()

	pa := NewPullAction()

	payload, _, err := pa.getDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboard() should not have returned an error")
	}

	if payload != expectedPayloadGetDashboard {
		t.Errorf("getDashboard() did not returned the expected payload")
	}
}

// TestDumpDashboard test function for dumpDashboard()
func TestDumpDashboard(t *testing.T) {

	expectedContentByte := []byte(expectedContent)
	notExistingFilePath := "/notexistingpath/myfile.json"
	existingFilePath := "/tmp/dogsitter_coverage_test_file1.txt"

	pa := NewPullAction()

	err := pa.dumpDashboard(expectedContentByte, notExistingFilePath)

	if err == nil {
		t.Errorf("dumpDashboard() didn't return an error")
	}

	err = pa.dumpDashboard(expectedContentByte, existingFilePath)

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

	pa := NewPullAction()

	prettyPayload := pa.beautify(payload)

	if string(prettyPayload) != expectedPrettyJSON {
		t.Errorf("beautify() return is not correct. Expected '%s', got '%s'", expectedPrettyJSON, string(prettyPayload))
	}
}

// TestStripBadField function that test stripBadField()
func TestStripBadField(t *testing.T) {
	input := `{"a":"b","c":"d","e":"f"}`
	expectedOutput := `{"a":"b","e":"f"}`

	pa := NewPullAction()

	output, err := pa.stripBadField([]byte(input), "c")

	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	if string(output) != expectedOutput {
		t.Errorf("output should be `%v`, got `%v`", expectedOutput, string(output))
	}

	_, err = pa.stripBadField([]byte("foo"), "c")

	if err == nil {
		t.Errorf("err should not be nil")
	}
}

// TestPullOK test for pull() with proper config
func TestPullOK(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(datadogSuccessfullGetDashboard))

	}))

	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	set.String("dh", ts.URL, "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")
	set.String("id", dasboardID, "doc")
	set.String("o", "/tmp/test", "doc")

	context := cli.NewContext(nil, set, nil)
	pa := NewPullAction()

	err := pa.pull(context)

	if err != nil {
		t.Errorf("pull() should not have returned an error, %v", err)
	}
	// Cleaning up
	err = os.Remove("/tmp/test")
	if err != nil {
		t.Errorf("Error while cleaning up %s", "/tmp/test")
	}

}

// TestPullKO test for pull() with wrong config
func TestPullKO(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	set.String("dh", "wronghost", "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	pa := NewPullAction()

	err := pa.pull(context)

	if err == nil {
		t.Errorf("pull() should have returned an error")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	defer ts.Close()

	set2 := flag.NewFlagSet("test2", 0)
	set2.String("dh", ts.URL, "doc")
	set2.String("api-key", apiKey, "doc")
	set2.String("app-key", appKey, "doc")

	context2 := cli.NewContext(nil, set, nil)

	err = pa.pull(context2)

	if err == nil {
		t.Errorf("pull() should have returned an error")
	}
}
