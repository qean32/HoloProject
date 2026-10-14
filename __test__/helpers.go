package test

import "errors"

// errTestsFailed — хотя бы один подтест провален.
var errTestsFailed = errors.New("есть проваленные подтесты")
