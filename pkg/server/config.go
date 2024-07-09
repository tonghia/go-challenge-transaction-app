package server

import (
	"fmt"
	"net"
)

func (l Listen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

// Listen represents a network end point address.
type Listen struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func (l *Listen) createListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", l.String())
	if err != nil {
		return nil, fmt.Errorf("failed to listen %s: %w", l.String(), err)
	}

	return lis, nil
}

func createDefaultConfig() *serverConfig {
	config := &serverConfig{
		grpc:    createDefaultGrpcConfig(),
		gateway: createDefaultGatewayConfig(),
	}

	return config
}

type serverConfig struct {
	grpc           *grpcConfig
	gateway        *gatewayConfig
	serviceServers []ServiceServer // your grpc-services
}
