package todo

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Todos []Todo

func Load(path string) (Todos, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	if path[:2] == "~/" {
		path = filepath.Join(home, path[2:])
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Todos{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var todos Todos
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}
	return todos, nil

}

func Save(path string, todos Todos) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	if path[:2] == "~/" {
		path = filepath.Join(home, path[2:])
	}
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
