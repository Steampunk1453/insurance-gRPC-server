package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/repository"
	"github.com/insurance-policy/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"log"
	"net"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	port       = flag.Int("port", 10000, "The server port")
)

var server *insurancePolicyServer

type insurancePolicyServer struct {
	pb.UnimplementedInsurancePolicyServer
	amsService service.AmsService
	userService service.UserService
	policyService service.PolicyService
}

func NewInsurancePolicyServer(amsService service.AmsService, userService service.UserService, policyService service.PolicyService) *insurancePolicyServer {
	return &insurancePolicyServer{
		amsService: amsService,
		userService: userService,
		policyService: policyService,
	}
}

func init() {
	// Initialization of the global server variable
	ser := NewInsurancePolicyServer(service.NewAmsService(), service.NewUserService(repository.NewMemdbUserRepository()),
		service.NewPolicyService(repository.NewMemdbPolicyRepository()))
	server = ser
}

func (s *insurancePolicyServer) GetContactAndPoliciesById(ctx context.Context, user *pb.User) (*pb.InsuranceResponse, error) {
	userResp, err := s.getUserById(user.ID)
	if err != nil {
		return nil, err
	}
	policiesResp, err := s.getPoliciesByMobile(userResp.MobileNumber)
	if err != nil {
		return nil, err
	}
	return &pb.InsuranceResponse{User: userResp, Policies: policiesResp}, nil
}

func (s *insurancePolicyServer) GetContactsAndPoliciesByMobileNumber(ctx context.Context, user *pb.User) (*pb.InsuranceResponse, error) {
	userResp, err := s.getUserByMobile(user.MobileNumber)
	policiesResp, err := s.getPoliciesByMobile(user.MobileNumber)
	if err != nil {
		return nil, err
	}
	return &pb.InsuranceResponse{User: userResp, Policies: policiesResp}, nil
}

func (s *insurancePolicyServer) loadUsers()  {
	users := s.amsService.GetUsersByAgentID(1)
	err := s.userService.Save(users)
	log.Printf("Saved users")
	if err != nil {
		log.Fatalf("Failed to load users in DB: %v", err)
	}
}

func (s *insurancePolicyServer) getUserById(id int64) (*pb.User, error) {
	user, err := s.userService.GetByID(id)
	if err != nil {
		log.Fatalf("Failed to get user from DB: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *insurancePolicyServer) getUserByMobile(mobileNumber string) (*pb.User, error) {
	user, err := s.userService.GetByMobileNumber(mobileNumber)
	if err != nil {
		log.Fatalf("Failed to get user from DB: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *insurancePolicyServer) loadPolicies()  {
	policies := s.amsService.GetPoliciesByAgentID(1)
	err := s.policyService.Save(policies)
	log.Printf("Saved policies")
	if err != nil {
		log.Fatalf("Failed to load policies in DB: %v", err)
	}
}

func (s *insurancePolicyServer) getPoliciesByMobile(mobileNumber string) ([]*pb.Policy, error) {
	policies, err := s.policyService.GetByMobileNumber(mobileNumber)
	if err != nil {
		log.Fatalf("Failed to get policies from DB: %v", err)
		return nil, err
	}
	return policies, nil
}

func main() {
	// Retrieve and save data from an external REST API
	server.loadUsers()
	server.loadPolicies()
	// Boot server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInsurancePolicyServer(grpcServer, server)
	grpcServer.Serve(lis)
}
