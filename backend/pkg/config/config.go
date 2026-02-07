package config
import(
	"log"
	"github.com/spf13/viper"
)
type Config struct{
	Port string `mapstructure:"port"`
}
func Load()(*Config,error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig();err != nil{
		log.Fatalf("Error reading config file,%s",err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg);err!= nil{
		log.Fatalf("Error unmarshalling config file,%s",err)
	}
	return &cfg,nil
}