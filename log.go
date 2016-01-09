package mdns

import (
	"log"
	"os"
)

var Log = log.New(os.Stderr, "", log.LstdFlags)
