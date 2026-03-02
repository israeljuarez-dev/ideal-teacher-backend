package config

type Container struct {
	HTTP *HTTP
}
type HTTP struct {
	Env            string
	URL            string
	Port           string
	AllowedOrigins string
}

func New()(*Container, error){
	
}