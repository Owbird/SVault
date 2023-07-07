CRPTIFY_DIR=lib/SVault-Engine
PY_INSTALLER=/env/bin/pyinstaller
WAILS_BUILD=wails build
VERSION=0.1.0
BUILD_DIR=build
BIN_DIR=$(BUILD_DIR)/bin

# Build engines
build_engines:
	@echo "Building engines"
	rm -rf $(BUILD_DIR)
	$(CRPTIFY_DIR)/$(PY_INSTALLER) --nowindow --name=ee1 --onefile $(CRPTIFY_DIR)/encrypt/encryptor.py
	$(CRPTIFY_DIR)/$(PY_INSTALLER) --nowindow --name=ee2 --onefile $(CRPTIFY_DIR)/decrypt/decryptor.py
	mkdir -p $(BIN_DIR)
	mv dist/* $(BIN_DIR)

# Compile targets
compile_windows: build_engines
	@echo "Building Windows"
	mkdir -p "$(BIN_DIR)/SVault-$(VERSION)-windows"
	$(WAILS_BUILD) -platform windows -o "SVault-$(VERSION)-windows"
	cp $(BIN_DIR)/ee* "$(BIN_DIR)/SVault-$(VERSION)-windows"
	zip -j $(BIN_DIR)/SVault-$(VERSION)-windows.zip $(BIN_DIR)/SVault-$(VERSION)-windows/*
	rm -rf $(BIN_DIR)/SVault-$(VERSION)-windows

compile_linux: build_engines
	@echo "Building Linux"
	mkdir -p "$(BIN_DIR)/SVault-$(VERSION)-linux"
	$(WAILS_BUILD) -platform linux -o "SVault-$(VERSION)-linux"
	cp $(BIN_DIR)/ee* "$(BIN_DIR)/SVault-$(VERSION)-linux"
	zip -j $(BIN_DIR)/SVault-$(VERSION)-linux.zip $(BIN_DIR)/SVault-$(VERSION)-linux/*
	rm -rf $(BIN_DIR)/SVault-$(VERSION)-linux

compile_darwin: build_engines
	@echo "Building macOS"
	mkdir -p "$(BIN_DIR)/SVault-$(VERSION)-darwin"
	$(WAILS_BUILD) -platform darwin -o "SVault-$(VERSION)-darwin"
	cp $(BIN_DIR)/ee* "$(BIN_DIR)/SVault-$(VERSION)-darwin"
	zip -j $(BIN_DIR)/SVault-$(VERSION)-darwin.zip $(BIN_DIR)/SVault-$(VERSION)-darwin/*
	rm -rf $(BIN_DIR)/SVault-$(VERSION)-darwin

# Main target
compile: compile_windows compile_linux compile_darwin
	rm $(BIN_DIR)/ee*
