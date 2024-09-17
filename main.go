package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	filePath := flag.String("file", "", "Chemin vers le fichier binaire")
	help := flag.Bool("h", false, "Afficher le menu d'aide")
	flag.Parse()

	if *help {
		fmt.Println("Un désassembleur d'instructions RISC-V RV32I")
		fmt.Println()
		flag.Usage()
		fmt.Println()
		fmt.Println("Exemple:")
		fmt.Println("  sae-emulateur -file program.bin")
		fmt.Println()
		fmt.Println("Auteur:")
		fmt.Println("  -Grégoire Launay--Bécue")
		fmt.Println("  -Gwendal Margely")

		os.Exit(0)
	}

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
