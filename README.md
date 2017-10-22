# JSON lint

> JSON linter using Go stdlib

## Install

### API

```sh
go get -u -v github.com/moxar/jsonlint
```

### CLI

```sh
go get -u -v github.com/moxar/jsonlint/cmd/jsonlint
```

## Usage

### API

See example in [`main.go`](cmd/jsonlint/main.go).

```go
jsonBytes, _ := ioutil.ReadAll(os.Stdin)
lintedBytes, _ := jsonlint.Lint(jsonBytes)
fmt.Println(string(lintedJSONBytes))
// $  echo '{"foo": "bar", "lorem": 1, "ipsum": false, "dolores": null, "ikamor": {"foo": ["bar", "foo", "bar"]}}' | go run cmd/jsonlint/main.go
```

### CLI

```sh
echo '{"foo": "bar", "lorem": 1, "ipsum": false, "dolores": null, "ikamor": {"foo": ["bar", "foo", "bar"]}}' | jsonlint
echo '{"foo": "bar", "lorem": 1, "ipsum": false, "dolores": null, "ikamor": {"foo": ["bar", "foo", "bar"]}}' | jsonlint | pygmentize -l json
```

## License

[MIT](LICENSE)
