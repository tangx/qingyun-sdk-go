package qingyun

import "time"

type CreateServerCertificateRequest struct {
	ServerCertificateName string `json:"server_certificate_name,omitempty" url:"server_certificate_name,omitempty"`
	CertificateContent    string `json:"certificate_content,omitempty" url:"certificate_content,omitempty"`
	PrivateKey            string `json:"private_key,omitempty" url:"private_key,omitempty"`
}

type CreateServerCertificateResponse struct {
	Action              string `json:"action,omitempty" url:"action,omitempty"`
	ServerCertificateID string `json:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	RetCode             int    `json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) CreateCertificate(params CreateServerCertificateRequest) (resp CreateServerCertificateResponse, err error) {
	err = cli.MethodGET("CreateServerCertificate", params, &resp)
	return
}

type DeleteServerCertificatesRequest struct {
	ServerCertificates []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
}

type DeleteServerCertificatesResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty" url:"job_id,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) DeleteCertificates(params DeleteServerCertificatesRequest) (resp DeleteServerCertificatesResponse, err error) {
	err = cli.MethodGET("DeleteServerCertificates", params, &resp)
	return
}

type AssociateCertsToLBListenerRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
}

type AssociateCertsToLBListenerResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	Retcode int    `yaml:"retcode,omitempty" json:"retcode,omitempty" url:"retcode,omitempty"`
}

func (cli *Client) AssociateCertsToLBListener(params AssociateCertsToLBListenerRequest) (resp AssociateCertsToLBListenerResponse, err error) {
	err = cli.MethodGET("AssociateServerCertsToLBListener", params, &resp)
	return
}

type DissociateCertsFromLBListenerRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
}

type DissociateCertsFromLBListenerResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) DissociateCertsFromLBListener(params DissociateCertsFromLBListenerRequest) (resp DissociateCertsFromLBListenerResponse, err error) {
	err = cli.MethodGET("DissociateServerCertsFromLBListener", params, &resp)
	return
}

type DescribeCertsRequest struct {
	ServerCertificates []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty"`
	SearchWord         string   `yaml:"search_word,omitempty" json:"search_word,omitempty" url:"search_word,omitempty"`
	Verbose            int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset             int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit              int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
}

type DescribeCertsResponse struct {
	Action               string              `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	ServerCertificateSet []ServerCertificate `yaml:"server_certificate_set,omitempty" json:"server_certificate_set,omitempty" url:"server_certificate_set,omitempty"`
	TotalCount           int                 `yaml:"total_count,omitempty" json:"total_count,omitempty" url:"total_count,omitempty"`
}

type ServerCertificate struct {
	ServerCertificateID   string    `yaml:"server_certificate_id,omitempty" json:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	ServerCertificateName string    `yaml:"server_certificate_name,omitempty" json:"server_certificate_name,omitempty" url:"server_certificate_name,omitempty"`
	PrivateKey            string    `yaml:"private_key,omitempty" json:"private_key,omitempty" url:"private_key,omitempty"`
	CertificateContent    string    `yaml:"certificate_content,omitempty" json:"certificate_content,omitempty" url:"certificate_content,omitempty"`
	Description           string    `yaml:"description,omitempty" json:"description,omitempty" url:"description,omitempty"`
	CreateTime            time.Time `yaml:"create_time,omitempty" json:"create_time,omitempty" url:"create_time,omitempty"`
}

func (cli *Client) DescribeCerts(params DescribeCertsRequest) (resp DescribeCertsResponse, err error) {
	err = cli.MethodGET("DescribeServerCertificates", params, &resp)
	return
}
