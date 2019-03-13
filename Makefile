GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=abrasion
URL='http://google.com'

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME) $(URL)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	rm -f $(BINARY_NAME)
	rm -i *.csv