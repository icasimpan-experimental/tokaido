package constants

import (
	"time"
)

// HTTPProtocol ...
const HTTPProtocol = "http://"

// HTTPSProtocol ...
const HTTPSProtocol = "https://"

// OneYear is a time.Duration representing a year's worth of seconds.
const OneYear = 8760 * time.Hour

// OneDay is a time.Duration representing a day's worth of seconds.
const OneDay = 24 * time.Hour
