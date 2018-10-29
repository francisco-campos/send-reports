package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

const addrMailSender = "smtp.gmail.com:587"
const hostMailSender = "smtp.gmail.com"

//Reports struct represent the structure complete of config.json
type Reports struct {
	Reports []Report `json:"reports"`
}

//Report is a item of reports
type Report struct {
	Title         string        `json:"title"`
	Frequency     string        `json:"frequency"`
	Hour          string        `json:"hour"`
	ConfigMailing ConfigMailing `json:"configMailing"`
	Query         string        `json:"query"`
	Headers       string        `json:"headers"`
}

//ConfigMailing is the configuration for a report
type ConfigMailing struct {
	MailFrom     string   `json:"mailFrom"`
	MailFromPass string   `json:"mailFromPass"`
	MailTo       []string `json:"mailTo"`
	Subject      string   `json:"subject"`
	Message      string   `json:"message"`
}

var reports Reports

//LoadConfiguration carga configuracion de archivo json
func LoadConfiguration(file string) Reports {
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println("error cargar la configuracion: " + err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&reports)
	configFile.Close()
	return reports
}

//SendEmail envia correo electronico
func SendEmail(from string, pass string, to []string, body string) {
	err := smtp.SendMail(addrMailSender,
		smtp.PlainAuth("", from, pass, hostMailSender),
		from, to, []byte(body))

	if err != nil {
		log.Printf("error SendEmail: %s", err)
		return
	}
}

func generateReport(report Report) string {
	return ""
}

func main() {
	fmt.Println("Comenzando ejecución!")
	filePath := "./config.json"
	reports = LoadConfiguration(filePath)
	reportsCount := len(reports.Reports)
	if reportsCount == 0 {
		fmt.Printf("No existen reportes para enviar.\n")
	}

	for _, element := range reports.Reports {
		reportPath := generateReport(element)
		//SendEmail(element.ConfigMailing, reportPath)
	}
}
