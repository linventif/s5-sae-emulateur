package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"os/exec"
	"testing"
)

func TestRiscvDecoder(t *testing.T) {
	// Prepare a binary test file
	testInstructions := []uint32{
		0x00a10093, // addi x1, x2, 10
		0x00800293, // addi x5, x0, 8
		0x00400313, // addi x6, x0, 4
	}

	fileName := "test_instructions.bin"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	for _, instr := range testInstructions {
		err := binary.Write(file, binary.LittleEndian, instr)
		if err != nil {
			t.Fatalf("Failed to write instruction: %v", err)
		}
	}
	file.Close()

	// Run the decoder program
	cmd := exec.Command("go", "run", "./main.go", fileName)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to execute decoder: %v\nOutput: %s", err, out.String())
	}

	// Expected output
	expected := `offset,valeur,opcode,encoding
00000000,00a10093,OP-IMM,I
00000004,00800293,OP-IMM,I
00000008,00400313,OP-IMM,I
`

	if out.String() != expected {
		t.Errorf("Output mismatch\nExpected:\n%s\nGot:\n%s", expected, out.String())
	}
}
