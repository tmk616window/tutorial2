package config

import "cloud.google.com/go/storage"

type GCP struct {
    GCSClient   *storage.Client `required:"true" envconfig:"DB_PORT"`
}
