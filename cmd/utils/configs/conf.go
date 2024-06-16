package configs

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

type Config struct {
	ALIEN_VAULT_KEY string `json:"alien_vault_key"`
	VIRUS_TOTAL_KEY string `json:"virus_total_key"`
	ABUSE_DB_KEY    string `json:"abuse_db_key"`
	CROWD_SEC_KEY   string `json:"crowd_sec_key"`
}

func confFile() (string, bool) {
	root, err := os.UserHomeDir()
	if err != nil {
		log.Error("Error getting home directory", err)
		return "", false
	}
	path := root + "/.config/chx/conf.json"

	_, statErr := os.Stat(path)
	if os.IsNotExist(statErr) {
		err := CreateConfig(path)
		if err != nil {
			log.Error("Error creating config file", err)
			return "", false
		}
	}

	return path, true
}

func ReadConfig() (*Config, error) {
	conf, ok := confFile()
	if !ok {
		return nil, nil
	}

	content, err := os.ReadFile(conf)
	if err != nil {
		log.Error("error reading conf file", err)
		return nil, errors.New("error reading conf")
	}

	c := Config{}
	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Error("error unmarshalling json from config file", err)
		return nil, errors.New("cannot unmarshal json")
	}

	return &c, nil
}

func UpdateConfig() {
	conf, ok := confFile()
	if !ok {
		return
	}

	// Check if the file is writable
	file, err := os.OpenFile(conf, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error("Config file is not writable", err)
		return
	}
	file.Close()

	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		log.Info("Could not get default editor from env")
		return
	}

	cmd := exec.Command(editor, conf)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Error("error launching editor or finding config file", err)
	}
}

func CreateConfig(confPath string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error("Error getting home directory", err)
		return err
	}

	dirPath := home + "/.config/chx/"
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Error("Error creating config directory", err)
		return err
	}

	content, err := json.MarshalIndent(&Config{}, "", " ")
	if err != nil {
		log.Error("Error marshalling config to JSON", err)
		return err
	}

	err = os.WriteFile(confPath, content, 0644)
	if err != nil {
		log.Error("Error writing config file", err)
		return err
	}

	// Ensure the file has the correct permissions
	err = os.Chmod(confPath, 0644)
	if err != nil {
		log.Error("Error setting file permissions", err)
		return err
	}

	return nil
}
