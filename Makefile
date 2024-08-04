# Define Go command and flags
GO = go 
GOFLAGS = -ldflags="-s -w"

# Define the target executable 
TARGET = wecat

# Default target: build the executable 
all: $(TARGET)

# Rule to build the target executable
$(TARGET): main.go
	@echo "--- building ---"
	$(GO) build $(GOFLAGS) -o $(TARGET) main.go

# Clean target : remove the target executable
clean:
	@echo "--- clean $(TARGET) ---"
	- rm -f $(TARGET)

# install target : install the target executable
install:
	@echo "-- install $(TARGET)  to /usr/local/bin/ ---"
	cp -a $(TARGET) /usr/local/bin/

# uninstall target : uninstall the target executable
uninstall:
	@echo "--- uninstall $(TARGET) from /usr/local/bin/ ---"
	-rm -f /usr/local/bin/$(TARGET)

# run target : run the target executable
run: $(TARGET)
	@echo "--- start run ${TARGET} ..."
	./$(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...

help:
	@echo "make"
	@echo "make clean"
	@echo "make install"
	@echo "make uninstall"
	@echo "make run"
	@echo "make test"
	@echo "make help"
