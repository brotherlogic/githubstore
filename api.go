package main

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
)

func (s *Server) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	return s.client.CreateIssue(ctx, req)
}
func (s *Server) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	return s.client.CloseIssue(ctx, req)
}
func (s *Server) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	return s.client.CommentOnIssue(ctx, req)
}
func (s *Server) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return s.client.GetIssue(ctx, req)
}
func (s *Server) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	return s.client.GetLabels(ctx, req)
}
func (s *Server) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	return s.client.GetIssues(ctx, req)
}
func (s *Server) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	return s.client.AddLabel(ctx, req)
}
func (s *Server) DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error) {
	return s.client.DeleteLabel(ctx, req)
}
func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	return s.client.GetComments(ctx, req)
}
func (s *Server) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	return s.client.GetRepos(ctx, req)
}
func (s *Server) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	return s.client.GetRepo(ctx, req)
}
func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	return s.client.ListFiles(ctx, req)
}
func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	return s.client.GetFile(ctx, req)
}
func (s *Server) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	return s.client.UpdateFile(ctx, req)
}
func (s *Server) GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	return s.client.GetProjects(ctx, req)
}
func (s *Server) GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.GetReleasesResponse, error) {
	return s.client.GetReleases(ctx, req)
}
