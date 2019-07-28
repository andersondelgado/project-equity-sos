package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func StrNoSQLDrive() string {
	//When running locally, get credentials from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	}
	cloudantUrl := os.Getenv("CLOUDANT_URL")

	appEnv, _ := cfenv.Current()
	if appEnv != nil {
		cloudantService, _ := appEnv.Services.WithLabel("cloudantNoSQLDB")
		if len(cloudantService) > 0 {
			cloudantUrl = cloudantService[0].Credentials["url"].(string)
		}
	}

	return cloudantUrl
}

func StrNoSQLDBname() string {
	//When running locally, get credentials from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	}
	cloudantUrl := os.Getenv("CLOUDANT_DBNAME")

	// appEnv, _ := cfenv.Current()
	// if appEnv != nil {
	// 	cloudantService, _ := appEnv.Services.WithLabel("cloudantNoSQLDB")
	// 	if len(cloudantService) > 0 {
	// 		cloudantUrl = cloudantService[0].Credentials["url"].(string)
	// 	}
	// }

	return cloudantUrl
}

func TestConnection() {
	env := EnviromentsRaw()
	dialect := env.DbDialect
	psqlInfo := fmt.Sprintf(`%s:%s@(%s:%s)/%s`,
		env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName)
	db, _ := sqlx.Connect(dialect, psqlInfo)

	_ = db.Ping()
}

func EnviromentsRaw() Enviroment {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("#string(data): ",string(data))
	byteData := []byte(string(data))
	var enviroment Enviroment
	json.Unmarshal(byteData, &enviroment)

	return enviroment
}

func IBMIAMCredentialRaw () IBMIAMCredential{
	data, err := ioutil.ReadFile("config/apiKey_ibm.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("#string(data): ",string(data))
	byteData := []byte(string(data))
	var enviroment IBMIAMCredential
	json.Unmarshal(byteData, &enviroment)

	return enviroment
}

type IBMIAMCredential struct {
	Apikey               string `json:"apikey"`
	Host                 string `json:"host"`
	IamApikeyDescription string `json:"iam_apikey_description"`
	IamApikeyName        string `json:"iam_apikey_name"`
	IamRoleCrn           string `json:"iam_role_crn"`
	IamServiceidCrn      string `json:"iam_serviceid_crn"`
	Url                  string `json:"url"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Port                 string `json:"port"`
}

type Enviroment struct {
	Domain       string       `json:"domain"`
	RemoteHost   []RemoteHost `json:"remote_host"`
	Port         string       `json:"port"`
	DbDialect    string       `json:"db_dialect"`
	DbHost       string       `json:"db_host"`
	DbName       string       `json:"db_name"`
	DbUser       string       `json:"db_user"`
	DbPassword   string       `json:"db_password"`
	DbSslmode    string       `json:"db_sslmode"`
	DbPort       string       `json:"db_port"`
	AppKey       string       `json:"app_key"`
	SmtpHost     string       `json:"smtp_host"`
	SmtpPort     string       `json:"smtp_port"`
	SmtpEmail    string       `json:"smtp_email"`
	SmtpPassword string       `json:"smtp_password"`
	SmtpTls      string       `json:"smtp_tls"`
}

type RemoteHost struct {
	Name string `json:"name"`
}
