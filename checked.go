package datakit

func Enabled(name string) bool {
	if enabled, ok := allInputs[name]; !ok {
		return false // not exist
	} else {
		if EnableUncheckInputs {
			return true
		} else {
			return enabled
		}
	}
}

var (
	EnableUncheckInputs = false

	allInputs = map[string]bool{
		"cshark":      false,
		"httpPacket":  false,
		"mock":        false,
		"secureexec":  false,
		"vsphere":     false,
		"zookeeper":   false,
		"qyt_all":     false,
		"httpProb":    false,
		"activemqlog": false,
		"kafkalog":    false,
		"rabbitmqlog": false,

		"active_directory":       true,
		"activemq":               true,
		"aliyunactiontrail":      true,
		"aliyuncdn":              true,
		"aliyuncms":              true,
		"aliyuncost":             true,
		"aliyunddos":             true,
		"aliyunfc":               true,
		"aliyunlog":              true,
		"aliyunobject":           true,
		"aliyunprice":            true,
		"aliyunrdsslowlog":       true,
		"aliyunsecurity":         true,
		"amqp_consumer":          true,
		"ansible":                true,
		"apache":                 true,
		"aspdotnet":              true,
		"awsbill":                true,
		"awscloudtrail":          true,
		"azure_monitor":          true,
		"baiduIndex":             true,
		"binlog":                 true,
		"bitbucket":              true,
		"cassandra":              true,
		"ceph":                   true,
		"clickhouse":             true,
		"cloudflare":             true,
		"cloudwatch":             true,
		"collectd":               true,
		"confluence":             true,
		"consul":                 true,
		"containerd":             true,
		"coredns":                true,
		"cpu":                    true,
		"csvmetric":              true,
		"csvobject":              true,
		"ddtrace":                true,
		"disk":                   true,
		"diskio":                 true,
		"dns_query":              true,
		"docker":                 true,
		"docker_containers":      true,
		"dockerlog":              true,
		"dotnetclr":              true,
		"druid":                  true,
		"elasticsearch":          true,
		"envoy":                  true,
		"etcd":                   true,
		"exec":                   true,
		"expressjs":              true,
		"external":               true,
		"flink":                  true,
		"fluentd":                true,
		"fluentdlog":             false,
		"file_collector":         false,
		"ginlog":                 true,
		"github":                 true,
		"gitlab":                 true,
		"goruntime":              true,
		"hadoop_hdfs":            true,
		"haproxy":                true,
		"harborMonitor":          true,
		"host_processes":         true,
		"hostobject":             true,
		"http":                   true,
		"http_response":          true,
		"httpjson":               true,
		"httpstat":               true,
		"huaweiyunces":           true,
		"huaweiyunobject":        true,
		"iis":                    true,
		"influxdb":               true,
		"internal":               true,
		"iptables":               true,
		"jboss":                  true,
		"jenkins":                true,
		"jira":                   true,
		"jolokia2_agent":         true,
		"jvm":                    true,
		"k8sobject":              true,
		"kafka":                  true,
		"kafka_consumer":         true,
		"kapacitor":              true,
		"kernel":                 true,
		"kibana":                 true,
		"kong":                   true,
		"kube_inventory":         true,
		"kubernetes":             true,
		"lighttpd":               true,
		"mem":                    true,
		"memcached":              true,
		"modbus":                 true,
		"mongodb":                true,
		"mongodb_oplog":          true,
		"mqtt_consumer":          true,
		"msexchange":             true,
		"mysqlMonitor":           true,
		"mysqlog":                true,
		"nats":                   true,
		"neo4j":                  true,
		"net":                    true,
		"net_response":           true,
		"netstat":                true,
		"nfsstat":                true,
		"nginx":                  true,
		"nginx_plus":             true,
		"nginx_plus_api":         true,
		"nginx_upstream_check":   true,
		"nginx_vts":              true,
		"nginxlog":               true,
		"nsq":                    true,
		"nsq_consumer":           true,
		"ntpq":                   true,
		"nvidia_smi":             true,
		"openldap":               true,
		"openntpd":               true,
		"oraclemonitor":          true,
		"phpfpm":                 true,
		"ping":                   true,
		"postgresql":             true,
		"postgresql_replication": true,
		"processes":              true,
		"procstat":               true,
		"prom":                   true,
		"proxy":                  true,
		"puppetagent":            true,
		"rabbitmq":               true,
		"redis":                  true,
		"redislog":               true,
		"rum":                    true,
		"scanport":               true,
		"self":                   true,
		"smart":                  true,
		"snmp":                   true,
		"socket_listener":        true, // collectd checked
		"solr":                   true,
		"sqlserver":              true,
		"squid":                  true,
		"ssh":                    true,
		"statsd":                 true,
		"swap":                   true,
		"syslog":                 true,
		"system":                 true,
		"systemd":                true,
		"systemd_units":          true,
		"tailf":                  true,
		"telegraf_http":          true,
		"tencentcms":             true,
		"tencentcost":            true,
		"tencentobject":          true,
		"tengine":                true,
		"tidb":                   true,
		"timezone":               true,
		"traceJaeger":            true,
		"traceSkywalking":        true,
		"traceZipkin":            true,
		"tracerouter":            true,
		"traefik":                true,
		"ucloud_monitor":         true,
		"uwsgi":                  true,
		"varnish":                true,
		"weblogic":               true,
		"wechatminiprogram":      true,
		"win_perf_counters":      true,
		"win_services":           true,
		"wmi":                    true,
		"x509_cert":              true,
		"yarn":                   true,
		"zabbix":                 true,
		"zaplog":                 true,
	}
)
