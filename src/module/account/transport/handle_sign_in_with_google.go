package accounttrpt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "fmt"
	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "1093847451472-cbaiohq0s3vh8vle1dvc8m076lhpbj4a.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/api/v1/oauth2/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

func HandlerSignInWithGoogle() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		googleConf := SetupConfig()
		url := googleConf.AuthCodeURL("randomstate")

		http.Redirect(ctx.Writer, ctx.Request, url, http.StatusSeeOther)

	}
}

func HandleCallbackGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle the exchange code to initiate a transport.
		// session := sessions.Default(c)
		// retrievedState := session.Get("state")
		// if retrievedState != c.Query("state") {
		// 	c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		// 	return
		// }

		conf := SetupConfig()

		tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		client := conf.Client(oauth2.NoContext, tok)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)

		type GoogleProfile struct {
			email string
		}

		var response GoogleUser
		error := json.Unmarshal([]byte(string(data)), &response)
		if err != nil {
			panic(error)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   response,
		})

	}
}
