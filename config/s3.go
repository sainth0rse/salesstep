package config

import (
	"context"
	"fmt"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioClient – глобальная переменная клиента
var MinioClient *minio.Client

// MinioBucket – имя бакета (создадим в init)
const MinioBucket = "ss-bucket"

func InitMinio() error {
	// Подключаемся к локальному MinIO
	endpoint := getEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKey := getEnv("MINIO_ROOT_USER", "root")
	secretKey := getEnv("MINIO_ROOT_PASSWORD", "rootpassword")
	useSSL := false

	// Создаём клиента
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("ошибка при инициализации MinIO: %w", err)
	}
	MinioClient = client

	// Создаём (или убеждаемся в наличии) бакет
	ctx := context.Background()
	err = client.MakeBucket(ctx, MinioBucket, minio.MakeBucketOptions{})
	if err != nil {
		// Если бакет уже есть, проверим
		exists, errBucketExists := client.BucketExists(ctx, MinioBucket)
		if errBucketExists == nil && exists {
			log.Printf("MinIO bucket %s уже существует\n", MinioBucket)
		} else {
			return fmt.Errorf("не удалось создать или проверить бакет: %w", err)
		}
	} else {
		log.Printf("Создан новый bucket: %s\n", MinioBucket)
	}

	log.Println("MinIO инициализирован успешно!")
	return nil
}
