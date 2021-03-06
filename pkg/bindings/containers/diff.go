package containers

import (
	"context"
	"net/http"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/storage/pkg/archive"
)

// Diff provides the changes between two container layers
func Diff(ctx context.Context, nameOrID string) ([]archive.Change, error) {
	conn, err := bindings.GetClient(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.DoRequest(nil, http.MethodGet, "/containers/%s/changes", nil, nil, nameOrID)
	if err != nil {
		return nil, err
	}
	var changes []archive.Change
	return changes, response.Process(&changes)
}
