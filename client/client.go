package main

import (
	"context"
	"flag"
	pb "github.com/insurance-policy/insurancepolicy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"log"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

const (
	userId       = 1
	mobileNumber = "1234567892"
)

func printInsuranceResponseById(client pb.InsurancePolicyClient, user *pb.User) {
	log.Printf("Getting policy holder and its policies with userID (%d)", user.ID)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	insuranceResponse, err := client.GetContactAndPoliciesById(ctx, user)
	if err != nil {
		log.Fatalf("%v.GetContactAndPoliciesById(_) = _, %v: ", client, err)
	}
	log.Println(insuranceResponse)
}

func printInsuranceResponseByyMobileNumber(client pb.InsurancePolicyClient, user *pb.User) {
	log.Printf("Getting policy holder and its policies with mobile number (%s)", user.MobileNumber)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insuranceResponse, err := client.GetContactsAndPoliciesByMobileNumber(ctx, user)
	if err != nil {
		log.Fatalf("%v.GetContactsAndPoliciesByMobileNumber(_) = _, %v: ", client, err)
	}
	log.Println(insuranceResponse)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewInsurancePolicyClient(conn)
	// Looking for a valid InsuranceResponse by Id with test data
	printInsuranceResponseById(client, &pb.User{ID: userId})
	// Looking for a valid InsuranceResponse by Mobile Number with test data
	printInsuranceResponseByyMobileNumber(client, &pb.User{MobileNumber: mobileNumber})
}