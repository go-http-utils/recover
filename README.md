# recover
[![Build Status](https://travis-ci.org/go-http-utils/recover.svg?branch=master)](https://travis-ci.org/go-http-utils/recover)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/recover/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/recover?branch=master)

HTTP panic recovery middleware for Go.

## Installation

```
go get -u github.com/go-http-utils/recover
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/go-http-utils/recover

## Usage

```go
import (
  "github.com/go-http-utils/recover"
)
```

```go
mux := http.NewServeMux()

// ...

http.ListenAndServe(":8080", recover.Handler(mux, recover.DefaultRecoverHandler))
```
