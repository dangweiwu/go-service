package apiconfig

type ApiConfig struct {
	Host       string `yaml:"host" default:"127.0.0.1:8000" validate:"required"`
	OpenGinLog bool   `yaml:"openGinLog" default:"false"`
	ViewDir    string `yaml:"viewDir" default:"./view"`
	Mode       string `yaml:"mode" default:"debug"`
	Domain     string `yaml:"domain" default:""` //提供域名则开启https
}
