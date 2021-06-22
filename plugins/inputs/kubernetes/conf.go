package kubernetes

const (
	configSample = `
[[inputs.kubernetes]]
    # required
    interval = "10s"
    ## URL for the Kubernetes API
    url = "https://kubernetes.default:443"
    ## Use bearer token for authorization. ('bearer_token' takes priority)
    ## at: /run/secrets/kubernetes.io/serviceaccount/token
    bearer_token = "/run/secrets/kubernetes.io/serviceaccount/token"

    ## Set http timeout (default 5 seconds)
    timeout = "5s"

     ## Optional TLS Config
    tls_ca = "/run/secrets/kubernetes.io/serviceaccount/ca.crt"

    ## Use TLS but skip chain & host verification
    insecure_skip_verify = false

    [inputs.kubernetes.tags]
    # tag1 = "val1"
    # tag2 = "val2"
`
)