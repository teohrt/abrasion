GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=abrasion

run: build
	./$(BINARY_NAME) -verbose

run-limited: build
	./$(BINARY_NAME) -scrapeLimit=150 -verbose

email: build
	./$(BINARY_NAME) -getEmail -verbose

build:
	$(GOBUILD)
	
clean:
	rm -f $(BINARY_NAME)
	rm *.txt