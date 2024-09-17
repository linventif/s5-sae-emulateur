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
		/*
			Un décodeur d'instruction RISC-V RV32I

			Utilisation: decode_riscv [OPTIONS] FICHIER_BIN

			Arguments:
			FICHIER_BIN Un fichier au format binaire contenant les instructions à décoder

			Options:
			-h Affiche ce message d'aide
		*/
		fmt.Println("Un désassembleur d'instructions RISC-V RV32I")
		fmt.Println()
		fmt.Println("Utilisation: decode_riscv [OPTIONS] FICHIER_BIN")
		fmt.Println()
		fmt.Println("Arguments:")
		fmt.Println("  FICHIER_BIN  Un fichier au format binaire contenant les instructions à décoder")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  -h  Affiche ce message d'aide")
		fmt.Println()
		fmt.Println("Exemples:")
		fmt.Println("  decode_riscv -file program.bin")
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
