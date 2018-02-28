all: tiny_x86_emu
tiny_x86_emu:
	go build ./...
test: guest_bin
	pkgs=$(go list ./... | grep -v /vendor/)
	go vet ${pkgs}
	golint ${pkgs}
	go test ${pkgs} --cover
	# goveralls -repotoken $COVERALL_REPO_TOKEN
guest_bin:
	gcc -Wl,--entry=inc,--oformat=binary -nostdlib -fno-asynchronous-unwind-tables -o guest/inc.bin guest/inc.c # binary
	gcc -c -g -o guest/inc.o guest/inc.c # elf
	ndisasm -b 32 guest/inc.bin
	hexdump -C guest/inc.bin
	objdump -d -S -M intel guest/inc.o
