package utils

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
)

func BasicAuthInteceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	const methodToBeIntercepted1 = "/MenuItemsService/CreateMultiple"
	const methodToBeIntercepted2 = "/RestaurantService/Create"

	if info.FullMethod == methodToBeIntercepted1 || info.FullMethod == methodToBeIntercepted2 {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing credentials")
		}

		authHeader, ok := md["authorization"]
		basicHeader := authHeader[0]
		if !ok || len(authHeader) < 1 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		const prefix = "Basic "
		if !strings.HasPrefix(authHeader[0], prefix) {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header")
		}

		encodedCredentials := strings.TrimPrefix(authHeader[0], prefix)
		credentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid base64 credentials")
		}

		parts := strings.SplitN(string(credentials), ":", 2)
		if len(parts) != 2 {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization credentials")
		}

		authorized := basicAuthentication(basicHeader, ctx)
		if !authorized {
			return nil, status.Error(codes.Unauthenticated, "User does not have access")
		}

	}

	return handler(ctx, req)
}

func basicAuthentication(authHeader string, ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/user-management/users/authorize", nil)
	if err != nil {
		log.Fatalf("request creation failed: %s", err)
	}
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		status.Error(codes.Unauthenticated, "request failed")
	}

	if resp.StatusCode == 200 {
		return true
	}

	return false
}
