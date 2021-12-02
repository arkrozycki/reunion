package util

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/arkrozycki/reunion/logger"
)

var log = logger.Get()

func Hash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	hashed := hex.EncodeToString(h.Sum(nil))
	log.Debug().Msgf("Hash %s for %s", hashed, str)
	return hashed
}
