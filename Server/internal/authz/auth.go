package authz

import (
	"context"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type jwt struct {
	token string
}

func NewTokenFromTokenFile(token string) (credentials.PerRPCCredentials, error) {
	data, err := ioutil.ReadFile(token)
	if err != nil {
		return jwt{}, err
	}
	return jwt{string(data)}, nil
}

func (j jwt) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": j.token,
	}, nil
}

func (j jwt) RequireTransportSecurity() bool {
	return true
}
