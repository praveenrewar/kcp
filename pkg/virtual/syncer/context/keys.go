/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"context"
	"errors"
)

// syncTargetNameContextKeyType is the type of the key for the request context value
// that will carry the name of the SyncTarget resources with be synced with.
type syncTargetNameContextKeyType string

// apiDomainKeyContextKey is the key for the request context value
// that will carry the name of the SyncTarget resources with be synced with.
const syncTargetNameContextKey syncTargetNameContextKeyType = "SyncerVirtualWorkspaceSyncTargetKey"

// WithSyncTargetName adds a SyncTarget name to the context.
func WithSyncTargetName(ctx context.Context, syncTargetName string) context.Context {
	return context.WithValue(ctx, syncTargetNameContextKey, syncTargetName)
}

// SyncTargetNameFrom retrieves the SyncTarget name key from the context, if any.
func SyncTargetNameFrom(ctx context.Context) (string, error) {
	wcn, hasSyncTargetName := ctx.Value(syncTargetNameContextKey).(string)
	if !hasSyncTargetName {
		return "", errors.New("context must contain a valid non-empty SyncTarget name")
	}
	return wcn, nil
}
