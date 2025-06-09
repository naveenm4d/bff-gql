package config

import (
	"os"

	flag "github.com/spf13/pflag"
)

type Configuration struct {
	HTTPPort *string

	BlogSvcURL *string
}

var (
	Config *Configuration

	httpPort = flag.String(
		"http-port",
		"8000",
		"the port to serve on")

	blogSvcURL = flag.String(
		"blog-svc-url",
		"127.0.0.1:8001",
		"Blog Service URL")
)

func updateStringEnvVariable(defValue *string, key string) *string {
	val := os.Getenv(key)

	if val == "" {
		return defValue
	}

	return &val
}

func init() {
	flag.Parse()

	httpPort = updateStringEnvVariable(httpPort, "HTTP_PORT")

	blogSvcURL = updateStringEnvVariable(blogSvcURL, "BLOG_SVC_URL")

	Config = &Configuration{
		HTTPPort:   httpPort,
		BlogSvcURL: blogSvcURL,
	}
}
