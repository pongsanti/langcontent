package route

import "errors"

var errConnectingDatabase = errors.New("Error connecting database")
var errContentTypeRequired = errors.New("Content type is required")

var errMcontentIdInvalid = errors.New("Content ID invalid")
var errLanguageRequired = errors.New("Language is required")
var errTitleRequired = errors.New("Title is required")
