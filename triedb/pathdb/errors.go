// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package pathdb

import "errors"

var (
	// errDatabaseReadOnly is returned if the database is opened in read only mode
	// to prevent any mutation.
	errDatabaseReadOnly = errors.New("read only")

	// errDatabaseWaitSync is returned if the initial state sync is not completed
	// yet and database is disabled to prevent accessing state.
	errDatabaseWaitSync = errors.New("waiting for sync")

	// errSnapshotStale is returned from data accessors if the underlying layer
	// had been invalidated due to the chain progressing forward far enough
	// to not maintain the layer's original state.
	errSnapshotStale = errors.New("layer stale")

	// errUnexpectedHistory is returned if an unmatched state history is applied
	// to the database for state rollback.
	errUnexpectedHistory = errors.New("unexpected state history")

	// errStateUnrecoverable is returned if state is required to be reverted to
	// a destination without associated state history available.
	errStateUnrecoverable = errors.New("state is unrecoverable")

	// errWriteImmutable is returned if write to background immutable nodecache
	// under asyncnodebuffer
	errWriteImmutable = errors.New("write immutable nodecache")

	// errFlushMutable is returned if flush the background mutable nodecache
	// to disk, under asyncnodebuffer
	errFlushMutable = errors.New("flush mutable nodecache")

	// errIncompatibleMerge is returned when merge node cache occurs error.
	errIncompatibleMerge = errors.New("incompatible nodecache merge")

	// errRevertImmutable is returned if revert the background immutable nodecache
	errRevertImmutable = errors.New("revert immutable nodecache")
)
