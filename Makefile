OUTPUT=build/haags
PREFIX=/usr/local

all: compile strip

compile:
	go build -o $(OUTPUT)

strip:
	strip $(OUTPUT)

install:
	install -Dm755 $(OUTPUT) $(PREFIX)/bin/haags
