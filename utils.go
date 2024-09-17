package main

import (
	"encoding/binary"
	"os"
)

// Fonction pour extraire l'opcode d'une instruction
func extractOpcode(inst uint32) uint32 {
	return inst & 0x7F // Masque pour extraire les 7 bits de poids faible
}

// Fonction pour étendre les valeurs immédiates de 12 bits à 32 bits
func extendImmediate(imm uint32) int32 {
	if imm&0x800 != 0 {
		return int32(imm | 0xFFFFF000)
	}
	return int32(imm)
}

// Fonction pour lire un fichier binaire et retourner les instructions
func readBinaryFile(filePath string) ([]uint32, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var instructions []uint32
	var inst uint32
	for {
		err := binary.Read(file, binary.LittleEndian, &inst)
		if err != nil {
			break
		}
		instructions = append(instructions, inst)
	}
	return instructions, nil
}
