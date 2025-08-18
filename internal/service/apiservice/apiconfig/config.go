package apiconfig

type ApiConfig struct {
	Host        string `yaml:"host" default:"127.0.0.1:8000" validate:"required"`
	OpenGinLog  bool   `yaml:"openGinLog" default:"false"`
	ViewDir     string `yaml:"viewDir" default:"./view"`
	DocDir      string `yaml:"docDir" default:"./doc"`
	DocUser     string `yaml:"docUser" default:"admin"`
	DocPassword string `yaml:"docPassword" default:"admin"`
	Mode        string `yaml:"mode" default:"debug"`
	Domain      string `yaml:"domain" default:""` //cert domain
	Email       string `yaml:"email"`             //cert email
}
