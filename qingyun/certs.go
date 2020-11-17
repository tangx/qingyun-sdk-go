package qingyun

type CreateServerCertficateRequest struct {
	ServerCertificateName string `json:"server_certificate_name,omitempty" url:"server_certificate_name,omitempty"`
	CertificateContent    string `json:"certificate_content,omitempty" url:"certificate_content,omitempty"`
	PrivateKey            string `json:"private_key,omitempty" url:"private_key,omitempty"`
}

type CreateServerCertficateResponse struct {
	Action              string `json:"action,omitempty" url:"action,omitempty"`
	ServerCertificateID string `json:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	RetCode             int    `json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

type DeleteServerCertificatesRequest struct {
	ServerCertificates []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
}

type DeleteServerCertificatesResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty" url:"job_id,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

type AssociateCertsToLBRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
}

type AssociateCertsToLBResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	Retcode int    `yaml:"retcode,omitempty" json:"retcode,omitempty" url:"retcode,omitempty"`
}

type DissociateCertsFromLBRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
}

type DissociateCertsFromLBResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}
