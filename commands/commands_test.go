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

const (
	dasboardID         = "666"
	apiKey             = "1234"
	appKey             = "5678"
	dumpFilePermission = "-rw-------"
	expectedContent    = "aaaa"
	expectedPrettyJSON = `{
	"a": "b"
}`
	datadogSuccessfullGetDashboard = (`{"notify_list":null,"description":"","author_name":"Renaud Hager","template_variables":[{"default":"*","prefix":"hostname","name":"source"},{"default":"dev","prefix":"environment","name":"environment"},{"default":"eu-west-1","prefix":"location","name":"location"},{"default":"*","prefix":"kube_namespace","name":"namespace"},{"default":"kubeworker","prefix":"group_role","name":"group"},{"default":"dev-sre-k8singress-alb","prefix":"name","name":"loadbalancer"}],"is_read_only":false,"id":"hfy-m49-ps3","title":"SRE - Kubernetes","url":"/dashboard/hfy-m49-ps3/sre---kubernetes","created_at":"2019-04-03T18:27:49.044613+00:00","modified_at":"2019-06-25T10:35:08.574408+00:00","author_handle":"renaud.hager@ef.com","widgets":[{"definition":{"autoscale":true,"title":"Pods desired","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.deployment.replicas_desired{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"custom_text","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":5,"x":0,"width":17,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Deployments","font_size":"18","background_color":"gray"},"layout":{"y":0,"x":0,"width":54,"height":5}},{"definition":{"title_size":"16","title":"Pods desired","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.deployment.replicas_desired{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":5,"x":17,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Pods available","title_align":"left","precision":2,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.deployment.replicas_available{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":20,"x":0,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Pods desired","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.deployment.replicas_available{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_width":"normal","palette":"green","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":20,"x":17,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Pods desired","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.daemonset.desired{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"custom_text","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":5,"x":54,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Pods desired","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.daemonset.desired{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":5,"x":71,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Pods ready","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.daemonset.ready{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":20,"x":54,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Pods desired","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.daemonset.ready{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_width":"normal","palette":"green","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":20,"x":71,"width":37,"height":15}},{"definition":{"title_align":"center","title_size":"16","title":"Kubelets Up","tags":["$location","$environment"],"group_by":[],"time":{},"type":"check_status","check":"kubernetes.kubelet.check","grouping":"cluster"},"layout":{"y":0,"x":108,"width":14,"height":20}},{"definition":{"title_align":"center","title_size":"16","title":"Kubelets Ping","tags":["$location","$environment"],"group_by":[],"time":{},"type":"check_status","check":"kubernetes.kubelet.check.ping","grouping":"cluster"},"layout":{"y":0,"x":122,"width":14,"height":20}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"DaemonSets","font_size":"18","background_color":"gray"},"layout":{"y":0,"x":54,"width":54,"height":5}},{"definition":{"autoscale":true,"title":"Pods unavailable","title_align":"left","precision":2,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.deployment.replicas_unavailable{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"palette":"red_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":35,"x":0,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Pods unavailable","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.deployment.replicas_unavailable{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":35,"x":17,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Pods missscheduled","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.daemonset.misscheduled{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"red_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":35,"x":54,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Pods missscheduled","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.daemonset.misscheduled{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":35,"x":71,"width":37,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"ReplicaSets","font_size":"18","background_color":"gray"},"layout":{"y":50,"x":0,"width":54,"height":5}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"StatefullSets","font_size":"18","background_color":"gray"},"layout":{"y":50,"x":54,"width":54,"height":5}},{"definition":{"autoscale":true,"title":"Ready","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"custom_text","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":55,"x":0,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Ready","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":55,"x":17,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Not Ready","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.replicaset.replicas_desired{$source,$environment,$location,$namespace,$group}-sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"red_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":70,"x":0,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Not Ready","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.replicaset.replicas_desired{$source,$environment,$location,$namespace,$group} by {replicaset}-sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group} by {replicaset}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":70,"x":17,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Desired","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"custom_text","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":55,"x":54,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Desired","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":55,"x":71,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Current","title_align":"left","precision":2,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.statefulset.replicas_current{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":70,"x":54,"width":17,"height":15}},{"definition":{"title_size":"16","title":"Current","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.statefulset.replicas_current{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_width":"normal","palette":"green","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":70,"x":71,"width":37,"height":15}},{"definition":{"autoscale":true,"title":"Not Ready","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group}-sum:kubernetes_state.statefulset.replicas_ready{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"custom_fg_color":"#6a53a1","palette":"red_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":84,"x":54,"width":17,"height":16}},{"definition":{"title_size":"16","title":"Not Ready","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group} by {statefulset}-sum:kubernetes_state.statefulset.replicas_ready{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":85,"x":71,"width":37,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Containers status","font_size":"18","background_color":"gray"},"layout":{"y":85,"x":0,"width":54,"height":5}},{"definition":{"title_size":"16","title":"Container status","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes_state.container.running{$source,$environment,$namespace,$location}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"running","expression":"sum:kubernetes_state.container.running{$source,$environment,$namespace,$location}"}]},{"q":"sum:kubernetes_state.container.waiting{$source,$environment,$namespace,$location}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"waiting","expression":"sum:kubernetes_state.container.waiting{$source,$environment,$namespace,$location}"}]},{"q":"sum:kubernetes_state.container.terminated{$source,$environment,$location,$namespace}","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"terminated","expression":"sum:kubernetes_state.container.terminated{$source,$environment,$location,$namespace}"}]}],"type":"timeseries","legend_size":"0"},"layout":{"y":90,"x":0,"width":54,"height":10}},{"definition":{"autoscale":true,"title":"Containers Running","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.container.running{$source,$environment,$location,$namespace}","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":20,"x":108,"width":14,"height":15}},{"definition":{"autoscale":true,"title":"Containers Waiting","title_align":"left","precision":2,"time":{},"title_size":"16","requests":[{"q":"sum:kubernetes_state.container.waiting{$source,$environment,$location,$namespace,$group}","aggregator":"last","conditional_formats":[{"palette":"red_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":20,"x":122,"width":14,"height":15}},{"definition":{"style":{"palette":"yellow_to_green","palette_flip":false},"title_size":"16","group":[],"title":"Running pods / Hosts","title_align":"left","no_metric_hosts":false,"scope":["$environment","$source"],"requests":{"fill":{"q":"sum:kubernetes.pods.running{$environment,$source} by {host}"}},"no_group_hosts":true,"type":"hostmap"},"layout":{"y":35,"x":108,"width":28,"height":22}},{"definition":{"style":{"palette":"YlOrRd","palette_flip":false},"title_size":"16","group":[],"title":"CPU usage / Hosts","title_align":"left","no_metric_hosts":false,"scope":["$environment","$source"],"requests":{"fill":{"q":"sum:kubernetes.cpu.usage.total{$environment,$source} by {host}"}},"no_group_hosts":true,"type":"hostmap"},"layout":{"y":57,"x":108,"width":28,"height":22}},{"definition":{"style":{"palette":"green_to_orange","palette_flip":false},"title_size":"16","group":[],"title":"Memory usage / Hosts","title_align":"left","no_metric_hosts":false,"scope":["$environment","$source"],"requests":{"fill":{"q":"sum:kubernetes.memory.usage{$environment,$source} by {host}"}},"no_group_hosts":true,"type":"hostmap"},"layout":{"y":79,"x":108,"width":28,"height":21}},{"definition":{"title_size":"16","title":"Total CPU by POD","title_align":"left","time":{},"requests":[{"q":"top(avg:kubernetes.cpu.usage.total{$environment,$source,$location,$namespace,!pod_name:no_pod} by {pod_name}, 10, 'mean', 'desc')","conditional_formats":[],"style":{"palette":"warm"}}],"type":"toplist"},"layout":{"y":105,"x":0,"width":54,"height":15}},{"definition":{"title_size":"16","title":"Total Memory by POD","title_align":"left","time":{},"requests":[{"q":"top(avg:kubernetes.memory.usage{$environment,$source,$location,$namespace,!pod_name:no_pod} by {pod_name}, 10, 'mean', 'desc')","conditional_formats":[],"style":{"palette":"purple"}}],"type":"toplist"},"layout":{"y":105,"x":54,"width":54,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Top list","font_size":"18","background_color":"gray"},"layout":{"y":100,"x":0,"width":108,"height":5}},{"definition":{"autoscale":true,"title":"Req/s","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":5,"x":177,"width":13,"height":15}},{"definition":{"title_size":"16","title":"Req/s per backend (priv ingress)","title_align":"left","time":{},"requests":[{"q":"top(sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress} by {backend}.as_rate(), 10, 'mean', 'desc')","conditional_formats":[],"style":{"palette":"cool"}}],"type":"toplist"},"layout":{"y":120,"x":0,"width":54,"height":15}},{"definition":{"title_size":"16","title":"Req/s per backend (pub ingress)","title_align":"left","time":{},"requests":[{"q":"top(sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public} by {backend}.as_rate(), 10, 'mean', 'desc')","conditional_formats":[],"style":{"palette":"cool"}}],"type":"toplist"},"layout":{"y":120,"x":54,"width":54,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Private Ingress","font_size":"18","background_color":"gray"},"layout":{"y":0,"x":136,"width":54,"height":5}},{"definition":{"title_size":"16","title":"Req/s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","style":{"line_width":"normal","palette":"green","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":5,"x":136,"width":41,"height":15}},{"definition":{"title_size":"16","title":"% 5XX/s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"((sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:500}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:502}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:503}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:504}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:501}.as_rate())*100)/sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","style":{"line_width":"normal","palette":"red","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":20,"x":136,"width":41,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Public Ingress","font_size":"18","background_color":"gray"},"layout":{"y":35,"x":136,"width":54,"height":5}},{"definition":{"title_size":"16","title":"Req/s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public}.as_rate()","style":{"line_width":"normal","palette":"green","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":40,"x":136,"width":41,"height":15}},{"definition":{"title_size":"16","title":"5XX/s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public,code:500,code:501,code:502,code:503,code:504}.as_rate()","style":{"line_width":"normal","palette":"red","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":55,"x":136,"width":41,"height":15}},{"definition":{"autoscale":true,"title":"% 5XX/s","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"((sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:500}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:502}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:503}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:504}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:501}.as_rate())*100)/sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":"<="},{"palette":"red_on_white","value":1,"comparator":">="},{"palette":"yellow_on_white","value":0.5,"comparator":">="},{"palette":"yellow_on_white","value":1,"comparator":"<"}]}],"type":"query_value"},"layout":{"y":20,"x":177,"width":13,"height":15}},{"definition":{"autoscale":true,"title":"Req/s","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public}.as_rate()","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":40,"x":177,"width":13,"height":15}},{"definition":{"autoscale":true,"title":"5XX/s","title_align":"left","precision":0,"time":{},"title_size":"16","requests":[{"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public,code:500,code:501,code:502,code:503,code:504}.as_rate()","aggregator":"last","conditional_formats":[{"palette":"green_on_white","value":0,"comparator":">"}]}],"type":"query_value"},"layout":{"y":55,"x":177,"width":13,"height":15}},{"definition":{"style":{"palette":"green_to_orange","palette_flip":false},"title_size":"16","group":[],"title":"Req/s per private ingress","title_align":"left","node_type":"host","no_metric_hosts":false,"scope":["$environment","group_role:kubeingress"],"requests":{"fill":{"q":"sum:traefik.backend.request.total{$environment,group_role:kubeingress} by {host}"}},"no_group_hosts":true,"type":"hostmap"},"layout":{"y":100,"x":108,"width":28,"height":20}},{"definition":{"style":{"palette":"green_to_orange","palette_flip":false},"title_size":"16","group":[],"title":"Req/s per public ingress","title_align":"left","node_type":"host","no_metric_hosts":false,"scope":["$environment","group_role:kubeingress-public"],"requests":{"fill":{"q":"sum:traefik.backend.request.total{$environment,group_role:kubeingress-public} by {host}"}},"no_group_hosts":true,"type":"hostmap"},"layout":{"y":117,"x":108,"width":28,"height":18}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"CoreDns","font_size":"18","background_color":"gray"},"layout":{"y":122,"x":136,"width":54,"height":5}},{"definition":{"title_size":"16","title":"Req/s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:coredns.request_count{$environment,$location}.as_rate()","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","legend_size":"0"},"layout":{"y":127,"x":136,"width":54,"height":15}},{"definition":{"title_size":"16","title":"Req/s per type","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:coredns.request_type_count{$environment,$location} by {type}.as_rate()","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"area"}],"type":"timeseries","legend_size":"0"},"layout":{"y":142,"x":136,"width":54,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Resources Overview","font_size":"18","background_color":"gray"},"layout":{"y":135,"x":0,"width":108,"height":5}},{"definition":{"title_size":"16","title":"CPU","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes.cpu.capacity{$source,$environment,$location,$group}","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"capacity","expression":"sum:kubernetes.cpu.capacity{$source,$environment,$location,$group}"}]},{"q":"sum:kubernetes.cpu.requests{$source,$environment,$location,$group}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"requested","expression":"sum:kubernetes.cpu.requests{$source,$environment,$location,$group}"}]}],"type":"timeseries","legend_size":"0"},"layout":{"y":140,"x":0,"width":54,"height":15}},{"definition":{"title_size":"16","title":"Memory","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:kubernetes.memory.capacity{$source,$environment,$location,$group}","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"capacity","expression":"sum:kubernetes.memory.capacity{$source,$environment,$location,$group}"}]},{"q":"sum:kubernetes.memory.requests{$source,$environment,$location,$group}, sum:kubernetes.memory.usage{$source,$environment,$location,$group}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line","metadata":[{"alias_name":"requested","expression":"sum:kubernetes.memory.requests{$source,$environment,$location,$group}"},{"alias_name":"usage","expression":"sum:kubernetes.memory.usage{$source,$environment,$location,$group}"}]}],"type":"timeseries","legend_size":"0"},"layout":{"y":140,"x":54,"width":54,"height":15}},{"definition":{"tick_pos":"50%","show_tick":true,"type":"note","tick_edge":"left","text_align":"center","content":"Load Balancer Overview","font_size":"18","background_color":"gray"},"layout":{"y":70,"x":136,"width":54,"height":5}},{"definition":{"title_size":"16","title":"5XX /s","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:aws.applicationelb.httpcode_elb_5xx{$environment,$loadbalancer}.as_rate()","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line"}],"type":"timeseries","legend_size":"0"},"layout":{"y":75,"x":136,"width":54,"height":14}},{"definition":{"title_size":"16","title":"RPS","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:aws.applicationelb.request_count{$environment,$loadbalancer}.as_rate()","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","legend_size":"0"},"layout":{"y":89,"x":136,"width":54,"height":16}},{"definition":{"title_size":"16","title":"Latency","title_align":"left","show_legend":false,"time":{},"requests":[{"q":"sum:aws.applicationelb.target_response_time.p95{$environment,$loadbalancer}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","legend_size":"0"},"layout":{"y":105,"x":136,"width":54,"height":17}}],"layout_type":"free"}`)
	expectedPayloadGetDashboard    = (`{"author_handle":"renaud.hager@ef.com","created_at":"2019-04-03T18:27:49.044613+00:00","description":"","id":"hfy-m49-ps3","is_read_only":false,"layout_type":"free","modified_at":"2019-06-25T10:35:08.574408+00:00","notify_list":null,"template_variables":[{"default":"*","name":"source","prefix":"hostname"},{"default":"dev","name":"environment","prefix":"environment"},{"default":"eu-west-1","name":"location","prefix":"location"},{"default":"*","name":"namespace","prefix":"kube_namespace"},{"default":"kubeworker","name":"group","prefix":"group_role"},{"default":"dev-sre-k8singress-alb","name":"loadbalancer","prefix":"name"}],"title":"SRE - Kubernetes","url":"/dashboard/hfy-m49-ps3/sre---kubernetes","widgets":[{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"custom_text","value":0}],"q":"sum:kubernetes_state.deployment.replicas_desired{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":0,"y":5}},{"definition":{"background_color":"gray","content":"Deployments","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":0,"y":0}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.deployment.replicas_desired{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_type":"solid","line_width":"normal","palette":"purple"}}],"show_legend":false,"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":17,"y":5}},{"definition":{"autoscale":true,"precision":2,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:kubernetes_state.deployment.replicas_available{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods available","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":0,"y":20}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.deployment.replicas_available{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_type":"solid","line_width":"normal","palette":"green"}}],"show_legend":false,"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":17,"y":20}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"custom_text","value":0}],"q":"sum:kubernetes_state.daemonset.desired{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":54,"y":5}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.daemonset.desired{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_type":"solid","line_width":"normal","palette":"purple"}}],"show_legend":false,"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":5}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"green_on_white","value":0}],"q":"sum:kubernetes_state.daemonset.ready{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods ready","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":54,"y":20}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.daemonset.ready{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_type":"solid","line_width":"normal","palette":"green"}}],"show_legend":false,"time":{},"title":"Pods desired","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":20}},{"definition":{"check":"kubernetes.kubelet.check","group_by":[],"grouping":"cluster","tags":["$location","$environment"],"time":{},"title":"Kubelets Up","title_align":"center","title_size":"16","type":"check_status"},"layout":{"height":20,"width":14,"x":108,"y":0}},{"definition":{"check":"kubernetes.kubelet.check.ping","group_by":[],"grouping":"cluster","tags":["$location","$environment"],"time":{},"title":"Kubelets Ping","title_align":"center","title_size":"16","type":"check_status"},"layout":{"height":20,"width":14,"x":122,"y":0}},{"definition":{"background_color":"gray","content":"DaemonSets","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":54,"y":0}},{"definition":{"autoscale":true,"precision":2,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"red_on_white","value":0}],"q":"sum:kubernetes_state.deployment.replicas_unavailable{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods unavailable","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":0,"y":35}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.deployment.replicas_unavailable{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"Pods unavailable","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":17,"y":35}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"red_on_white","value":0}],"q":"sum:kubernetes_state.daemonset.misscheduled{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Pods missscheduled","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":54,"y":35}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.daemonset.misscheduled{$source,$environment,$location,$namespace,$group} by {daemonset}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"Pods missscheduled","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":35}},{"definition":{"background_color":"gray","content":"ReplicaSets","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":0,"y":50}},{"definition":{"background_color":"gray","content":"StatefullSets","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":54,"y":50}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"custom_text","value":0}],"q":"sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Ready","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":0,"y":55}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group} by {deployment}","style":{"line_type":"solid","line_width":"normal","palette":"purple"}}],"show_legend":false,"time":{},"title":"Ready","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":17,"y":55}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"red_on_white","value":0}],"q":"sum:kubernetes_state.replicaset.replicas_desired{$source,$environment,$location,$namespace,$group}-sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Not Ready","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":0,"y":70}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.replicaset.replicas_desired{$source,$environment,$location,$namespace,$group} by {replicaset}-sum:kubernetes_state.replicaset.replicas_ready{$source,$environment,$location,$namespace,$group} by {replicaset}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"Not Ready","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":17,"y":70}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"custom_text","value":0}],"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Desired","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":54,"y":55}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_type":"solid","line_width":"normal","palette":"purple"}}],"show_legend":false,"time":{},"title":"Desired","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":55}},{"definition":{"autoscale":true,"precision":2,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:kubernetes_state.statefulset.replicas_current{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Current","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":17,"x":54,"y":70}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.statefulset.replicas_current{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_type":"solid","line_width":"normal","palette":"green"}}],"show_legend":false,"time":{},"title":"Current","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":70}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","custom_fg_color":"#6a53a1","palette":"red_on_white","value":0}],"q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group}-sum:kubernetes_state.statefulset.replicas_ready{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Not Ready","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":16,"width":17,"x":54,"y":84}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:kubernetes_state.statefulset.replicas_desired{$source,$environment,$location,$namespace,$group} by {statefulset}-sum:kubernetes_state.statefulset.replicas_ready{$source,$environment,$location,$namespace,$group} by {statefulset}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"Not Ready","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":37,"x":71,"y":85}},{"definition":{"background_color":"gray","content":"Containers status","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":0,"y":85}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","metadata":[{"alias_name":"running","expression":"sum:kubernetes_state.container.running{$source,$environment,$namespace,$location}"}],"q":"sum:kubernetes_state.container.running{$source,$environment,$namespace,$location}","style":{"line_type":"solid","line_width":"normal","palette":"dog_classic"}},{"display_type":"line","metadata":[{"alias_name":"waiting","expression":"sum:kubernetes_state.container.waiting{$source,$environment,$namespace,$location}"}],"q":"sum:kubernetes_state.container.waiting{$source,$environment,$namespace,$location}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}},{"display_type":"line","metadata":[{"alias_name":"terminated","expression":"sum:kubernetes_state.container.terminated{$source,$environment,$location,$namespace}"}],"q":"sum:kubernetes_state.container.terminated{$source,$environment,$location,$namespace}","style":{"line_type":"solid","line_width":"normal","palette":"grey"}}],"show_legend":false,"time":{},"title":"Container status","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":10,"width":54,"x":0,"y":90}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:kubernetes_state.container.running{$source,$environment,$location,$namespace}"}],"time":{},"title":"Containers Running","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":14,"x":108,"y":20}},{"definition":{"autoscale":true,"precision":2,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"red_on_white","value":0}],"q":"sum:kubernetes_state.container.waiting{$source,$environment,$location,$namespace,$group}"}],"time":{},"title":"Containers Waiting","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":14,"x":122,"y":20}},{"definition":{"group":[],"no_group_hosts":true,"no_metric_hosts":false,"requests":{"fill":{"q":"sum:kubernetes.pods.running{$environment,$source} by {host}"}},"scope":["$environment","$source"],"style":{"palette":"yellow_to_green","palette_flip":false},"title":"Running pods / Hosts","title_align":"left","title_size":"16","type":"hostmap"},"layout":{"height":22,"width":28,"x":108,"y":35}},{"definition":{"group":[],"no_group_hosts":true,"no_metric_hosts":false,"requests":{"fill":{"q":"sum:kubernetes.cpu.usage.total{$environment,$source} by {host}"}},"scope":["$environment","$source"],"style":{"palette":"YlOrRd","palette_flip":false},"title":"CPU usage / Hosts","title_align":"left","title_size":"16","type":"hostmap"},"layout":{"height":22,"width":28,"x":108,"y":57}},{"definition":{"group":[],"no_group_hosts":true,"no_metric_hosts":false,"requests":{"fill":{"q":"sum:kubernetes.memory.usage{$environment,$source} by {host}"}},"scope":["$environment","$source"],"style":{"palette":"green_to_orange","palette_flip":false},"title":"Memory usage / Hosts","title_align":"left","title_size":"16","type":"hostmap"},"layout":{"height":21,"width":28,"x":108,"y":79}},{"definition":{"requests":[{"conditional_formats":[],"q":"top(avg:kubernetes.cpu.usage.total{$environment,$source,$location,$namespace,!pod_name:no_pod} by {pod_name}, 10, 'mean', 'desc')","style":{"palette":"warm"}}],"time":{},"title":"Total CPU by POD","title_align":"left","title_size":"16","type":"toplist"},"layout":{"height":15,"width":54,"x":0,"y":105}},{"definition":{"requests":[{"conditional_formats":[],"q":"top(avg:kubernetes.memory.usage{$environment,$source,$location,$namespace,!pod_name:no_pod} by {pod_name}, 10, 'mean', 'desc')","style":{"palette":"purple"}}],"time":{},"title":"Total Memory by POD","title_align":"left","title_size":"16","type":"toplist"},"layout":{"height":15,"width":54,"x":54,"y":105}},{"definition":{"background_color":"gray","content":"Top list","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":108,"x":0,"y":100}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()"}],"time":{},"title":"Req/s","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":13,"x":177,"y":5}},{"definition":{"requests":[{"conditional_formats":[],"q":"top(sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress} by {backend}.as_rate(), 10, 'mean', 'desc')","style":{"palette":"cool"}}],"time":{},"title":"Req/s per backend (priv ingress)","title_align":"left","title_size":"16","type":"toplist"},"layout":{"height":15,"width":54,"x":0,"y":120}},{"definition":{"requests":[{"conditional_formats":[],"q":"top(sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public} by {backend}.as_rate(), 10, 'mean', 'desc')","style":{"palette":"cool"}}],"time":{},"title":"Req/s per backend (pub ingress)","title_align":"left","title_size":"16","type":"toplist"},"layout":{"height":15,"width":54,"x":54,"y":120}},{"definition":{"background_color":"gray","content":"Private Ingress","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":136,"y":0}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"green"}}],"show_legend":false,"time":{},"title":"Req/s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":41,"x":136,"y":5}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"((sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:500}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:502}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:503}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:504}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:501}.as_rate())*100)/sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"red"}}],"show_legend":false,"time":{},"title":"% 5XX/s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":41,"x":136,"y":20}},{"definition":{"background_color":"gray","content":"Public Ingress","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":136,"y":35}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"green"}}],"show_legend":false,"time":{},"title":"Req/s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":41,"x":136,"y":40}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public,code:500,code:501,code:502,code:503,code:504}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"red"}}],"show_legend":false,"time":{},"title":"5XX/s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":41,"x":136,"y":55}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003c=","palette":"green_on_white","value":0},{"comparator":"\u003e=","palette":"red_on_white","value":1},{"comparator":"\u003e=","palette":"yellow_on_white","value":0.5},{"comparator":"\u003c","palette":"yellow_on_white","value":1}],"q":"((sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:500}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:502}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:503}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:504}.as_rate()+sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress,code:501}.as_rate())*100)/sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress}.as_rate()"}],"time":{},"title":"% 5XX/s","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":13,"x":177,"y":20}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public}.as_rate()"}],"time":{},"title":"Req/s","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":13,"x":177,"y":40}},{"definition":{"autoscale":true,"precision":0,"requests":[{"aggregator":"last","conditional_formats":[{"comparator":"\u003e","palette":"green_on_white","value":0}],"q":"sum:traefik.backend.request.total{$environment,$location,group_role:kubeingress-public,code:500,code:501,code:502,code:503,code:504}.as_rate()"}],"time":{},"title":"5XX/s","title_align":"left","title_size":"16","type":"query_value"},"layout":{"height":15,"width":13,"x":177,"y":55}},{"definition":{"group":[],"no_group_hosts":true,"no_metric_hosts":false,"node_type":"host","requests":{"fill":{"q":"sum:traefik.backend.request.total{$environment,group_role:kubeingress} by {host}"}},"scope":["$environment","group_role:kubeingress"],"style":{"palette":"green_to_orange","palette_flip":false},"title":"Req/s per private ingress","title_align":"left","title_size":"16","type":"hostmap"},"layout":{"height":20,"width":28,"x":108,"y":100}},{"definition":{"group":[],"no_group_hosts":true,"no_metric_hosts":false,"node_type":"host","requests":{"fill":{"q":"sum:traefik.backend.request.total{$environment,group_role:kubeingress-public} by {host}"}},"scope":["$environment","group_role:kubeingress-public"],"style":{"palette":"green_to_orange","palette_flip":false},"title":"Req/s per public ingress","title_align":"left","title_size":"16","type":"hostmap"},"layout":{"height":18,"width":28,"x":108,"y":117}},{"definition":{"background_color":"gray","content":"CoreDns","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":136,"y":122}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","q":"sum:coredns.request_count{$environment,$location}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"dog_classic"}}],"show_legend":false,"time":{},"title":"Req/s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":54,"x":136,"y":127}},{"definition":{"legend_size":"0","requests":[{"display_type":"area","q":"sum:coredns.request_type_count{$environment,$location} by {type}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"dog_classic"}}],"show_legend":false,"time":{},"title":"Req/s per type","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":54,"x":136,"y":142}},{"definition":{"background_color":"gray","content":"Resources Overview","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":108,"x":0,"y":135}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","metadata":[{"alias_name":"capacity","expression":"sum:kubernetes.cpu.capacity{$source,$environment,$location,$group}"}],"q":"sum:kubernetes.cpu.capacity{$source,$environment,$location,$group}","style":{"line_type":"solid","line_width":"normal","palette":"grey"}},{"display_type":"line","metadata":[{"alias_name":"requested","expression":"sum:kubernetes.cpu.requests{$source,$environment,$location,$group}"}],"q":"sum:kubernetes.cpu.requests{$source,$environment,$location,$group}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"CPU","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":54,"x":0,"y":140}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","metadata":[{"alias_name":"capacity","expression":"sum:kubernetes.memory.capacity{$source,$environment,$location,$group}"}],"q":"sum:kubernetes.memory.capacity{$source,$environment,$location,$group}","style":{"line_type":"solid","line_width":"normal","palette":"grey"}},{"display_type":"line","metadata":[{"alias_name":"requested","expression":"sum:kubernetes.memory.requests{$source,$environment,$location,$group}"},{"alias_name":"usage","expression":"sum:kubernetes.memory.usage{$source,$environment,$location,$group}"}],"q":"sum:kubernetes.memory.requests{$source,$environment,$location,$group}, sum:kubernetes.memory.usage{$source,$environment,$location,$group}","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"Memory","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":15,"width":54,"x":54,"y":140}},{"definition":{"background_color":"gray","content":"Load Balancer Overview","font_size":"18","show_tick":true,"text_align":"center","tick_edge":"left","tick_pos":"50%","type":"note"},"layout":{"height":5,"width":54,"x":136,"y":70}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","q":"sum:aws.applicationelb.httpcode_elb_5xx{$environment,$loadbalancer}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"warm"}}],"show_legend":false,"time":{},"title":"5XX /s","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":14,"width":54,"x":136,"y":75}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","q":"sum:aws.applicationelb.request_count{$environment,$loadbalancer}.as_rate()","style":{"line_type":"solid","line_width":"normal","palette":"dog_classic"}}],"show_legend":false,"time":{},"title":"RPS","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":16,"width":54,"x":136,"y":89}},{"definition":{"legend_size":"0","requests":[{"display_type":"line","q":"sum:aws.applicationelb.target_response_time.p95{$environment,$loadbalancer}","style":{"line_type":"solid","line_width":"normal","palette":"dog_classic"}}],"show_legend":false,"time":{},"title":"Latency","title_align":"left","title_size":"16","type":"timeseries"},"layout":{"height":17,"width":54,"x":136,"y":105}}]}`)
	datadogSuccessfullResponse     = (`{"notify_list":[],"description":"created by renaud.hager@nospam.com","author_name":"Renaud Hager","template_variables":[{"default":"eu-west-1","prefix":"location","name":"location"},{"default":"*","prefix":"environment","name":"environment"},{"default":"sre","prefix":"team","name":"team"}],"is_read_only":false,"id":"dnq-s5w-h5j","title":"SRE - Consul Overview","url":"/dashboard/dnq-s5w-h5j/sre---consul-overview","created_at":"2019-06-14T14:58:34.760504+00:00","modified_at":"2019-06-14T14:58:34.760504+00:00","author_handle":"renaud.hager@nospam.com","widgets":[{"definition":{"widgets":[{"definition":{"type":"query_value","requests":[{"q":"avg:consul.autopilot.healthy{$location,$team,$environment,group_role:consul-server}","aggregator":"avg","conditional_formats":[{"palette":"white_on_red","comparator":"<=","value":0.5},{"palette":"white_on_green","comparator":">=","value":0.9},{"palette":"white_on_yellow","comparator":"<=","value":0.89}]}],"autoscale":false,"precision":0,"title":"Overall health"},"id":118514630952118},{"definition":{"requests":[{"q":"avg:consul.kvs.apply.avg{$location,$environment,$team,group_role:consul-server}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"area"}],"type":"timeseries","title":"KV and transaction apply latency"},"id":8450620905653722},{"definition":{"requests":[{"q":"max:consul.raft.commitTime.avg{$location,$team,$environment,group_role:consul-server}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","title":"Max Raft commit time (ms)"},"id":7470639491355588},{"definition":{"requests":[{"q":"avg:consul.kvs.apply.max{$location,$team,$environment}, avg:consul.txn.apply{$location,$team,$environment}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"bars"}],"type":"timeseries","title":"Various latency related to Raft (ms)"},"id":6899892404741622},{"definition":{"requests":[{"q":"avg:consul.raft.leader.lastContact.max{*}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line"}],"type":"timeseries","title":"Last contact from leader (ms)"},"id":3355181384970402}],"layout_type":"ordered","type":"group","title":"Latency"},"id":5254103359870646},{"definition":{"widgets":[{"definition":{"requests":[{"q":"max:consul.catalog.total_nodes{$location,$environment,$team}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"max:consul.catalog.total_nodes{$location,$environment,$team}","alias_name":"Total nodes"}]},{"q":"avg:consul.peers{$location,$environment,$team}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:consul.peers{$location,$environment,$team}","alias_name":"Peer nodes"}]}],"type":"timeseries","title":"Consul nodes"},"id":77898},{"definition":{"requests":[{"q":"avg:consul.rpc.query{$location,$environment,$team,group_role:consul-server}.as_count(), avg:consul.rpc.request{$location,$environment,$team,group_role:consul-server}.as_count()","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"bars","metadata":[{"expression":"avg:consul.rpc.query{$location,$environment,$team,group_role:consul-server}.as_count()","alias_name":"rpc query"},{"expression":"avg:consul.rpc.request{$location,$environment,$team,group_role:consul-server}.as_count()","alias_name":"rpc request"}]}],"type":"timeseries","title":"RPC request/query"},"id":4566201650946910},{"definition":{"requests":[{"q":"sum:consul.serf.member.join{$location,$team,$environment,group_role:consul-server}.as_count(), sum:consul.serf.member.left{$location,$team,$environment,group_role:consul-server}.as_count(), sum:consul.serf.member.failed{$location,$team,$environment,group_role:consul-server}.as_count()","style":{"line_width":"normal","palette":"orange","line_type":"solid"},"display_type":"bars","metadata":[{"expression":"sum:consul.serf.member.join{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf join"},{"expression":"sum:consul.serf.member.failed{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf failed"},{"expression":"sum:consul.serf.member.left{$location,$team,$environment,group_role:consul-server}.as_count()","alias_name":"serf left"}]}],"type":"timeseries","title":"Serf activity"},"id":3286132655611000}],"layout_type":"ordered","type":"group","title":"Network and Serf"},"id":286350767033660},{"definition":{"widgets":[{"definition":{"requests":[{"q":"max:consul.runtime.total_gc_pause_ns{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"area"}],"type":"timeseries","title":"Max GC time (ns)"},"id":6725921563822062},{"definition":{"requests":[{"q":"avg:consul.runtime.alloc_bytes{$location,$team,$environment,group_role:consul-server} by {host}, avg:consul.runtime.sys_bytes{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"purple","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:consul.runtime.alloc_bytes{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"alloc bytes"},{"expression":"avg:consul.runtime.sys_bytes{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"sys bytes"}]}],"type":"timeseries","title":"Memory usage"},"id":497692778043364}],"layout_type":"ordered","type":"group","title":"Memory"},"id":5524458375359210},{"definition":{"widgets":[{"definition":{"requests":[{"q":"avg:system.cpu.user{$location,$environment,$team,group_role:consul-server} by {host}+avg:system.cpu.system{$location,$environment,$team,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"area","metadata":[{"expression":"avg:system.cpu.user{$location,$environment,$team,group_role:consul-server} by {host}+avg:system.cpu.system{$location,$environment,$team,group_role:consul-server} by {host}","alias_name":"system+ users"}]}],"type":"timeseries","title":"CPU Usage"},"id":77901},{"definition":{"requests":[{"q":"(100*avg:system.disk.used{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs} by {host,device})/avg:system.disk.total{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs,!device:overlay} by {host,device}","style":{"line_width":"normal","palette":"dog_classic","line_type":"solid"},"display_type":"line","metadata":[{"expression":"(100*avg:system.disk.used{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs} by {host,device})/avg:system.disk.total{$location,$environment,$team,group_role:consul-server,!device:shm,!device:tmpfs,!device:overlay} by {host,device}","alias_name":"% used"}]}],"yaxis":{"max":"100","min":"0"},"type":"timeseries","title":"Disk usage"},"id":5907332623733836},{"definition":{"requests":[{"q":"avg:system.mem.usable{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"warm","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:system.mem.usable{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"usable"}]},{"q":"avg:system.mem.total{$location,$team,$environment,group_role:consul-server} by {host}","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"line","metadata":[{"expression":"avg:system.mem.total{$location,$team,$environment,group_role:consul-server} by {host}","alias_name":"total"}]}],"type":"timeseries","title":"Memory Usable"},"id":7740976102574542},{"definition":{"requests":[{"q":"per_second(avg:system.net.bytes_sent{$location,$environment,$team,group_role:consul-server} by {host})","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"area","metadata":[{"expression":"per_second(avg:system.net.bytes_sent{$location,$environment,$team,group_role:consul-server} by {host})","alias_name":"bytes_sent"}]}],"type":"timeseries","title":"Bytes sent/s"},"id":6219502221062534},{"definition":{"requests":[{"q":"per_second(avg:system.net.bytes_rcvd{$location,$environment,$team,group_role:consul-server} by {host})","style":{"line_width":"normal","palette":"grey","line_type":"solid"},"display_type":"area","metadata":[{"expression":"per_second(avg:system.net.bytes_rcvd{$location,$environment,$team,group_role:consul-server} by {host})","alias_name":"bytes_rcvd"}]}],"type":"timeseries","title":"Bytes rcvd/s"},"id":3344996939592842}],"layout_type":"ordered","type":"group","title":"System"},"id":140191778868366}],"layout_type":"ordered"}`)
	dashboardPayload               = (`{
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
	datadogSuccessfullDashboardListResponse = (`{"dashboards":[{"created_at":"2019-04-03T18:27:49.044613+00:00","author_handle":"renaud.hager@spam.com","is_read_only":false,"description":"","title":"SRE - Kubernetes","url":"/dashboard/hfy-m49-ps3/sre---kubernetes","layout_type":"free","modified_at":"2019-06-25T10:35:08.574408+00:00","id":"hfy-m49-ps3"},{"created_at":"2019-02-27T12:03:28.403848+00:00","author_handle":"renaud.hager@spam.com","is_read_only":false,"description":"created by renaud.hager@spam.com (cloned)","title":"SRE - Vault Overview","url":"/dashboard/hyv-dzi-xas/sre---vault-overview","layout_type":"ordered","modified_at":"2019-06-19T23:39:20.305904+00:00","id":"hyv-dzi-xas"}]}`)
	expectedJSONOutput                      = `{"dashboards":[{"created_at":"2019-04-03T18:27:49.044613+00:00","is_read_only":false,"description":"","id":"hfy-m49-ps3","title":"SRE - Kubernetes","url":"/dashboard/hfy-m49-ps3/sre---kubernetes","layout_type":"free","modified_at":"2019-06-25T10:35:08.574408+00:00","author_handle":"renaud.hager@spam.com"},{"created_at":"2019-02-27T12:03:28.403848+00:00","is_read_only":false,"description":"created by renaud.hager@spam.com (cloned)","id":"hyv-dzi-xas","title":"SRE - Vault Overview","url":"/dashboard/hyv-dzi-xas/sre---vault-overview","layout_type":"ordered","modified_at":"2019-06-19T23:39:20.305904+00:00","author_handle":"renaud.hager@spam.com"}]}
`
	expectedTextOutput = `SRE - Kubernetes | hfy-m49-ps3
SRE - Vault Overview | hyv-dzi-xas
`
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

	payload, _, err := getDashboard(ts.URL, dasboardID, apiKey, appKey)

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

// TestStripBadField function that test stripBadField()
func TestStripBadField(t *testing.T) {
	input := `{"a":"b","c":"d","e":"f"}`
	expectedOutput := `{"a":"b","e":"f"}`

	output, err := stripBadField([]byte(input), "c")

	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	if string(output) != expectedOutput {
		t.Errorf("output should be `%v`, got `%v`", expectedOutput, string(output))
	}

	_, err = stripBadField([]byte("foo"), "c")

	if err == nil {
		t.Errorf("err should not be nil")
	}
}

// TestDeleteDashboard test function for deleteDashboard()
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

	err := deleteDashboard(ts.URL, dasboardID, apiKey, appKey)

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

	err := deleteDashboard(ts.URL, dasboardID, apiKey, appKey)

	if err == nil {
		t.Errorf("deleteDashboard() should have returned an error")
	}
}

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

	dashboardList, err := getDashboardList(ts.URL, apiKey, appKey)

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

	_, err := getDashboardList(ts.URL, apiKey, appKey)

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

	err := output(dashboardList, "text", false)

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

// TestOutput test function for output() with text format
func TestOutputJsonFormat(t *testing.T) {
	var dashboardList DashboardList

	_ = json.Unmarshal([]byte(datadogSuccessfullDashboardListResponse), &dashboardList)

	previousStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := output(dashboardList, "json", false)

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

	err := list(context)

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

	err := list(context)

	if err == nil {
		t.Errorf("list() should have returned an error")
	}
}

// TestDeleteOK test for delete() with proper config
func TestDeleteOK(t *testing.T) {
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

	set := flag.NewFlagSet("test", 0)
	set.String("dh", ts.URL, "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")
	set.String("id", dasboardID, "doc")

	context := cli.NewContext(nil, set, nil)

	err := delete(context)

	if err != nil {
		t.Errorf("list() should not have returned an error")
	}
}

// TestDeleteKO test for delete() with wrong config
func TestDeleteKO(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	set.String("dh", "wronghost", "doc")
	set.String("api-key", apiKey, "doc")
	set.String("app-key", appKey, "doc")

	context := cli.NewContext(nil, set, nil)

	err := delete(context)

	if err == nil {
		t.Errorf("delete() should have returned an error")
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

	err := pull(context)

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

	err := pull(context)

	if err == nil {
		t.Errorf("pull() should have returned an error")
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

	err := push(context)

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

	err := push(context)

	if err == nil {
		t.Errorf("push() should have returned an error")
	}
}
