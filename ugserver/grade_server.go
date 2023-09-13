package ugserver

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_growth/pb"
)

type UgGradeServer struct {
	pb.UnimplementedUserGradeServer
}

func (s *UgGradeServer) ListGrades(ctx context.Context, in *pb.ListGradesRequest) (*pb.ListGradesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGrades not implemented")
}
func (s *UgGradeServer) ListGradePrivileges(ctx context.Context, in *pb.ListGradePrivilegesRequest) (*pb.ListGradePrivilegesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGradePrivileges not implemented")
}
func (s *UgGradeServer) CheckUserPrivilege(ctx context.Context, in *pb.CheckUserPrivilegeRequest) (*pb.CheckUserPrivilegeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserPrivilege not implemented")
}
func (s *UgGradeServer) UserGradeInfo(ctx context.Context, in *pb.UserGradeInfoRequest) (*pb.UserGradeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGradeInfo not implemented")
}
func (s *UgGradeServer) UserGradeChange(ctx context.Context, in *pb.UserGradeChangeRequest) (*pb.UserGradeChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGradeChange not implemented")
}
