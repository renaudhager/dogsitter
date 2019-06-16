package commands

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
	datadogSuccessfullResponse = (`{"notify_list":[],"description":"created by renaud.hager@nospam.com","author_name":"Renaud Hager","template_variables":[{"default":"eu-west-1","prefix":"location","name":"location"},{"default":"*","prefix":"environment","name":"environment"},{"default":"sre","prefix":"team","name":"team"}],"is_read_only":false,"id":"dnq-s5w-h5j","title":"SRE - Consul Overview","url":"/dashboard/dnq-s5w-h5j/sre---consul-overview","created_at":"2019-06-14T14:58:34.760504+00:00","modified_at":"2019-06-14T14:58:34.760504+00:00","author_handle":"renaud.hager@nospam.com","widgets":[{"definition":{"widgets":[{"definition":{"type":"query_value","requests":[{"q":"avg:consul.autopilot.healthy{$location,$team,$environment,group_role:consul-server}","aggregator":"avg","conditional_formats":[{"palette":"white_on_red","comparator":"<=","value":0.5},{"palette":"white_on_green","comparator":">=","value":0.9},{"palette":"white_on_yellow","comparator":"<=","value":0.89}]}],"autoscale":false,"precision":0,"title":"Overall health"},"id":118514630952118},{"definition":{"requests":[{"q":"avg:consul.kvs.apply.avg{$location,$environment,$team,group_role:consul-server}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"area"}],"type":"timeseries","title":"KV and transaction apply latency"},"id":8450620905653722},{"definition":{"requests":[{"q":"max:consul.raft.commitTime.avg{$location,$team,$environment,group_role:consul-server}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","title":"Max Raft commit time (ms)"},"id":7470639491355588},{"definition":{"requests":[{"q":"avg:consul.kvs.apply.max{$location,$team,$environment}, avg:consul.txn.apply{$location,$team,$environment}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"bars"}],"type":"timeseries","title":"Various latency related to Raft (ms)"},"id":6899892404741622},{"definition":{"requests":[{"q":"avg:consul.raft.leader.lastContact.max{*}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","title":"Last contact from leader (ms)"},"id":3355181384970402}],"layout_type":"ordered","type":"group","title":"Latency"},"id":5254103359870646},{"definition":{"widgets":[{"definition":{"requests":[{"q":"max:consul.catalog.total_nodes{$location,$environment,$team}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"max:consul.catalog.total_nodes{$location,$environment,$team}","alias_name":"Total nodes"}]},{"q":"avg:consul.peers{$location,$environment,$team}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:consul.peers{$location,$environment,$team}","alias_name":"Peer nodes"}]}],"type":"timeseries","title":"Consul nodes"},"id":77898},{"definition":{"requests":[{"q":"avg:consul.rpc.query{$location,$environment,$team,group_role:consul-server}.as_count(), avg:consul.rpc.request{$location,$environment,$team,group_role:consul-server}.as_count()","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"bars","metadata":[{"expression":"avg:consul.rpc.query{$location,$environment,$team,group_role:consul-server}.as_count()","alias_name":"rpc query"},{"expression":"avg:consul.rpc.request{$location,$environment,$team,group_role:consul-server}.as_count()","alias_name":"rpc request"}]}],"type":"timeseries","title":"RPC request/query"},"id":4566201650946910},{"definition":{"requests":[{"q":"sum:consul.serf.member.join{$location,$team,$environment,group_role:consul-server}.as_count(), sum:consul.serf.member.left{$location,$team,$environment,group_role:consul-server}.as_count(), sum:consul.serf.member.failed{$location,$team,$environment,group_role:consul-server}.as_count()","style":{"line_width":"normal","palette":"orange","line_type":"solid"},"display_type":"bars","metadata":[{"expression":"sum:consul.serf.member.join{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf join"},{"expression":"sum:consul.serf.member.failed{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf failed"},{"expression":"sum:consul.serf.member.left{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf left"}]}],"type":"timeseries","title":"Serf activity"},"id":3286132655611000}],"layout_type":"ordered","type":"group","title":"Network and Serf"},"id":286350767033660},{"definition":{"widgets":[{"definition":{"requests":[{"q":"max:consul.runtime.total_gc_pause_ns{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"area"}],"type":"timeseries","title":"Max GC time (ns)"},"id":6725921563822062},{"definition":{"requests":[{"q":"avg:consul.runtime.alloc_bytes{$location,$team,$environment,group_role:consul-server} by {host}, avg:consul.runtime.sys_bytes{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:consul.runtime.alloc_bytes{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"alloc bytes"},{"expression":"avg:consul.runtime.sys_bytes{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"sys bytes"}]}],"type":"timeseries","title":"Memory usage"},"id":497692778043364}],"layout_type":"ordered","type":"group","title":"Memory"},"id":5524458375359210},{"definition":{"widgets":[{"definition":{"requests":[{"q":"avg:system.cpu.user{$location,$environment,$team,group_role:consul-server} by {host}+avg:system.cpu.system{$location,$environment,$team,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area","metadata":[{"expression":"avg:system.cpu.user{$location,$environment,$team,group_role:consul-server} by {host}+avg:system.cpu.system{$location,$environment,$team,group_role:consul-server} by {host}","alias_name":"system+ users"}]}],"type":"timeseries","title":"CPU Usage"},"id":77901},{"definition":{"requests":[{"q":"(100*avg:system.disk.used{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs} by {host,device})/avg:system.disk.total{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs,!device:overlay} by {host,device}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"(100*avg:system.disk.used{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs} by {host,device})/avg:system.disk.total{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs,!device:overlay} by {host,device}","alias_name":"% used"}]}],"yaxis":{"max":"100","min":"0"},"type":"timeseries","title":"Disk usage"},"id":5907332623733836},{"definition":{"requests":[{"q":"avg:system.mem.usable{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:system.mem.usable{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"usable"}]},{"q":"avg:system.mem.total{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:system.mem.total{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"total"}]}],"type":"timeseries","title":"Memory Usable"},"id":7740976102574542},{"definition":{"requests":[{"q":"per_second(avg:system.net.bytes_sent{$location,$environment,$team,group_role:consul-server} by {host})","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"area","metadata":[{"expression":"per_second(avg:system.net.bytes_sent{$location,$environment,$team,group_role:consul-server} by {host})","alias_name":"bytes_sent"}]}],"type":"timeseries","title":"Bytes sent/s"},"id":6219502221062534},{"definition":{"requests":[{"q":"per_second(avg:system.net.bytes_rcvd{$location,$environment,$team,group_role:consul-server} by {host})","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"area","metadata":[{"expression":"per_second(avg:system.net.bytes_rcvd{$location,$environment,$team,group_role:consul-server} by {host})","alias_name":"bytes_rcvd"}]}],"type":"timeseries","title":"Bytes rcvd/s"},"id":3344996939592842}],"layout_type":"ordered","type":"group","title":"System"},"id":140191778868366}],"layout_type":"ordered"}`)
	dashboardPayload           = (`{
		"notify_list": [],
		"description": "created by renaud.hager@nospam.com",
		"template_variables": [
			{
				"default": "eu-west-1",
				"prefix": "location",
				"name": "location"
			},
			{
				"default": "*",
				"prefix": "environment",
				"name": "environment"
			},
			{
				"default": "sre",
				"prefix": "team",
				"name": "team"
			}
		],
		"is_read_only": false,
		"id": "hkv-c7t-fyu",
		"title": "SRE - Consul Overview",
		"url": "/dashboard/hkv-c7t-fyu/sre---consul-overview",
		"created_at": "2018-12-14T14:51:59.037558+00:00",
		"modified_at": "2019-05-01T14:47:14.687078+00:00",
		"author_handle": "renaud.hager@nospam.com"
	}
`)
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

