package main

type Memory struct {
	data []uint32
}

func initMemory(memory *Memory, size uint32, defaultValue uint32) {
	memory.data = make([]uint32, size)
	for i := 0; i < len(memory.data); i++ {
		memory.data[i] = defaultValue
	}
	logDebug("INIT", "Memory initialized with default value %d\n", defaultValue)
}

func readMemory(memory *Memory, address uint32) uint32 {
	if address < uint32(len(memory.data)) {
		return memory.data[address]
	}
	return 0
}

func writeMemory(memory *Memory, address uint32, value uint32) {
	if address < uint32(len(memory.data)) {
		memory.data[address] = value
	}
}

func lenMemory(memory *Memory) uint32 {
	return uint32(len(memory.data))
}
