// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package nats

import (
	"fmt"
	"time"

	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"

	"github.com/nats-io/nats.go"
	"github.com/saifhamdan/go-apigateway-blueprint/config"
)

type Nats struct {
	*nats.Conn
}

func NewNATS(cfg *config.Config, logger *logger.Logger) (*Nats, error) {
	url := fmt.Sprintf("nats://%s:%s", cfg.NatsHost, cfg.NatsPort)

	logger.Infof("Connecting to NATS server: %s", url)

	opts := []nats.Option{
		nats.Name("go-apigateway-blueprint"),
		nats.MaxReconnects(-1), // Infinite reconnects
		nats.ReconnectWait(2 * time.Second),
		nats.Timeout(10 * time.Second),
		nats.PingInterval(20 * time.Second),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Errorf("Reconnected to NATS server: %v", nc.ConnectedUrl())
		}),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			logger.Errorf("Disconnected from NATS server: %v", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.Errorf("Connection to NATS server closed: %v", nc.LastError())
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %v", err)
	}

	logger.Infof("Connected to NATS server: %s", nc.ConnectedUrl())

	return &Nats{Conn: nc}, nil
}
