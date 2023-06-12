package oss

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/web3eye-io/Web3Eye/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var ErrOssClientNotInit = errors.New("oss client not init")

var (
	s3Client  *s3.Client
	_s3Config S3Config
	client    = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 20 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   20 * time.Second,
			ResponseHeaderTimeout: 20 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
		},
	}
)

type S3Config struct {
	Region    string `json:"region"`
	EndPoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket,omitempty"`
}

// nolint
func Init(region, bucket string) error {
	s3Config := S3Config{
		Region:    region,
		EndPoint:  config.GetConfig().Minio.Address,
		AccessKey: config.GetConfig().Minio.AccessKey,
		SecretKey: config.GetConfig().Minio.SecretKey,
		Bucket:    bucket,
	}
	_s3Config = S3Config{
		Region:    s3Config.Region,
		EndPoint:  s3Config.EndPoint,
		AccessKey: s3Config.AccessKey,
		SecretKey: s3Config.SecretKey,
		Bucket:    s3Config.Bucket,
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               s3Config.EndPoint,
			HostnameImmutable: true,
			SigningRegion:     s3Config.Region,
		}, nil
	})

	cfg, err := s3config.LoadDefaultConfig(context.Background(),
		s3config.WithRegion(s3Config.Region),
		s3config.WithHTTPClient(client),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3Config.AccessKey, s3Config.SecretKey, "")),
		s3config.WithEndpointResolverWithOptions(customResolver),
		s3config.WithClientLogMode(aws.LogRetries),
	)
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		// o.EndpointOptions.DisableHTTPS = true
	})

	return nil
}

// GetStringValueWithNameSpace not network invoke
func getS3Bucket() string {
	return _s3Config.Bucket
}

func PutObject(ctx context.Context, key string, body []byte) error {
	if s3Client == nil {
		return ErrOssClientNotInit
	}

	_, err := s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	return err
}

func GetObject(ctx context.Context, key string) ([]byte, error) {
	if s3Client == nil {
		return nil, ErrOssClientNotInit
	}
	s3out, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer s3out.Body.Close()

	out, err := io.ReadAll(s3out.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func UploadFile(ctx context.Context, filePath, key string) error {
	if s3Client == nil {
		return ErrOssClientNotInit
	}

	s, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return errors.New("please input a file path,not a dir path")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func DownloadFile(ctx context.Context, filePath, key string) error {
	if s3Client == nil {
		return ErrOssClientNotInit
	}

	downloadFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	downloader := manager.NewDownloader(s3Client)
	_, err = downloader.Download(context.TODO(), downloadFile, &s3.GetObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(key),
	})
	return err
}

func GetObjectAttributes(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {
	if s3Client == nil {
		return nil, ErrOssClientNotInit
	}

	return s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(getS3Bucket()),
		Key:    aws.String(key),
	})
}
