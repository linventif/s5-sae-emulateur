package main

import (
	"fmt"
)

// Fonction pour désassembler une instruction
func disassembleInstruction(inst uint32, opcodeTable OpcodeTable) string {
	opcode := extractOpcode(inst)
	encoding := getEncodingType(opcode, opcodeTable)

	switch encoding {
	case "I":
		return disassembleIType(inst)
	case "R":
		return disassembleRType(inst)
	case "S":
		return disassembleSType(inst)
	case "U":
		return disassembleUType(inst)
	case "UJ":
		return disassembleUJType(inst)
	case "SB":
		return disassembleSBType(inst)
	default:
		return "UNKNOWN"
	}
}

// Fonction pour désassembler une instruction de type I
func disassembleIType(inst uint32) string {
	rd := (inst >> 7) & 0x1F
	rs1 := (inst >> 15) & 0x1F
	imm := (inst >> 20) & 0xFFF
	immExtended := extendImmediate(imm)
	return fmt.Sprintf("addi x%d, x%d, %d", rd, rs1, immExtended)
}

// Fonction pour désassembler une instruction de type R
func disassembleRType(inst uint32) string {
	rd := (inst >> 7) & 0x1F
	rs1 := (inst >> 15) & 0x1F
	rs2 := (inst >> 20) & 0x1F
	return fmt.Sprintf("add x%d, x%d, x%d", rd, rs1, rs2)
}

// Fonction pour désassembler une instruction de type S
func disassembleSType(inst uint32) string {
	rs1 := (inst >> 15) & 0x1F
	rs2 := (inst >> 20) & 0x1F
	imm := ((inst >> 7) & 0x1F) | ((inst>>25)&0x7F)<<5
	immExtended := extendImmediate(imm)
	return fmt.Sprintf("sw x%d, %d(x%d)", rs2, immExtended, rs1)
}

// Fonction pour désassembler une instruction de type U
func disassembleUType(inst uint32) string {
	rd := (inst >> 7) & 0x1F
	imm := (inst >> 12) & 0xFFFFF
	return fmt.Sprintf("lui x%d, %d", rd, imm)
}

// Fonction pour désassembler une instruction de type UJ
func disassembleUJType(inst uint32) string {
	rd := (inst >> 7) & 0x1F
	imm := ((inst >> 12) & 0xFF) | ((inst>>20)&0x1)<<11 | ((inst>>21)&0x3FF)<<1 | ((inst>>31)&0x1)<<20
	return fmt.Sprintf("jal x%d, %d", rd, imm)
}

// Fonction pour désassembler une instruction de type SB
func disassembleSBType(inst uint32) string {
	rs1 := (inst >> 15) & 0x1F
	rs2 := (inst >> 20) & 0x1F
	imm := ((inst>>7)&0x1)<<11 | ((inst>>8)&0xF)<<1 | ((inst>>25)&0x3F)<<5 | ((inst>>31)&0x1)<<12
	return fmt.Sprintf("beq x%d, x%d, %d", rs1, rs2, imm)
}

// Fonction pour désassembler un fichier binaire
func disassembleFile(filePath string, opcodeTable OpcodeTable) {
	instructions, err := readBinaryFile(filePath)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier binaire: %v\n", err)
		return
	}

	offset := 0
	for _, inst := range instructions {
		disassembled := disassembleInstruction(inst, opcodeTable)
		fmt.Printf("%08x: %s\n", offset, disassembled)
		offset += 4
	}
}
