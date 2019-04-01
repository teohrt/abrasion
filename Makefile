GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=abrasion

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
clean:
	rm -f $(BINARY_NAME)
	rm -i *.csv