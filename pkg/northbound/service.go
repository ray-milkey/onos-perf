// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package northbound

import (
	"context"
	perfapi "github.com/onosproject/onos-api/go/onos/perf"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
	"io"
	"math"
)

var log = logging.GetLogger("northbound", "perf")

// NewService returns a new topo Service
func NewService() northbound.Service {
	return &Service{}
}

// Service is a Service implementation for administration.
type Service struct {
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{}
	perfapi.RegisterPerfServiceServer(r, server)
}

// Server implements the gRPC service for administrative facilities.
type Server struct {
}

// Ping issues pong response
func (s *Server) Ping(ctx context.Context, req *perfapi.PingRequest) (*perfapi.PingResponse, error) {
	log.Infof("Received PingRequest %+v", req)
	res := &perfapi.PingResponse{Payload: req.Payload, Timestamp: req.Timestamp}
	log.Infof("Sending PingResponse %+v", res)
	return res, nil
}

// PingStream streams pong responses to ping messages
func (s *Server) PingStream(stream perfapi.PerfService_PingStreamServer) error {
	log.Infof("Received PingStream setup")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Infof("Received PingStream PingRequest %+v", in)
		count := int(math.Max(1, float64(in.GetRepeatCount())))

		for i := 0; i < count; i++ {
			if err := stream.Send(&perfapi.PingResponse{Payload: in.Payload, Timestamp: in.Timestamp}); err != nil {
				return err
			}
		}
	}
}
