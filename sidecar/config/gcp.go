package config

import "cloud.google.com/go/storage"

type GCP struct {
    GCSClient   *storage.Client
}
