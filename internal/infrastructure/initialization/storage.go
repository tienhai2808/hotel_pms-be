package initialization

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	s3        *s3.Client
	presignS3 *s3.PresignClient
}

func InitStorage(cfg config.MinIOConfig) (*Storage, error) {
	staticCredentials := credentials.NewStaticCredentialsProvider(
		cfg.AccessKeyID,
		cfg.SecretAccessKey,
		"",
	)

	aCfg, err := awsCfg.LoadDefaultConfig(
		context.TODO(),
		awsCfg.WithRegion(cfg.Region),
		awsCfg.WithCredentialsProvider(staticCredentials),
		awsCfg.WithRetryMaxAttempts(3),
		awsCfg.WithRetryMode(aws.RetryModeStandard),
		awsCfg.WithHTTPClient(&http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	protocol := "http"
	if cfg.UseSSL {
		protocol += "s"
	}
	client := s3.NewFromConfig(aCfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(fmt.Sprintf("%s://%s", protocol, cfg.Endpoint))
	})

	publicClient := s3.NewFromConfig(aCfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(cfg.PublicDomain)
	})
	presigner := s3.NewPresignClient(publicClient)

	return &Storage{
		client,
		presigner,
	}, nil
}

func (s *Storage) Client() *s3.Client {
	return s.s3
}

func (s *Storage) Presigner() *s3.PresignClient {
	return s.presignS3
}
