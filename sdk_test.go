// Copyright 2018, Boon Logic Inc

package amber_client

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	// "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	// amberClient "github.com/boonlogic/amber-go-sdk/client"
	// amberOps "github.com/boonlogic/amber-go-sdk/client/operations"
	// amberModels "github.com/boonlogic/amber-go-sdk/models"
	// "io/ioutil"
	_ "log"
	"os"
)

var licenseProfile LicenseProfile

func init() {
	var err error
	licenseProfile, err = getUserSecrets()
	if err != nil {
		fmt.Printf("getUserSecrets failed\n")
		os.Exit(3)
	}
}

func getUserSecrets() (LicenseProfile, error) {

	var lp licenseProfiles

	// retrieve the deployment from the environment.  If not set, default to 'qa'
	deployment := os.Getenv("AMBER_TEST_PROFILE")
	if deployment == "" {
		deployment = "qa"
	}

	region := "us-east-1"
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("amber-test-users"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(*result.SecretString), &lp); err != nil {
		return LicenseProfile{}, err
	}

	profile := lp[deployment]
	if profile.Username == "" {
		fmt.Printf("deployment %v not found\n", deployment)
		os.Exit(3)
	}

	return LicenseProfile{}, nil
}

func loadCredentialsIntoEnv() {
	os.Setenv("AMBER_USERNAME", licenseProfile.Username)
	os.Setenv("AMBER_PASSWORD", licenseProfile.Password)
	os.Setenv("AMBER_SERVER", licenseProfile.Server)
	os.Setenv("AMBER_OAUTH_SERVER", licenseProfile.OauthServer)
}

func clearEnv() {
	os.Unsetenv("AMBER_USERNAME")
	os.Unsetenv("AMBER_PASSWORD")
	os.Unsetenv("AMBER_SERVER")
	os.Unsetenv("AMBER_OAUTH_SERVER")
}

func restoreEnv(savedProfile LicenseProfile) {
	clearEnv()
	os.Setenv("AMBER_USERNAME", savedProfile.Username)
	os.Setenv("AMBER_PASSWORD", savedProfile.Password)
	os.Setenv("AMBER_SERVER", savedProfile.Server)
	os.Setenv("AMBER_OAUTH_SERVER", savedProfile.OauthServer)
}

// Runs before each test. Initializes by setting env variables.
func TestNewAmberClientFromProfile(t *testing.T) {

	clearEnv()

	// expected license profile
	profileA := LicenseProfile{
		Username:    "mruser",
		Password:    "mypassword",
		Server:      "https://fakeserver.com/v1",
		OauthServer: "https://fakeserver.com/v1/oauth",
	}
	amberClientA, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.Username, "mruser")
	require.Equal(t, amberClientA.licenseProfile.Password, "mypassword")
	require.Equal(t, amberClientA.licenseProfile.Server, "https://fakeserver.com/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth")

	// expected license profile override from environment
	os.Setenv("AMBER_USERNAME", "mruser-env")
	os.Setenv("AMBER_PASSWORD", "mypassword-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.Username, "mruser-env")
	require.Equal(t, amberClientB.licenseProfile.Password, "mypassword-env")
	require.Equal(t, amberClientB.licenseProfile.Server, "https://fakeserver.com/v1/env")
	require.Equal(t, amberClientB.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth/env")
}

func TestNewAmberClientFromFile(t *testing.T) {

	clearEnv()

	id := "default"
	file := "test/test.Amber.license"

	amberClientA, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.Username, "amber-dev-user")
	require.Equal(t, amberClientA.licenseProfile.Password, "phi{Ch2obovoe")
	require.Equal(t, amberClientA.licenseProfile.Server, "http://dev-build-1:5888/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://amber-local.boonlogic.com/qa")

	// expected license profile override from environment
	os.Setenv("AMBER_USERNAME", "mruser-env")
	os.Setenv("AMBER_PASSWORD", "mypassword-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.Username, "mruser-env")
	require.Equal(t, amberClientB.licenseProfile.Password, "mypassword-env")
	require.Equal(t, amberClientB.licenseProfile.Server, "https://fakeserver.com/v1/env")
	require.Equal(t, amberClientB.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth/env")
}
