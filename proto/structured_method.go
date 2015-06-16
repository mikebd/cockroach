// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Vivek Menezes (vivek@cockroachlabs.com)
//

package proto

// StructuredMethod is the enumerated type for structured methods.
type StructuredMethod int

const (
	// CreateTable creates a table from the specified Schema.
	CreateTable StructuredMethod = iota
	// GetTableRow returns a set of requested columns from a row, respecting
	// a possibly historical timestamp. If the timestamp is 0, returns
	// the most recent value.
	GetTableRow
	// PutTableRow updates/adds a set of columns at a row at the specified
	// timestamp. If the timestamp is 0, the value is set with the current
	// time as timestamp.
	PutTableRow
	// ConditionalPutTableRow sets the value for a set of columns in a row,
	// if the existing column values match the value specified in the request.
	// Specifying a null value for existing means the value must not yet exist.
	ConditionalPutTableRow
	// IncrementTableRow increments the specified column values at a row.
	IncrementTableRow
	// DeleteTableRow deletes a table row.
	DeleteTableRow
	// DeleteiTableRowRange deleted a range of table rows.
	DeleteTableRowRange
	// ScanTable fetches all the specified columns within the row range.
	ScanTable
	// BatchTable runs all the commands in parallel.
	BatchTable
)
