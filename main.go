package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
)

type config struct {
	PersonalAccessToken string `env:"bitbucket_token,required"`
	CommitHash          string `env:"commit_hash,required"`
	APIBaseURL          string `env:"api_base_url,required"`
	BuildKey            string `env:"build_key,required"`

	Status      string `env:"status,opt[auto,in_progress,successful,failed]"`
	TargetURL   string `env:"target_url"`
	BuildName   string `env:"build_name"`
	Description string `env:"description"`
}

func getState(status string) string {
	switch status {
	case "in_progress":
		return "INPROGRESS"
	case "successful":
		return "SUCCESSFUL"
	case "failed":
		return "FAILED"
	}

	if os.Getenv("BITRISE_BUILD_STATUS") == "0" {
		return "successful"
	}
	return "failed"
}

func getName(cfg config) string {
	if cfg.BuildName == "" {
		return cfg.BuildKey
	}

	return cfg.BuildName
}

func getTargetURL(cfg config) string {
	if cfg.TargetURL == "" {
		return os.Getenv("BITRISE_BUILD_URL")
	}

	return cfg.TargetURL
}

func getAPIURL(cfg config) string {
	return fmt.Sprintf("%s/build-status/1.0/commits/%s", cfg.APIBaseURL, cfg.CommitHash)
}

func postStatus(cfg config) error {
	params := map[string]string{
		"state":       getState(cfg.Status),
		"key":         cfg.BuildKey,
		"name":        getName(cfg),
		"url":         getTargetURL(cfg),
		"description": cfg.Description,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	url := getAPIURL(cfg)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.PersonalAccessToken))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send the request: %s", err)
	}

	if err := response.Body.Close(); err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("server error: %s", response.Status)
	}

	return err
}

func main() {
	if os.Getenv("commit_hash") == "" {
		log.Errorf("Cannot publish a build status without a commit")
		os.Exit(1)
	}

	var cfg config
	if err := stepconf.Parse(&cfg); err != nil {
		log.Errorf("Error parsing configuration: %s\n", err)
		os.Exit(1)
	}
	stepconf.Print(cfg)

	if err := postStatus(cfg); err != nil {
		log.Errorf("Error posting build status: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
