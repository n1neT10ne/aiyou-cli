// Package main fournit un point d'entrée pour la compilation multi-plateforme du CLI.
// Ce programme utilise le package builder pour générer des binaires optimisés
// pour Windows et Linux dans le dossier dist/.
package main

import (
	"fmt"
	"os"

	"github.com/n1neT10ne/aiyou-cli/scripts/builder"
)

// main exécute le processus de build multi-plateforme.
// En cas d'erreur, le message d'erreur est affiché sur stderr
// et le programme se termine avec un code de sortie 1.
func main() {
	if err := builder.BuildAll(); err != nil {
		fmt.Fprintf(os.Stderr, "Build failed: %v\n", err)
		os.Exit(1)
	}
}
