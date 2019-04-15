GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=abrasion

run: build
	./$(BINARY_NAME)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
clean:
	rm -f $(BINARY_NAME)
	rm *.csv