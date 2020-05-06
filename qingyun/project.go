package qingyun

type AddProjectResourceItemsRequest struct {
	ProjectID string   `yaml:"project_id,omitempty" json:"project_id,omitempty" url:"project_id,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`
}

type AddProjectResourceItemsResponse struct {
	Message     string   `json:"message" name:"message"`
	Action      string   `json:"action" name:"action" location:"elements"`
	ProjectID   string   `json:"project_id" name:"project_id" location:"elements"`
	ResourceIDs []string `json:"resource_ids" name:"resource_ids" location:"elements"`
	RetCode     int      `json:"ret_code" name:"ret_code" location:"elements"`
	ZoneID      string   `json:"zone_id" name:"zone_id" location:"elements"`
}

func (cli Client) AddProjectResourceItems(params AddProjectResourceItemsRequest) (resp AddProjectResourceItemsResponse, err error) {
	err = cli.MethodGET("AddProjectResourceItems", params, &resp)
	return
}

type DeleteProjectResourceItemsRequest struct {
	ProjectID []string `yaml:"project_id,omitempty" json:"project_id,omitempty" url:"project_id,omitempty,dotnumbered,numbered1"` // Required
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`    // Required
}

type DeleteProjectResourceItemsResponse struct {
	Message     string   `yaml:"message,omitempty" json:"message,omitempty" url:"message,omitempty"`
	Action      string   `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	ProjectID   []string `yaml:"project_id,omitempty" json:"project_id,omitempty" url:"project_id,omitempty,dotnumbered,numbered1"`
	ResourceIDs []string `yaml:"resource_i_ds,omitempty" json:"resource_i_ds,omitempty" url:"resource_i_ds,omitempty,dotnumbered,numbered1"`
	RetCode     int      `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
	ZoneID      string   `yaml:"zone_id,omitempty" json:"zone_id,omitempty" url:"zone_id,omitempty"`
}
