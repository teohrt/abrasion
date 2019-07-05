GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=abrasion

run: build
	./$(BINARY_NAME)

run-v: build
	./$(BINARY_NAME) -verbose

build:
	$(GOBUILD)
	
clean:
	rm -f $(BINARY_NAME)
	rm *.csv