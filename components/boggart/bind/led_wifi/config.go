package led_wifi

type Config struct {
	Address string `valid:"host,required"`
}

func (t Type) Config() interface{} {
	return &Config{}
}