// TestUploadDashboardAssertRequest test function to ensure that uploadDashboard() send the correct request
func TestUploadDashboardAssertRequest(t *testing.T) {

	// Test when response is succesfull
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		if string(body) != dashboardPayload {
			t.Errorf("Did not get expected body,expected '%s' got %s", expectedPrettyJSON, string(body))
		}

		w.Write([]byte(datadogSuccessfullResponse))

	}))

	defer ts.Close()

	err := uploadDashboard(ts.URL, []byte(dashboardPayload), apiKey, appKey)

	if err != nil {
		t.Errorf("uploadDashboard() should not have returned an error: %v", err)
	}

	// Test when response is failling
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("badJson"))

	}))

	defer ts.Close()

	err = uploadDashboard(ts.URL, []byte(dashboardPayload), apiKey, appKey)

	if err == nil {
		t.Errorf("uploadDashboard() should  have returned an error")
	}
}

// TestGetDashboardInfo function to test getDashboardInfo()
func TestGetDashboardInfo(t *testing.T) {

	id, url, err := getDashboardInfo(datadogSuccessfullResponse)

	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	if id != "dnq-s5w-h5j" {
		t.Errorf("Did not get expected id, got %s", id)
	}

	if url != "/dashboard/dnq-s5w-h5j/sre---consul-overview" {
		t.Errorf("Did not get expected url, got %s", url)
	}

	id, url, err = getDashboardInfo("badJson")

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
