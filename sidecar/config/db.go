package config

type Database struct {
    DBPORT   string `required:"true" envconfig:"DB_PORT"`
    DBNAME   string `required:"true" envconfig:"DB_NAME"`
    DBHOST   string `required:"true" envconfig:"DB_HOST"`
    PASSWORD string `required:"true" envconfig:"DB_PASSWORD"`
    DBUSER   string `required:"true" envconfig:"DB_USER"`
}
