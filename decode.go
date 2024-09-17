package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// Structure pour stocker la table de correspondance des opcodes
type OpcodeTable map[string]string

// Fonction pour lire le fichier JSON et charger la table de correspondance des opcodes
//
//goland:noinspection GoDeprecation
func loadOpcodeTable(filePath string) (OpcodeTable, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var table OpcodeTable
	err = json.Unmarshal(data, &table)
	if err != nil {
		return nil, err
	}

	return table, nil
}

// Fonction pour déterminer le type d'encodage à partir de l'opcode
func getEncodingType(opcode uint32, table OpcodeTable) string {
	opcodeStr := fmt.Sprintf("0x%02x", opcode)
	if encoding, ok := table[opcodeStr]; ok {
		return encoding
	}
	return "UNKNOWN"
}

// Fonction pour décoder le fichier binaire et afficher les informations en CSV
func decodeFile(filePath string, opcodeTable OpcodeTable) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier: %v\n", err)
		return
	}
	defer file.Close()

	var inst uint32
	offset := 0
	fmt.Println("offset,valeur,opcode,encoding")
	for {
		err := binary.Read(file, binary.LittleEndian, &inst)
		if err != nil {
			break
		}
		opcode := extractOpcode(inst)
		encoding := getEncodingType(opcode, opcodeTable)
		fmt.Printf("%08x,%08x,%02x,%s\n", offset, inst, opcode, encoding)
		offset += 4
	}
}
