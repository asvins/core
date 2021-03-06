package main

// Config struct for this service
type Config struct {
	Server struct {
		Addr string
		Port string
	}
	Service struct {
		Env string
	}
	Database struct {
		User    string
		DbName  string
		SSLMode string
	}
}

func GetServiceBaseURI(name string) string {
	mapping := map[string]string{"auth": "http://localhost:8001"}
	return mapping[name]
}
