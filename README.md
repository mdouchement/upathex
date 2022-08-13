# upathex

`upathex` is a package that expands paths like a shell interpreter.

## Usage

```go
path := upathex.ExpandTilde("~/bin")
// => /home/mdouchement/bin

path := upathex.ExpandTilde("~root/bin")
// => /root/bin


path := upathex.ExpandEnv("$HOME/bin")
// => /home/mdouchement/bin

path := upathex.ExpandEnv("${HOME}/bin")
// => /home/mdouchement/bin

path := upathex.ExpandEnvWithCustom("$TROLOLO/bin", map[string]string{
    "TROLOLO": "trololo-lolo",
})
// => trololo-lolo/bin
```

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request