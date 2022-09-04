package aws

type Config struct {
	Aws struct {
		S3 struct {
			Region          string
			Bucket          string
			AccessKeyID     string
			SecretAccessKey string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	// ex) アジアパシフィック (東京): ap-northeast-1
	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = "YOUR Bucket Name"
	c.Aws.S3.AccessKeyID = "YOUR AWS Access key ID"
	c.Aws.S3.SecretAccessKey = "YOUR AWS Secret Access Key"

	return c
}
