// Package builder fournit des fonctionnalités pour compiler le CLI aiyou pour différentes plateformes
// avec des optimisations de taille (stripping des symboles de debug et compression).
package builder

import (
	"fmt"
	"os"
	"os/exec"
)

// BuildAll compile le CLI pour Windows et Linux avec des optimisations de taille.
// Les binaires sont générés dans le dossier dist/ avec les optimisations suivantes :
// - Stripping des symboles de debug via -ldflags="-s -w"
//   - -s : supprime la table des symboles
//   - -w : supprime les informations DWARF pour le debug
//
// Les binaires générés sont :
// - dist/aiyou-cli.exe pour Windows
// - dist/aiyou-cli-linux pour Linux
func BuildAll() error {
	// Build pour Windows
	fmt.Println("Building for Windows...")
	cmdWin := exec.Command("go", "build", "-ldflags", "-s -w", "-o", "../../dist/aiyou-cli.exe")
	cmdWin.Env = append(os.Environ(),
		"GOOS=windows",
		"GOARCH=amd64",
	)
	if err := cmdWin.Run(); err != nil {
		fmt.Printf("Error building Windows binary: %v\n", err)
		return err
	}

	// Build pour Linux
	fmt.Println("Building for Linux...")
	cmdLinux := exec.Command("go", "build", "-ldflags", "-s -w", "-o", "../../dist/aiyou-cli-linux")
	cmdLinux.Env = append(os.Environ(),
		"GOOS=linux",
		"GOARCH=amd64",
	)
	if err := cmdLinux.Run(); err != nil {
		fmt.Printf("Error building Linux binary: %v\n", err)
		return err
	}

	fmt.Printf("Build complete!\nBinaries created in dist/:\n- aiyou-cli.exe (Windows)\n- aiyou-cli-linux (Linux)\n")
	return nil
}
