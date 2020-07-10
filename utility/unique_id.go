// Copyright © 2020 Debashish Patra <suvam0451@outlook.com> -- MPL-2.0

// Package utility provides utility functions and converters
package utility

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// UID : Returns a unique string name
func UID() (retval string) {
	uuidNew := uuid.New()
	bytearr := []byte(uuidNew.String())
	uid := bytearr[0:8]
	retval = base64.RawURLEncoding.EncodeToString(uid)
	return
}
