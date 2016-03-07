# Golang PHP Serialize

[![Build Status](https://travis-ci.org/mantyr/phpserialize.svg?branch=master)](https://travis-ci.org/mantyr/phpserialize) [![GoDoc](https://godoc.org/github.com/mantyr/phpserialize?status.png)](http://godoc.org/github.com/mantyr/phpserialize/)

## Example

```Go
package main

import (
    php_serialize "github.com/mantyr/phpserialize"
)

func main() {
    v := make(map[string]interface{})
    v["width"] = 300
    v["file"] = "file.jpg"

    w, err := Encode(v)
    fmt.Println(w)         // print `a:2:{s:5:"width";i:300;s:4:"file";s:8:"file.jpg";}` or  `a:2:{s:4:"file";s:8:"file.jpg";s:5:"width";i:300;}`
    fmt.Println(err)       // print <nil>
}
```

## Installation

    $ go get github.com/mantyr/phpserialize


[mantyr]: https://github.com/mantyr
