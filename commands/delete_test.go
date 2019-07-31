package commands

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli"
)

// TestDeleteDashboard test function for deleteDashboard() when response is successful.
func TestDeleteDashboard(t *testing.T) {

	// Test when response is succesfull
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "DELETE" {
			t.Errorf("Expected 'DELTET' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard/"+dasboardID+"?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

		w.Write([]byte(datadogSuccessfullResponse))

	}))

	defer ts.Close()

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("dh", ts.URL, "doc")
	globalSet.String("api-key", apiKey, "doc")
	globalSet.String("app-key", appKey, "doc")
	globalContext := cli.NewContext(nil, globalSet, nil)

	set := flag.NewFlagSet("test", 0)
	set.String("id", dasboardID, "doc")

	context := cli.NewContext(nil, set, globalContext)

	da := NewDeleteAction()

	err := da.deleteDashboard(context)

	if err != nil {
		t.Errorf("deleteDashboard() should not have returned an error: %v", err)
	}

}

// TestDeleteDashboard test for deleteDashboard() when response is unsuccesfull
func TestDeleteDashboardWrongResponseStatus(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)

	}))

	defer ts.Close()

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("dh", ts.URL, "doc")
	globalSet.String("api-key", apiKey, "doc")
	globalSet.String("app-key", appKey, "doc")
	globalContext := cli.NewContext(nil, globalSet, nil)

	set := flag.NewFlagSet("test", 0)
	set.String("id", dasboardID, "doc")

	context := cli.NewContext(nil, set, globalContext)

	da := NewDeleteAction()

	err := da.deleteDashboard(context)

	if err == nil {
		t.Errorf("deleteDashboard() should have returned an error")
	}
}

// TestDeleteDashboardOK test for deleteDashboard() with proper config
func TestDeleteDashboardOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "DELETE" {
			t.Errorf("Expected 'DELETE' request, got '%s'", r.Method)
		}

		if r.RequestURI != "/api/v1/dashboard/"+dasboardID+"?api_key="+apiKey+"&application_key="+appKey {
			t.Errorf("Did not get expected uri, got '%s'", r.RequestURI)
		}

		w.Write([]byte(datadogSuccessfullResponse))

	}))

	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	set.String("dh", ts.URL, "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")
	set.String("id", dasboardID, "doc")

	context := cli.NewContext(nil, set, nil)

	da := NewDeleteAction()

	err := da.deleteDashboard(context)

	if err != nil {
		t.Errorf("deleteDashboard() should not have returned an error")
	}
}

// TestDeleteDashboardKO test for deleteDashboard() with wrong config
func TestDeleteDashboardKO(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	set.String("dh", "wronghost", "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	da := NewDeleteAction()

	err := da.deleteDashboard(context)

	if err == nil {
		t.Errorf("delete() should have returned an error")
	}
}
