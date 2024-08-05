package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// This replaces {"query": "$uuid,@uuid0"}
// with {"query":"e72b5bc3-98be-46b1-a154-ce1fb2c24c31,e72b5bc3-98be-46b1-a154-ce1fb2c24c31"}
func InjectUUID(query string) string {
	uuids := []string{}
	for strings.Contains(query, "$uuid") {
		uuid := uuid.NewString()
		uuids = append(uuids, uuid)
		query = strings.Replace(query, "$uuid", uuid, 1)
	}
	for i, uuid := range uuids {
		query = strings.Replace(query, fmt.Sprintf("@uuid%d", i), uuid, 1)
	}

	return query
}
