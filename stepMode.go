package main

import (
	"fmt"
	"os"
)

var stepMode = false

func executeCommand(cpu *CPUState, memory *[]uint32, command string, startAddress uint32, defaultRegisterValue uint32) {
	switch command {
	case "step":
		stepMode = true
	case "continue":
		stepMode = false
	case "reset":
		initCPUState(cpu, startAddress, defaultRegisterValue)
		fmt.Println("CPU reset avec PC =", startAddress, "et registre par défaut =", defaultRegisterValue)
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Commande inconnue.")
	}
}

func handleStepMode(cpu *CPUState, memory *[]uint32, startAddress uint32, defaultRegisterValue uint32) {
	for stepMode {
		// Affiche l'état des registres
		fmt.Printf("PC: 0x%08x\n", cpu.pc)
		for i := 0; i < 32; i++ {
			fmt.Printf("x%d: 0x%08x\n", i, cpu.x[i])
		}

		// Affiche l'instruction
		if cpu.pc/4 < uint32(len(*memory)) {
			instruction := (*memory)[cpu.pc/4]
			opcode, _ := GetOpcodeFromInstruction(instruction)
			fmt.Println(opcode.Encoding.Decode(opcode, instruction, cpu, memory))
		} else {
			fmt.Println("Instruction hors mémoire.")
		}

		// Attend la commande suivante
		var command string
		fmt.Print("> ")
		fmt.Scanln(&command)
		executeCommand(cpu, memory, command, startAddress, defaultRegisterValue)
	}
}
