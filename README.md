# Client for DaData.ru

Forked from https://github.com/ekomobile/dadata.


[![Build Status](https://travis-ci.org/smoreg/dadata.svg)](https://travis-ci.org/smoreg/dadata)
[![GitHub release](https://img.shields.io/github/release/smoreg/dadata.svg)](https://github.com/smoreg/dadata/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/smoreg/dadata)](https://goreportcard.com/report/github.com/smoreg/dadata)
[![GoDoc](https://godoc.org/github.com/smoreg/dadata?status.svg)](https://godoc.org/github.com/smoreg/dadata)

DaData API v2

Implemented [Clean](https://dadata.ru/api/clean/) and [Suggest](https://dadata.ru/api/suggest/) methods.

## Installation

`go get github.com/smoreg/dadata`

## Usage
```go
import (
	"context"
	"fmt"

	"github.com/smoreg/dadata/v2"
	"github.com/smoreg/dadata/v2/api/suggest"
)

func DaDataExample()  {
	api := dadata.NewSuggestApi()

	params := suggest.RequestParams{
		Query: "ул Свободы",
	}

	suggestions, err := api.Address(context.Background(), &params)
	if err != nil {
		return
	}

	for _, s := range suggestions {
		fmt.Printf("%s", s.Value)
	}
}
```

## Licence
MIT see [LICENSE](LICENSE)
