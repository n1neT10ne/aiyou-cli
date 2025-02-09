/*
MIT License

Copyright (c) 2025 Cyrille BARTHELEMY

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config représente la configuration du CLI
type Config struct {
	Debug       bool    `yaml:"debug"`
	Temperature float64 `yaml:"temperature"`
	Retry       int     `yaml:"retry"`
	Prompt      string  `yaml:"prompt"`
	Message     string  `yaml:"message"`
	Token       string  `yaml:"token"`
	Model       string  `yaml:"model"`
	Stream      bool    `yaml:"stream"`
	AssistantID string  `yaml:"assistantId"`
}

// loadConfig charge la configuration depuis un fichier YAML et les variables d'environnement
func loadConfig(path string) (*Config, error) {
	config := &Config{
		Temperature: 1.0, // Valeur par défaut
	}

	// Charge la configuration depuis le fichier YAML s'il existe
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(data, config); err != nil {
			return nil, err
		}
	}

	// Charge le token depuis la variable d'environnement si elle existe
	if envToken := os.Getenv("AIYOU_CLI_TOKEN"); envToken != "" && config.Token == "" {
		config.Token = envToken
	}

	return config, nil
}
