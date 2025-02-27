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
	"fmt"
	"io"
	"os"
	"time"

	"github.com/n1neT10ne/aiyou-go-sdk"
	"github.com/spf13/cobra"
)

const version = "v0.2.0"

var (
	// Configuration globale
	cfg *Config
	// Flags de ligne de commande
	configFile  string
	debug       bool
	temperature float64
	retry       int
	prompt      string
	message     string
	token       string
	model       string
	stream      bool
	assistantID string
	listModels  bool
	useStdin    bool
)

// rootCmd représente la commande de base
var rootCmd = &cobra.Command{
	Use:   "aiyou-cli",
	Short: "--] AI.You platform - simple command line client",
	Long: `A fast and simple command line client for AI.You API build with the platform and Cline.
				Complete documentation available on github.com`,
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// 1. Récupérer le chemin du fichier de config depuis les arguments
		configPath := configFile
		if cmd.Flags().Changed("config") {
			configPath = cmd.Flag("config").Value.String()
		}

		// 2. Charger la configuration depuis le fichier
		cfg, err = loadConfig(configPath)
		if err != nil {
			return fmt.Errorf("error loading config: %w", err)
		}

		// Activer le debug si spécifié en ligne de commande
		if cmd.Flags().Changed("debug") {
			cfg.Debug = debug
		}

		// Afficher la config chargée si debug
		if cfg.Debug {
			fmt.Fprintf(os.Stderr, "[DEBUG] Loaded config: stream=%v\n", cfg.Stream)
		}

		// 3. Appliquer les overrides de la ligne de commande
		if cmd.Flags().Changed("temp") {
			cfg.Temperature = temperature
		}
		if cmd.Flags().Changed("retry") {
			cfg.Retry = retry
		}
		if cmd.Flags().Changed("prompt") && prompt != "" {
			cfg.Prompt = prompt
		}
		if cmd.Flags().Changed("message") && message != "" {
			cfg.Message = message
		}
		if cmd.Flags().Changed("token") {
			cfg.Token = token
		}

		// Debug: afficher la source du token
		if cfg.Debug {
			if cmd.Flags().Changed("token") {
				fmt.Fprintf(os.Stderr, "[DEBUG] Token source: command line flag\n")
			} else if os.Getenv("AIYOU_CLI_TOKEN") != "" {
				fmt.Fprintf(os.Stderr, "[DEBUG] Token source: environment variable\n")
			} else if cfg.Token != "" {
				fmt.Fprintf(os.Stderr, "[DEBUG] Token source: config file\n")
			}
		}
		if cmd.Flags().Changed("model") && model != "" {
			cfg.Model = model
		}
		if cmd.Flags().Changed("stream") {
			cfg.Stream = stream
		}
		if cmd.Flags().Changed("assistant") && assistantID != "" {
			cfg.AssistantID = assistantID
		}

		// Afficher la config finale si debug
		if cfg.Debug {
			fmt.Fprintf(os.Stderr, "[DEBUG] Final config after merge: stream=%v\n", cfg.Stream)
		}

		// Lire depuis stdin si le flag est activé
		if useStdin {
			stdinData, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("error reading from stdin: %w", err)
			}
			cfg.Message = string(stdinData)
		}

		// Si --list-models est utilisé, on ignore les autres paramètres
		if listModels {
			if cfg.Token == "" {
				return fmt.Errorf("token is required (via -k/--token flag, AIYOU_CLI_TOKEN environment variable, or token field in config file)")
			}

			models, err := aiyou.ListModels(cfg.Token, aiyou.WithDebug(cfg.Debug))
			if err != nil {
				return fmt.Errorf("error listing models: %w", err)
			}

			// Affichage des noms de modèles uniquement
			// fmt.Println("Available models:")
			for _, model := range models {
				fmt.Printf("- %s\n", model.Name)
			}
			return nil
		}

		// Validation des paramètres requis pour le mode normal
		var errors []string
		if cfg.Message == "" {
			errors = append(errors, "message is required (via -m/--message flag or -i/--stdin)")
		}
		if cfg.Token == "" {
			errors = append(errors, "token is required (via -k/--token flag, AIYOU_CLI_TOKEN environment variable, or token field in config file)")
		}
		if cfg.Temperature < 0.0 || cfg.Temperature > 2.0 {
			errors = append(errors, "temperature must be between 0.0 and 2.0")
		}

		if len(errors) > 0 {
			return fmt.Errorf("validation error(s):\n  - %s", joinErrors(errors))
		}

		// Préparation des options
		opts := []aiyou.Option{
			aiyou.WithDebug(cfg.Debug),
			aiyou.WithTemperature(cfg.Temperature),
			aiyou.WithStream(cfg.Stream),
		}

		if cfg.Retry > 0 {
			opts = append(opts, aiyou.WithRetry(cfg.Retry, 2*time.Second))
		}

		if cfg.Prompt != "" {
			opts = append(opts, aiyou.WithSystemPrompt(cfg.Prompt))
		}

		if cfg.AssistantID != "" {
			opts = append(opts, aiyou.WithAssistantID(cfg.AssistantID))
		}

		// Appel à l'API
		response, err := aiyou.Completion(
			cfg.Model,
			cfg.Token,
			cfg.Message,
			opts...,
		)
		if err != nil {
			return fmt.Errorf("API error: %w", err)
		}

		// Affichage de la réponse
		fmt.Println(response)
		return nil
	},
}

// joinErrors joint les erreurs avec un séparateur
func joinErrors(errors []string) string {
	result := ""
	for i, err := range errors {
		if i > 0 {
			result += "\n  - "
		}
		result += err
	}
	return result
}

func init() {
	// Flags persistants (disponibles pour toutes les commandes)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "aiyou-cli.yaml", "Configuration file path (default : ./aiyou-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode on AI.You package, see JSON requests (default : false))")
	rootCmd.PersistentFlags().Float64VarP(&temperature, "temp", "t", 1.0, "Temperature (0.0-2.0) (default : 1)")
	rootCmd.PersistentFlags().IntVarP(&retry, "retry", "r", 0, "Number of retries (default : 0)")
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "System instructions")
	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message to send to the LLM")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "k", "", "API token")
	rootCmd.PersistentFlags().StringVarP(&model, "model", "M", "", "LLM to use")
	rootCmd.PersistentFlags().BoolVarP(&stream, "stream", "s", false, "Enable streaming mode (default : false)")
	rootCmd.PersistentFlags().StringVarP(&assistantID, "assistant", "a", "", "Assistant ID")
	rootCmd.PersistentFlags().BoolVarP(&listModels, "list-models", "L", false, "List available models")
	rootCmd.PersistentFlags().BoolVarP(&useStdin, "stdin", "i", false, "Read message from standard input")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
