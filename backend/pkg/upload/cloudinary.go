package upload

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() (*CloudinaryService, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return nil, err
	}
	return &CloudinaryService{cld: cld}, nil
}

func (s *CloudinaryService) UploadImage(ctx context.Context, file multipart.File, folder string) (string, error) {
	resp, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
