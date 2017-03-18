package settings

type Server struct {
    Port int `json:"port"`
}

type Config struct {
    Env string
    Server Server `json:"server"`
    ApiVersion string `json:"apiVersion"`
}
