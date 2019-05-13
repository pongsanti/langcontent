package route

import "errors"

var errConnectingDatabase = errors.New("Error connecting database")
var errContentTypeRequired = errors.New("Content type is required")
