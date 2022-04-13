package mqtt

import (
	"context"
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/mqtt"
)

type ServerOption func(o *Server)

func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.bOpts = append(s.bOpts, broker.Addrs(addr))
	}
}

func WithLogger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

func WithTLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
		if c != nil {
			s.bOpts = append(s.bOpts, broker.Secure(true))
		}
		s.bOpts = append(s.bOpts, broker.TLSConfig(c))
	}
}

func WithCleanSession(enable bool) ServerOption {
	return func(s *Server) {
		s.bOpts = append(s.bOpts, mqtt.WithCleanSession(enable))
	}
}

func WithAuth(username string, password string) ServerOption {
	return func(s *Server) {
		s.bOpts = append(s.bOpts, mqtt.WithAuth(username, password))
	}
}

func WithClientId(clientId string) ServerOption {
	return func(s *Server) {
		s.bOpts = append(s.bOpts, mqtt.WithClientId(clientId))
	}
}

func Subscribe(topic string, h broker.Handler) ServerOption {
	return func(s *Server) {
		if s.baseCtx == nil {
			s.baseCtx = context.Background()
		}

		_ = s.RegisterSubscriber(topic, h,
			broker.SubscribeContext(s.baseCtx),
			//broker.Queue(topic),
		)
	}
}