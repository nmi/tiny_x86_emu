rebuild: clean build
build:
	go build
test: guest_bin
	pkgs=$(go list ./... | grep -v /vendor/)
	go vet ${pkgs}
	golint ${pkgs}
	go test ${pkgs} --cover
clean:
	go clean
guest_bin:
	# binary
	gcc -Wl,--entry=inc,--oformat=binary -nostdlib -fno-asynchronous-unwind-tables -o guest/inc.bin guest/inc.c
	nasm -f bin ./guest/addjmp.asm  -o ./guest/addjmp.bin
	nasm -f bin ./guest/modrm-test.asm  -o ./guest/modrm-test.bin
	nasm -f bin ./guest/call-test.asm  -o ./guest/call-test.bin
	# elf
	gcc -c -g -o guest/inc.o guest/inc.c
	# disasm
	# objdump -D -b binary -m i386:x86-64 ./guest/addjmp.bin
	# objdump -D -b binary -m i386:x86-64 ./guest/modrm-test.bin
	# objdump -D -b binary -m i386:x86-64 ./guest/call-test.bin
	ndisasm -b 32 guest/call-test.bin
	# hexdump -C guest/inc.bin
	# objdump -D -b binary -m i386:x86-64 ./guest/inc.bin
	# objdump -d -S -M intel guest/inc.o
