CGO_ENABLED=1
CC=x86_64-w64-mingw32-cc
OUTPUT_WINDOWS=./bin/Installer-Windows.exe
OUTPUT_LINUX=./bin/Installer-Linux
THEME=dark

windows:
	@echo "Windows için proje oluşturuluyor..."
	@CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o $(OUTPUT_WINDOWS) ./src/*.go

linux:
	@echo "Linux için proje oluşturuluyor..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_LINUX) ./src/*.go

all:
	@echo "Windows için proje oluşturuluyor..."
	@CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o $(OUTPUT_WINDOWS) ./src/*.go
	@echo "Building a project for Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_LINUX) ./src/*.go

clean:
	@echo "Temizleniyor..."
	@rm -f $(OUTPUT_WINDOWS) $(OUTPUT_LINUX)

help:
	@echo "Windows versiyonunu kurmak için 'make windows' komutunu kullanın."
	@echo "Linux versiyonunu kurmak için 'make linux' komutunu kullanın."
	@echo "Temizlemek için 'make clean' komutunu kullanın."
	
.PHONY: all windows linux clean help
