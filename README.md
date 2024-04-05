# query

Go package to parse the results from a query.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]


## Overview

Package[query][] provides convenience functions for parsing the results from a
Query.

```go
type Querier interface {
	Query(cmd string) (string, error)
}
```

## Contributing

To contribute, please fork the repository, create a feature branch, and then
submit a [pull request][].


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
