package qingyun

type RunInstancesRequest struct {
	ImageID       string   `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           int      `yaml:"cpu,omitempty" json:"cpu,omitempty" url:"cpu,omitempty"`
	CpuModel      string   `yaml:"cpu_model,omitempty" json:"cpu_model,omitempty" url:"cpu_model,omitempty"`
	Memory        int      `yaml:"memory,omitempty" json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass int      `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	InstanceType  string   `yaml:"instance_type,omitempty" json:"instance_type,omitempty" url:"instance_type,omitempty"`
	OsDiskSize    int      `yaml:"os_disk_size,omitempty" json:"os_disk_size,omitempty" url:"os_disk_size,omitempty"`
	LoginMode     string   `yaml:"login_mode,omitempty" json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string   `yaml:"login_keypair,omitempty" json:"login_keypair,omitempty" url:"login_keypair,omitempty,dotnumbered,numbered1"`
	InstanceName  string   `yaml:"instance_name,omitempty" json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string `yaml:"vxnets,omitempty" json:"vxnets,omitempty" url:"vxnets,omitempty,dotnumbered,numbered1"`
	GPU           int      `yaml:"gpu,omitempty" json:"gpu,omitempty" url:"gpu,omitempty"`
	GpuClass      int      `yaml:"gpu_class,omitempty" json:"gpu_class,omitempty" url:"gpu_class,omitempty"`
}

type RunInstancesResponse struct {
	Action    string   `json:"action"`
	Instances []string `json:"instances"`
	JobID     string   `json:"job_id"`
	RetCode   int      `json:"ret_code"`
	Message   string   `json:"message"`
}

func (cli *Client) RunInstances(params RunInstancesRequest) (resp RunInstancesResponse, err error) {
	err = cli.MethodGET("RunInstances", params, &resp)
	return
}

type DescribeInstancesRequest struct {
	Instances     []string `yaml:"instances,omitempty" json:"instances,omitempty" url:"instances,omitempty,dotnumbered,numbered1"`
	ImageID       string   `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	InstanceType  string   `yaml:"instance_type,omitempty" json:"instance_type,omitempty" url:"instance_type,omitempty,dotnumbered,numbered1"`
	InstanceClass int      `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	Status        []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Verbose       int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
}

type DescribeInstancesResponse struct {
	Action      string        `json:"action"`
	InstanceSet []InstanceSet `json:"instance_set"`
	RetCode     int           `json:"ret_code"`
	TotalCount  int           `json:"total_count"`
}

type InstanceSet struct {
	Vxnets           []InstanceSetVxnet  `json:"vxnets,omitempty"`
	MemoryCurrent    int                 `json:"memory_current,omitempty"`
	VcpusCurrent     int                 `json:"vcpus_current,omitempty"`
	Image            InstanceSetImage    `json:"image,omitempty"`
	InstanceName     string              `json:"instance_name,omitempty"`
	InstanceClass    int                 `json:"instance_class,omitempty"`
	Status           string              `json:"status,omitempty"`
	Description      string              `json:"description,omitempty"`
	ReservedContract string              `json:"reserved_contract,omitempty"`
	VolumeIDS        []string            `json:"volume_ids,omitempty"`
	ZoneID           string              `json:"zone_id,omitempty"`
	InstanceID       string              `json:"instance_id,omitempty"`
	InstanceType     string              `json:"instance_type,omitempty"`
	Volumes          []InstanceSetVolume `json:"volumes,omitempty"`
	KeypairIDS       []string            `json:"keypair_ids,omitempty"`
	Extra            InstanceSetExtra    `json:"extra"`
}

type InstanceSetExtra struct {
	OSDiskSize int `json:"os_disk_size"`
}
type InstanceSetImage struct {
	ImageID string `json:"image_id"`
}

type InstanceSetVolume struct {
	Device   string `json:"device"`
	VolumeID string `json:"volume_id"`
}

type InstanceSetVxnet struct {
	Ipv6Address string `json:"ipv6_address"`
	VxnetType   int    `json:"vxnet_type"`
	VxnetID     string `json:"vxnet_id"`
	VxnetName   string `json:"vxnet_name"`
	Role        int    `json:"role"`
	PrivateIP   string `json:"private_ip"`
	NICID       string `json:"nic_id"`
}

type DescribeInstancesResponseInstanceSet struct{}

func (cli *Client) DescribeInstances(params DescribeInstancesRequest) (resp DescribeInstancesResponse, err error) {
	err = cli.MethodGET("DescribeInstances", params, &resp)
	return
}
