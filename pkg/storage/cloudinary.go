package storage

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/yonisaka/dating-service/config"
)

// Compile time check to verify implements the Storage interface.
var _ Storage = (*cloudinaryClient)(nil)

// cloudinaryClient is a Storage  dummy for testing purpose
type cloudinaryClient struct {
	client *cloudinary.Cloudinary
}

func NewCloudinaryClient(cfg config.Cloudinary) Storage {
	client, _ := cloudinary.NewFromParams(cfg.CloudName, cfg.APIKey, cfg.APISecret)
	return &cloudinaryClient{
		client: client,
	}
}

// Put do nothing
func (s *cloudinaryClient) Put(ctx context.Context, parent, name string, contents []byte, _ bool, contentType string) error {
	_, err := s.client.Upload.Upload(ctx, bytes.NewReader(contents), uploader.UploadParams{
		PublicID:     name,
		Folder:       parent,
		ResourceType: contentType,
	})
	if err != nil {
		return err
	}

	return nil
}

// Delete do nothing
func (s *cloudinaryClient) Delete(ctx context.Context, parent, name string) error {
	_, err := s.client.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: name,
	})
	if err != nil {
		return err
	}

	return nil
}

// Get do nothing
func (s *cloudinaryClient) Get(_ context.Context, _, _ string) ([]byte, error) {
	return nil, nil
}

// GetSignedURL do nothing
func (s *cloudinaryClient) GetSignedURL(parent, name string) (*url.URL, error) {
	image, err := s.client.Image(fmt.Sprintf("%s/%s", parent, name))
	if err != nil {
		return nil, err
	}

	signedURL, err := image.String()
	if err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(signedURL)
	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}
