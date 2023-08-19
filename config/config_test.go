package config

import (
	"reflect"
	"testing"
)

func TestConfig_GetConfig(t *testing.T) {
	type fields struct {
		TodoList struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"todolist"`

		Database struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Name     string `yaml:"name"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Driver   string `yaml:"driver"`
		} `yaml:"database"`
	}

	currentFields := fields{
		TodoList: struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}{
			Host: "127.0.0.1",
			Port: "9000",
		},
		Database: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Name     string `yaml:"name"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Driver   string `yaml:"driver"`
		}{
			Host:     "127.0.0.1",
			Port:     "3306",
			Name:     "db",
			Username: "username",
			Password: "pasword",
			Driver:   "driver",
		},
	}

	wantFields := Config{
		TodoList: struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}{
			Host: "127.0.0.1",
			Port: "9000",
		},
		Database: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Name     string `yaml:"name"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Driver   string `yaml:"driver"`
		}{
			Host:     "127.0.0.1",
			Port:     "3306",
			Name:     "db",
			Username: "username",
			Password: "pasword",
			Driver:   "driver",
		},
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name:   "Success Get Config",
			fields: currentFields,
			want:   &wantFields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				TodoList: tt.fields.TodoList,
				Database: tt.fields.Database,
			}
			if got := c.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
