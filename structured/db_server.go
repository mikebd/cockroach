// Copyright 2014 The Cockroach Authors.
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

package structured

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cockroachdb/cockroach/client"
	"github.com/cockroachdb/cockroach/proto"
	"github.com/cockroachdb/cockroach/rpc"
	"github.com/cockroachdb/cockroach/util"
)

const (
	// DBPrefix is the prefix for the structured database endpoint used
	// to interact with a Cockroach node via HTTP RPC.
	DBPrefix = client.StructuredDBEndpoint
)

var allowedEncodings = []util.EncodingType{util.JSONEncoding, util.ProtoEncoding}

// verifyRequest checks for illegal inputs in request proto and
// returns an error indicating which, if any, were found.
// TODO(vivek): fill this up.
func verifyRequest(args proto.StructuredRequest) error {
	return nil
}

// The structured Table API.
var allPublicMethods = map[string]proto.StructuredMethod{
	proto.CreateTable.String():            proto.CreateTable,
	proto.GetTableRow.String():            proto.GetTableRow,
	proto.PutTableRow.String():            proto.PutTableRow,
	proto.ConditionalPutTableRow.String(): proto.ConditionalPutTableRow,
	proto.IncrementTableRow.String():      proto.IncrementTableRow,
	proto.DeleteTableRow.String():         proto.DeleteTableRow,
	proto.DeleteTableRowRange.String():    proto.DeleteTableRowRange,
	proto.ScanTable.String():              proto.ScanTable,
	proto.BatchTable.String():             proto.BatchTable,
}

// createArgsAndReply returns allocated request and response pairs
// according to the specified method. Note that createArgsAndReply
// only knows about public methods and explicitly returns nil for
// an unknown method.
func createArgsAndReply(method string) (proto.StructuredRequest, proto.StructuredResponse) {
	if m, ok := allPublicMethods[method]; ok {
		switch m {
		// The structured Table API.
		//case proto.CreateTable:
		//	return &proto.CreateTableRequest{}, &proto.CreateTableResponse{}
		case proto.GetTableRow:
			return &proto.GetTableRowRequest{}, &proto.GetTableRowResponse{}
		case proto.PutTableRow:
			return &proto.PutTableRowRequest{}, &proto.PutTableRowResponse{}
		case proto.ConditionalPutTableRow:
			return &proto.ConditionalPutTableRowRequest{}, &proto.ConditionalPutTableRowResponse{}
		case proto.IncrementTableRow:
			return &proto.IncrementTableRowRequest{}, &proto.IncrementTableRowResponse{}
		case proto.DeleteTableRow:
			return &proto.DeleteTableRowRequest{}, &proto.DeleteTableRowResponse{}
		case proto.DeleteTableRowRange:
			return &proto.DeleteTableRowRangeRequest{}, &proto.DeleteTableRowRangeResponse{}
		case proto.ScanTable:
			return &proto.ScanTableRequest{}, &proto.ScanTableResponse{}
		case proto.BatchTable:
			return &proto.BatchTableRequest{}, &proto.BatchTableResponse{}
		}
	}
	return nil, nil
}

// A DBServer provides an HTTP server endpoint serving the key-value API.
// It accepts either JSON or serialized protobuf content types.
type DBServer struct {
	sender client.Sender
}

// NewDBServer allocates and returns a new DBServer.
func NewDBServer(sender client.Sender) *DBServer {
	return &DBServer{sender: sender}
}

// ServeHTTP serves the Structured API by treating the request URL path
// as the method, the request body as the arguments, and sets the
// response body as the method reply. The request body is unmarshalled
// into arguments based on the Content-Type request header. Protobuf
// and JSON-encoded requests are supported. The response body is
// encoded according the the request's Accept header, or if not
// present, in the same format as the request's incoming Content-Type
// header.
func (s *DBServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path
	if !strings.HasPrefix(method, DBPrefix) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	method = strings.TrimPrefix(method, DBPrefix)
	args, reply := createArgsAndReply(method)
	if args == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Unmarshal the request.
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := util.UnmarshalRequest(r, reqBody, args, allowedEncodings); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verify the request for public API.
	if err := verifyRequest(args); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	run(proto.StructuredCall{Args: args, Reply: reply})

	// Marshal the response.
	body, contentType, err := util.MarshalResponse(r, reply, allowedEncodings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType)
	w.Write(body)
}

// RegisterRPC registers the RPC endpoints.
func (s *DBServer) RegisterRPC(rpcServer *rpc.Server) error {
	return rpcServer.RegisterName("Server", (*rpcDBServer)(s))
}

func run(call proto.StructuredCall) {
	call.Reply.Header().Error = &proto.Error{Message: "Not implemented"}
}

// rpcDBServer is used to provide a separate method namespace for RPC
// registration.
type rpcDBServer DBServer

// s a client.Call struct and sends it via our local sender.
func (s *rpcDBServer) executeCmd(args proto.StructuredRequest, reply proto.StructuredResponse) error {
	run(proto.StructuredCall{Args: args, Reply: reply})
	return nil
}

func (s *rpcDBServer) GetTableRow(args *proto.GetTableRowRequest, reply *proto.GetTableRowResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) PutTableRow(args *proto.PutTableRowRequest, reply *proto.PutTableRowResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) ConditionalPutTableRow(args *proto.ConditionalPutTableRowRequest, reply *proto.ConditionalPutTableRowResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) IncrementTableRow(args *proto.IncrementTableRowRequest, reply *proto.IncrementTableRowResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) DeleteTableRow(args *proto.DeleteTableRowRequest, reply *proto.DeleteTableRowResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) DeleteTableRowRange(args *proto.DeleteTableRowRangeRequest, reply *proto.DeleteTableRowRangeResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) ScanTable(args *proto.ScanTableRequest, reply *proto.ScanTableResponse) error {
	return s.executeCmd(args, reply)
}

func (s *rpcDBServer) BatchTable(args *proto.BatchTableRequest, reply *proto.BatchTableResponse) error {
	return s.executeCmd(args, reply)
}
