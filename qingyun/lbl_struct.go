package qingyun

type LoadBalanceListener struct {
	Forwardfor               int          `json:"forwardfor"`
	LoadbalancerListenerID   string       `json:"loadbalancer_listener_id"`
	BalanceMode              string       `json:"balance_mode"`
	ListenerProtocol         string       `json:"listener_protocol"`
	BackendProtocol          string       `json:"backend_protocol"`
	HealthyCheckMethod       string       `json:"healthy_check_method"`
	HealthyCheckOption       string       `json:"healthy_check_option"`
	SessionSticky            string       `json:"session_sticky"`
	LoadbalancerListenerName string       `json:"loadbalancer_listener_name"`
	Backends                 []LblBackend `json:"backends"`
	CreateTime               string       `json:"create_time"`
	LoadbalancerID           string       `json:"loadbalancer_id"`
	ListenerPort             int          `json:"listener_port"`
	ListenerOption           int          `json:"listener_option"`
}

type LblBackend struct {
	LoadbalancerBackendID   string `json:"loadbalancer_backend_id"`
	LoadbalancerBackendName string `json:"loadbalancer_backend_name"`
	Weight                  int    `json:"weight"`
	Port                    int    `json:"port"`
	ResourceID              string `json:"resource_id"`
	LoadbalancerListenerID  string `json:"loadbalancer_listener_id"`
	LoadbalancerID          string `json:"loadbalancer_id"`
	CreateTime              string `json:"create_time"`
}
