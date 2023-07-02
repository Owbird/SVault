CRPTIFY_DIR=lib/SVault-Engine
PY_INSTALLER=/env/bin/pyinstaller
WAILS_BUILD=wails build
VERSION=0.1.0

build_engines:
	@echo "Building engines"
	rm -rf build/
	$(CRPTIFY_DIR)/$(PY_INSTALLER) --nowindow --name=ee1 --onefile $(CRPTIFY_DIR)/encrypt/encryptor.py
	$(CRPTIFY_DIR)/$(PY_INSTALLER) --nowindow --name=ee2 --onefile $(CRPTIFY_DIR)/decrypt/decryptor.py
	mkdir -p build/bin
	mv dist/* build/bin/

compile_windows: build_engines
	@echo "Building windows"
	$(WAILS_BUILD) -platform windows -o "SVault-$(VERSION)-windows"

compile_linux: build_engines
	@echo "Building linux"
	$(WAILS_BUILD) -platform linux -o "SVault-$(VERSION)-linux" 


compile_darwin: build_engines
	@echo "Building all"
	$(WAILS_BUILD) -platform darwin -o "SVault-$(VERSION)-darwin" 


compile: build_engines compile_windows compile_linux compile_darwin
