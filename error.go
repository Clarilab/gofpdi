package gofpdi

import "github.com/pkg/errors"

// ErrInvalidPageNumber indicates that the provided page number is invalid,
// either because it's less or equal to zero or bigger max pages of the pdf.
var ErrInvalidPageNumber = errors.New("provided page number is invalid")