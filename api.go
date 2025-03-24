package main

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/status"
)

var (
	serverRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubstore_requests",
		Help: "The number of server requests",
	}, []string{"method", "status", "from_cache"})
)

func (s *Server) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	val, err := s.client.CreateIssue(ctx, req)
	serverRequests.WithLabelValues("CreateIssue", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	val, err := s.client.CloseIssue(ctx, req)
	serverRequests.WithLabelValues("CloseIssue", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	val, err := s.client.CommentOnIssue(ctx, req)
	serverRequests.WithLabelValues("CommentOnIssue", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	val, err := s.client.GetIssue(ctx, req)
	serverRequests.WithLabelValues("GetIssue", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	val, err := s.client.GetLabels(ctx, req)
	serverRequests.WithLabelValues("GetLabels", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	val, err := s.client.GetIssues(ctx, req)
	serverRequests.WithLabelValues("GetIssues", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	val, err := s.client.AddLabel(ctx, req)
	serverRequests.WithLabelValues("AddLabel", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error) {
	val, err := s.client.DeleteLabel(ctx, req)
	serverRequests.WithLabelValues("CreateIssue", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	val, err := s.client.GetComments(ctx, req)
	serverRequests.WithLabelValues("GetComments", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	val, err := s.client.GetRepos(ctx, req)
	serverRequests.WithLabelValues("GetRepos", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	val, err := s.client.GetRepo(ctx, req)
	serverRequests.WithLabelValues("GetRepo", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	val, err := s.client.ListFiles(ctx, req)
	serverRequests.WithLabelValues("ListFiles", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	val, err := s.client.GetFile(ctx, req)
	serverRequests.WithLabelValues("GetFile", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	val, err := s.client.UpdateFile(ctx, req)
	serverRequests.WithLabelValues("UpdateFile", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	val, err := s.client.GetProjects(ctx, req)
	serverRequests.WithLabelValues("GetProjects", status.Code(err).String(), "direct")
	return val, err
}
func (s *Server) GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.GetReleasesResponse, error) {
	val, err := s.client.GetReleases(ctx, req)
	serverRequests.WithLabelValues("GetReleases", status.Code(err).String(), "direct")
	return val, err
}
