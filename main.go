package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Définir les options de la ligne de commande
	filePath := flag.String("file", "", "Chemin vers le fichier binaire")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Veuillez spécifier un fichier binaire avec l'option -file")
		os.Exit(1)
	}

	// Charger la table de correspondance des opcodes
	opcodeTable, err := loadOpcodeTable("opcodes.json")
	if err != nil {
		fmt.Printf("Erreur lors du chargement de la table des opcodes: %v\n", err)
		os.Exit(1)
	}

	// Appeler les fonctions de décodage et de désassemblage
	decodeFile(*filePath, opcodeTable)
	disassembleFile(*filePath, opcodeTable)
}
