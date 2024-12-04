package jwtconfig

type JwtConfig struct {
	Secret string `yaml:"secret" validate:"required"`
	Exp    int64  `yaml:"exp" validate:"required"`
}
