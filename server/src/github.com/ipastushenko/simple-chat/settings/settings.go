package settings

type Server struct {
    Port int `json:"port"`
    SecretKeyPath string `json:"secret_key_path"`
    VerifyKeyPath string `json:"verify_key_path"`
    TokenExpiration int `json:"token_expiration"`
}

type Config struct {
    Env string
    Server Server `json:"server"`
    ApiVersion string `json:"apiVersion"`
}
