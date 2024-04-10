# query

Go package to parse the results from a query.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

## Overview

Package [query][] provides convenience functions for parsing the results from a
Query.

```go
type Querier interface {
	Query(cmd string) (string, error)
}
```

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
$ make lint
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[query][] is released under the MIT license. Please see the [LICENSE.txt][]
file for more information.

[godoc badge]: https://godoc.org/github.com/gotmc/query?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/query
[query]: https://github.com/gotmc/query
[LICENSE.txt]: https://github.com/gotmc/query/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/gotmc/query
[report card]: https://goreportcard.com/report/github.com/gotmc/query
