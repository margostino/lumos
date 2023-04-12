package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/margostino/lumos/common"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

type Credentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func Append(variable *Variable) error {
	ctx := context.Background()

	privateKey := common.NewString(os.Getenv("GAPI_PRIVATE_KEY")).
		ReplaceAll("\\n", "\n").
		ReplaceAll("\\", "").
		Value()

	credentialsJson := &Credentials{
		Type:                    os.Getenv("GAPI_CREDENTIALS_TYPE"),
		ProjectId:               os.Getenv("GAPI_PROJECT_ID"),
		PrivateKeyId:            os.Getenv("GAPI_PRIVATE_KEY_ID"),
		PrivateKey:              privateKey,
		ClientEmail:             os.Getenv("GAPI_CLIENT_EMAIL"),
		ClientId:                os.Getenv("GAPI_CLIENT_ID"),
		AuthUri:                 "https://accounts.google.com/o/oauth2/auth",
		TokenUri:                "https://oauth2.googleapis.com/token",
		AuthProviderX509CertUrl: "https://www.googleapis.com/oauth2/v1/certs",
		ClientX509CertUrl:       "https://www.googleapis.com/robot/v1/metadata/x509/lumos-923%40climateline.iam.gserviceaccount.com",
	}

	credentialsJsonBytes, err := json.Marshal(credentialsJson)
	if err != nil {
		fmt.Printf("Error marshalling struct to JSON: %v", err)
		return err
	}

	creds, err := google.CredentialsFromJSON(ctx, credentialsJsonBytes, "https://www.googleapis.com/auth/spreadsheets")

	if err != nil {
		fmt.Printf("Unable to find default credentials: %v", err)
		return err
	}

	api, err := sheets.NewService(ctx, option.WithCredentials(creds))

	if err != nil {
		fmt.Printf("Unable to create Google Sheets API service: %v", err)
		return err
	}

	if !common.IsError(err, "when creating new Google API Service") {
		spreadsheetId := os.Getenv("SPREADSHEET_ID")
		updateRange := os.Getenv("SPREADSHEET_RANGE")

		var values = [][]interface{}{
			{variable.Datetime, variable.Latitude, variable.Longitude, variable.Name, variable.Value, variable.Observation},
		}

		valueRange := &sheets.ValueRange{Values: values}
		valueInputOption := "RAW"
		insertDataOption := "INSERT_ROWS"

		resp, err := api.Spreadsheets.Values.Append(spreadsheetId, updateRange, valueRange).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()

		if resp.HTTPStatusCode != 200 {
			log.Println("Appender call failed")
			return err
		}
	}

	return nil

}
