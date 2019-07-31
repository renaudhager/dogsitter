package commands

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

// TestGetDashboardList test function for getDashboardList()
func TestGetDashboardList(t *testing.T) {
	var expectedDashboardList DashboardList

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

		w.Write([]byte(datadogSuccessfullDashboardListResponse))

	}))

	defer ts.Close()

	la := NewListAction()

	dashboardList, err := la.getDashboardList(ts.URL, apiKey, appKey)

	if err != nil {
		t.Errorf("getDashboardList() should not have returned an error")
	}

	_ = json.Unmarshal([]byte(datadogSuccessfullDashboardListResponse), &expectedDashboardList)

	if !reflect.DeepEqual(dashboardList, expectedDashboardList) {
		t.Errorf("getDashboardList() did not return the right list")
	}

}

// TestGetDashboardList test function for getDashboardList() when response is unsuccesfull
func TestGetDashboardListWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	defer ts.Close()

	la := NewListAction()

	_, err := la.getDashboardList(ts.URL, apiKey, appKey)

	if err == nil {
		t.Errorf("getDashboardList() should have returned an error")
	}
}

// TestOutput test function for output() with text format
func TestOutputTextFormat(t *testing.T) {
	var dashboardList DashboardList

	_ = json.Unmarshal([]byte(datadogSuccessfullDashboardListResponse), &dashboardList)

	previousStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	la := NewListAction()

	err := la.output(dashboardList, "text", false)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = previousStdout

	if err != nil {
		t.Errorf("output() should not have returned an error")
	}

	if string(out) != expectedTextOutput {
		t.Errorf("output() did not print the expected information. Expected \n%v\n, got \n%v\n", expectedTextOutput, string(out))
	}
}

// TestOutputTextFormatVerbose test function for output() with text format and verbosity enabled
func TestOutputTextFormatVerbose(t *testing.T) {
	var dashboardList DashboardList

	_ = json.Unmarshal([]byte(datadogSuccessfullDashboardListResponse), &dashboardList)

	previousStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	la := NewListAction()

	err := la.output(dashboardList, "text", true)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = previousStdout

	if err != nil {
		t.Errorf("output() should not have returned an error")
	}

	if string(out) != expectedTextVerboseOutput {
		t.Errorf("output() did not print the expected information. Expected \n%v\n, got \n%v\n", expectedTextVerboseOutput, string(out))
	}
}

// TestOutput test function for output() with text format
func TestOutputJsonFormat(t *testing.T) {
	var dashboardList DashboardList

	_ = json.Unmarshal([]byte(datadogSuccessfullDashboardListResponse), &dashboardList)

	previousStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	la := NewListAction()

	err := la.output(dashboardList, "json", false)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = previousStdout

	if err != nil {
		t.Errorf("output() should not have returned an error")
	}

	if string(out) != expectedJSONOutput {
		t.Errorf("output() did not print the expected information. Expected \n%v\n, got \n%v\n", expectedJSONOutput, string(out))
	}
}

// TestListOK test for list() with proper config
func TestListOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

		w.Write([]byte(datadogSuccessfullDashboardListResponse))

	}))

	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	set.String("dh", ts.URL, "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	la := NewListAction()

	err := la.list(context)

	if err != nil {
		t.Errorf("list() should not have returned an error")
	}
}

// TestListKO test for list() with wrong config
func TestListKO(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	set.String("dh", "wronghost", "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	la := NewListAction()

	err := la.list(context)

	if err == nil {
		t.Errorf("list() should have returned an error")
	}
}
