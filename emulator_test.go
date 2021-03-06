package main

import (
	"bytes"
	"testing"
)

func TestAddJmp(t *testing.T) {
	e, _ := run(t, "guest/addjmp.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x0029)
	assetRegister32(t, e, "ECX", ECX, 0x0000)
	assetRegister32(t, e, "EDX", EDX, 0x0080)
	assetRegister32(t, e, "EBX", EBX, 0x0000)
	assetRegister32(t, e, "ESI", ESI, 0x0000)
	assetRegister32(t, e, "EDI", EDI, 0x0000)
	assetRegister32(t, e, "EBP", EBP, 0x0000)
}

func TestCall(t *testing.T) {
	e, _ := run(t, "guest/call-test.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x00f1)
	assetRegister32(t, e, "ECX", ECX, 0x011a)
	assetRegister32(t, e, "EDX", EDX, 0x0080)
	assetRegister32(t, e, "EBX", EBX, 0x0029)
	assetRegister32(t, e, "ESI", ESI, 0x0000)
	assetRegister32(t, e, "EDI", EDI, 0x0000)
	assetRegister32(t, e, "EBP", EBP, 0x0000)
}

// func TestInc(t *testing.T) {
// 	e := run(t, "guest/inc.bin")
// 	assetRegister32(t, e, "EAX", EAX, 0x0000)
// 	assetRegister32(t, e, "ECX", ECX, 0x0000)
// 	assetRegister32(t, e, "EDX", EDX, 0x0000)
// 	assetRegister32(t, e, "EBX", EBX, 0x0000)
// 	assetRegister32(t, e, "ESI", ESI, 0x0000)
// 	assetRegister32(t, e, "EDI", EDI, 0x0000)
// 	assetRegister32(t, e, "EBP", EBP, 0x0000)
// }

func TestModRM(t *testing.T) {
	e, _ := run(t, "guest/modrm-test.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x0002)
	assetRegister32(t, e, "ECX", ECX, 0x0000)
	assetRegister32(t, e, "EDX", EDX, 0x0080)
	assetRegister32(t, e, "EBX", EBX, 0x0000)
	assetRegister32(t, e, "ESI", ESI, 0x0007)
	assetRegister32(t, e, "EDI", EDI, 0x0008)
	assetRegister32(t, e, "EBP", EBP, 0x7be0)
}

func Test132(t *testing.T) {
	e, _ := run(t, "guest/test132.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x0003)
	assetRegister32(t, e, "ECX", ECX, 0x0000)
	assetRegister32(t, e, "EDX", EDX, 0x0080)
	assetRegister32(t, e, "EBX", EBX, 0x0000)
	assetRegister32(t, e, "ESI", ESI, 0x0000)
	assetRegister32(t, e, "EDI", EDI, 0x0000)
	assetRegister32(t, e, "EBP", EBP, 0x0000)
}

func Test133(t *testing.T) {
	e, _ := run(t, "guest/test133.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x0037)
	assetRegister32(t, e, "ECX", ECX, 0x0000)
	assetRegister32(t, e, "EDX", EDX, 0x0080)
	assetRegister32(t, e, "EBX", EBX, 0x0000)
	assetRegister32(t, e, "ESI", ESI, 0x0000)
	assetRegister32(t, e, "EDI", EDI, 0x0000)
	assetRegister32(t, e, "EBP", EBP, 0x0000)
}

func Test134(t *testing.T) {
	e, _ := run(t, "guest/test134.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x0000)
}

func Test135(t *testing.T) {
	e, _ := run(t, "guest/test135.bin", true)
	assetRegister32(t, e, "EAX", EAX, 0x800fffff)
}

func Test141(t *testing.T) {
	e, _ := run(t, "guest/test141.bin", true)
	// expected := "A\x0a"
	//
	// if actual != expected {
	// 	t.Fatalf("expected=%x actual=%x", expected, actual)
	// }
	assetRegister32(t, e, "EAX", EAX, 0x000a)
	assetRegister32(t, e, "ECX", ECX, 0x0000)
	assetRegister32(t, e, "EDX", EDX, 0x03f8)
	assetRegister32(t, e, "EBX", EBX, 0x0000)
	assetRegister32(t, e, "ESI", ESI, 0x0000)
	assetRegister32(t, e, "EDI", EDI, 0x0000)
	assetRegister32(t, e, "EBP", EBP, 0x0000)
}

func TestMbr(t *testing.T) {
	_, actual := run(t, "guest/mbr.bin", false)

	expected := "Congratulations!\x0d\x0a" +
		"You are on a way to hacker!!\x0d\x0a" +
		"The system has halted.\x0a"

	if expected != actual {
		t.Fatalf("\nexpected=%x\nactual  =%x\n", expected, actual)
	}
}

func run(t *testing.T, filename string, protectedEnable bool) (*Emulator, string) {
	bin, err := LoadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}
	e := NewEmulator(0x7c00+0x10000, 0x7c00, 0x7c00, protectedEnable, true, reader, writer, map[uint64]string{})
	// for i := uint32(0); i < 0x7c00 + 0x10000; i++ {
	// 	e.memory[i] = 0
	// }
	for i := 0; i < len(bin); i++ {
		e.memory[uint32(i+0x7c00)] = bin[i]
	}
	for e.eip < 0x7c00+0x10000 {
		err := e.execInst()
		if err != nil {
			t.Fatal(err.Error())
		}
		if e.eip == 0 || e.eip == 0x7c00 {
			break
		}
	}
	return e, writer.String()
}

func assetRegister32(t *testing.T, e *Emulator, name string, index uint8, expected uint32) {
	if e.getRegister32(index) != expected {
		t.Fatalf("Bad %s, expected=%08x, actual=%08x\n",
			name, expected, e.getRegister32(index))
	}
}
