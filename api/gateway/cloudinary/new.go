package cloudinary

import (
	"PopcornMovie/config"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

type GateWay interface {
	UploadToCloudinary(ctx context.Context, file *os.File, filePath string) (string, error)
}

type impl struct {
	cld *cloudinary.Cloudinary
}

func (i impl) UploadToCloudinary(ctx context.Context, file *os.File, filePath string) (string, error) {
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	result, err := i.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", errors.WithStack(err)
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}

func New(config config.Configurations, logger *zap.Logger) (GateWay, error) {
	configCloudinary := config.Cloundinary
	newConfigCld, err := cloudinary.NewFromParams(configCloudinary.Name, configCloudinary.ApiKey, configCloudinary.ApiSecret)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.WithStack(err)
	}

	return &impl{
		cld: newConfigCld,
	}, err
}
