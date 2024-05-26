package grpcserver

import "go-sport-lines-practice/api/proto/pkg/lines"

type Server struct {
	lines.UnimplementedLinesServiceServer
}
