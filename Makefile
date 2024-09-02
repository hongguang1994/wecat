# Define Go command and flags
GO = go 
GOFLAGS = -ldflags="-s -w"

# Define the target executable 
APP = wecat

# Default target: build the executable 

# Rule to build the target executable
.PHONY: build
build: clean
	@echo "--- building ---"
	$(GO) build $(GOFLAGS) -o $(APP) main.go

# Clean target : remove the target executable
.PHONY: clean
clean:
	@echo "--- clean  ---"
	$(GO) clean

# install target : install the target executable
.PHONY: install
install:
	@echo "-- install $(TARGET)  to /usr/local/bin/ ---"
	cp -a $(TARGET) /usr/local/bin/

# uninstall target : uninstall the target executable
.PHONY: uninstall
uninstall:
	@echo "--- uninstall $(TARGET) from /usr/local/bin/ ---"
	-rm -f /usr/local/bin/$(TARGET)

# run target : run the target executable
.PHONY: run
run: 
	@echo "--- start run ${TARGET} ..."
	./$(TARGET)

# Test target: run Go tests for the project
.PHONY: test
test:
	$(GO) test ./...

.PHONY: help
help:
	@echo "make"
	@echo "make clean"
	@echo "make install"
	@echo "make uninstall"
	@echo "make run"
	@echo "make test"
	@echo "make help"
