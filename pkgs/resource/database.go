package resource

import (
	"context"
	"encoding/json"
	"github.com/qazxcdswe123/pwctl/pkgs/util"
	"net/http"
)

type Database struct {
	Name                    string  `json:"md_name"`
	Connstr                 string  `json:"md_connstr"`
	IsSuperuser             bool    `json:"md_is_superuser"`
	PresetConfigName        string  `json:"md_preset_config_name"`
	Config                  *string `json:"md_config"`
	IsEnabled               bool    `json:"md_is_enabled"`
	LastModifiedOn          string  `json:"md_last_modified_on"`
	Dbtype                  string  `json:"md_dbtype"`
	IncludePattern          *string `json:"md_include_pattern"`
	ExcludePattern          *string `json:"md_exclude_pattern"`
	CustomTags              *string `json:"md_custom_tags"`
	Group                   string  `json:"md_group"`
	Encryption              string  `json:"md_encryption"`
	HostConfig              *string `json:"md_host_config"`
	OnlyIfMaster            bool    `json:"md_only_if_master"`
	PresetConfigNameStandby *string `json:"md_preset_config_name_standby"`
	ConfigStandby           *string `json:"md_config_standby"`
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) List(ctx context.Context, options *ListOptions) ([]Database, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", options.Hostname+"/db", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Token", util.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var databases []Database
	err = json.NewDecoder(resp.Body).Decode(&databases)
	if err != nil {
		return nil, err
	}

	return databases, nil
}
