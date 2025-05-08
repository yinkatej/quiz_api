package quiz
import (
	// "fmt"
	"github.com/spf13/viper"
)


type Quiz struct {
	Quiz struct {
	Title string `json:"title"`
	Time  int	`json:"time"`
	Questions []Question `json:"questions"`	
}
}
type Question struct {
	Type    string `json:"type"`
	Id 	   int `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Question string `json:"question"`
	SelectedOption string `json:"selectedOption"`
	Options  []string `json:"options"`
}


type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	    Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Database struct {
		Url string `mapstructure:"url"`
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`
}
var any Quiz
func LoadQuiz() (*Quiz, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("questions")
	viper.SetConfigType("json")
	// viper.AddConfigPath(".")
	// viper.SetConfigName("questions")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&any)
	if err != nil {
		return nil, err
	}
	return &any, nil
}