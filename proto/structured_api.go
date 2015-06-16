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

package proto

import gogoproto "github.com/gogo/protobuf/proto"

// StructuredRequest is an interface for RPC requests.
type StructuredRequest interface {
	gogoproto.Message
	// Header returns the request header.
	Header() *TableRequestHeader
	// Method returns the request method.
	Method() StructuredMethod
	// CreateReply creates a new response object.
	CreateReply() StructuredResponse
}

// StructuredResponse is an interface for RPC responses.
type StructuredResponse interface {
	gogoproto.Message
	// Header returns the response header.
	Header() *TableResponseHeader
	// Verify verifies .
	Verify(req StructuredRequest) error
}

// A StructuredCall is a pending database API call.
type StructuredCall struct {
	Args  StructuredRequest  // The argument to the command
	Reply StructuredResponse // The reply from the command
}

// Header implements the StructuredRequest interface for TableRequestHeader.
func (rh *TableRequestHeader) Header() *TableRequestHeader {
	return rh
}

// Header implements the StructuredResponse interface for TableResponseHeader.
func (rh *TableResponseHeader) Header() *TableResponseHeader {
	return rh
}

func (rh *TableResponseHeader) Verify(req StructuredRequest) error {
	return nil
}

// StructuredMethod implements the StructuredRequest interface.
func (*CreateTableRequest) Method() StructuredMethod { return CreateTable }

// StructuredMethod implements the StructuredRequest interface.
func (*GetTableRowRequest) Method() StructuredMethod { return GetTableRow }

// StructuredMethod implements the StructuredRequest interface.
func (*PutTableRowRequest) Method() StructuredMethod { return PutTableRow }

// StructuredMethod implements the StructuredRequest interface.
func (*ConditionalPutTableRowRequest) Method() StructuredMethod { return ConditionalPutTableRow }

// StructuredMethod implements the StructuredRequest interface.
func (*IncrementTableRowRequest) Method() StructuredMethod { return IncrementTableRow }

// StructuredMethod implements the StructuredRequest interface.
func (*DeleteTableRowRequest) Method() StructuredMethod { return DeleteTableRow }

// StructuredMethod implements the StructuredRequest interface.
func (*DeleteTableRowRangeRequest) Method() StructuredMethod { return DeleteTableRowRange }

// StructuredMethod implements the StructuredRequest interface.
func (*ScanTableRequest) Method() StructuredMethod { return ScanTable }

// StructuredMethod implements the StructuredRequest interface.
func (*BatchTableRequest) Method() StructuredMethod { return BatchTable }

// CreateReply implements the StructuredRequest interface.
//func (*CreateTableRequest) CreateReply() StructuredResponse { return &CreateTableResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*GetTableRowRequest) CreateReply() StructuredResponse { return &GetTableRowResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*PutTableRowRequest) CreateReply() StructuredResponse { return &PutTableRowResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*ConditionalPutTableRowRequest) CreateReply() StructuredResponse {
	return &ConditionalPutTableRowResponse{}
}

// CreateReply implements the StructuredRequest interface.
func (*IncrementTableRowRequest) CreateReply() StructuredResponse { return &IncrementTableRowResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*DeleteTableRowRequest) CreateReply() StructuredResponse { return &DeleteTableRowResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*DeleteTableRowRangeRequest) CreateReply() StructuredResponse {
	return &DeleteTableRowRangeResponse{}
}

// CreateReply implements the StructuredRequest interface.
func (*ScanTableRequest) CreateReply() StructuredResponse { return &ScanTableResponse{} }

// CreateReply implements the StructuredRequest interface.
func (*BatchTableRequest) CreateReply() StructuredResponse { return &BatchTableResponse{} }
