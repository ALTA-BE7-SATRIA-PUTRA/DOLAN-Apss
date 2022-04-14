package configs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var theSession *session.Session

//GetConfig Initiatilize config in singleton way
func GetSession() *session.Session {
	lock.Lock()
	defer lock.Unlock()

	if theSession == nil {
		theSession = initSession()
	}

	return theSession
}

func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_Region")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_Access_key_ID"), os.Getenv("AWS_Secret_access_key"), ""),
	}))
	return newSession
}
