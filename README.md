# haags
Read from standard input and output a Haags translation of the given input.

## Building
```bash
make && sudo make install
```
> You may also run `go build` on systems without GNU make.

## Example
```bash
# Use it as a REPL
haags

# Pipe a string to it
echo "ik heb werkelijk niets beters te doen met mijn leven" | haags

# Pipe a file to it
haags < something.txt
```

## Author
Kees van Voorthuizen

## License
[MIT](./LICENSE)
