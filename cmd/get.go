/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type Database struct {
	MdName                string      `json:"md_name"`
	MdGroup               string      `json:"md_group"`
	MdConfig              interface{} `json:"md_config"`
	MdDBType              string      `json:"md_dbtype"`
	MdConnStr             string      `json:"md_connstr"`
	MdEncryption          string      `json:"md_encryption"`
	MdIsEnabled           bool        `json:"md_is_enabled"`
	MdCustomTags          interface{} `json:"md_custom_tags"`
	MdHostConfig          interface{} `json:"md_host_config"`
	MdIsSuperuser         bool        `json:"md_is_superuser"`
	MdConfigStandby       interface{} `json:"md_config_standby"`
	MdOnlyIfMaster        bool        `json:"md_only_if_master"`
	MdExcludePattern      interface{} `json:"md_exclude_pattern"`
	MdIncludePattern      interface{} `json:"md_include_pattern"`
	MdLastModifiedOn      string      `json:"md_last_modified_on"`
	MdPresetConfigName    string      `json:"md_preset_config_name"`
	MdPresetConfigNameStd interface{} `json:"md_preset_config_name_standby"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func prettyPrintJson(str []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, str, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func getDatabase(ctx context.Context, token, url string) ([]Database, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url+"/db", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Token", token)
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

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(databaseSubCommand)
}

var databaseSubCommand = &cobra.Command{
	Use:   "database",
	Short: "Get all monitored databases",
	Run:   runGetDatabase,
}

func runGetDatabase(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := getJWTToken(ctx, username, password, hostname)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := getDatabase(ctx, token, hostname)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyStr, err := prettyPrintJson(jsonBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(prettyStr)
}
