package handlers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"

	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var (
    googleClientId string
    gsuiteDomain   string
    authBaseUrl    *url.URL
    authServiceKey string
)

func init() {
    //TODO: figure out a better place for this configuration
    googleClientId = os.Getenv("GOOGLE_CLIENT_ID")
    if googleClientId == "" {
        panic("Missing GOOGLE_CLIENT_ID environment variable")
    }
    gsuiteDomain = os.Getenv("GSUITE_DOMAIN")
    if gsuiteDomain == "" {
        log.Println("GSUITE_DOMAIN not set, accepting any google account")
    }
    base := os.Getenv("AUTH_SERVICE_URL")
    if base != "UNIT_TEST" {
        baseUrl, err := url.Parse(base)
        if len(base) == 0 || err != nil {
            panic("Invalid or missing AUTH_SERVICE_URL")
        }

        authBaseUrl = baseUrl

        authServiceKey = os.Getenv("AUTH_SERVICE_KEY")
        if authServiceKey == "" {
            panic("Missing AUTH_SERVICE_KEY")
        }
    }
}

type googleInfo struct {
    GSuiteDomain  string `json:"hd"`
    FirstName     string `json:"given_name"`
    LastName      string `json:"family_name"`
    Picture       string `json:"picture"`
    Email         string `json:"email"`
    EmailVerified bool   `json:"email_verified"`
    jwt.RegisteredClaims
}

func fetchGooglePubKey(kid string) (string, error) {
    resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	keys := map[string]string{}
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return "", err
	}
	key, ok := keys[kid]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return key, nil
}

func validateGoogleJWT(raw string) (*googleInfo, error) {
    token, err := jwt.ParseWithClaims(raw, &googleInfo{}, func(token *jwt.Token) (interface{}, error) {
        pem, err := fetchGooglePubKey(fmt.Sprintf("%s", token.Header["kid"]))
        if err != nil {
            return nil, err
        }
        return jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*googleInfo)
    if !ok {
        return nil, fmt.Errorf("Invalid google JWT")
    }

    if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
        return nil, fmt.Errorf("Invalid issuer")
    }

    if err := claims.Valid(); err != nil {
        return nil, err
    }

    if !claims.VerifyAudience(googleClientId, true) {
        return nil, fmt.Errorf("Invalid client id")
    }

    if gsuiteDomain != "" && claims.GSuiteDomain != gsuiteDomain {
        return nil, fmt.Errorf("Wrong GSuite domain")
    }

    return claims, nil
}

func (h *handler) googleLogin(c *gin.Context) {
    c.File("assets/login/google.html")
}

func (h *handler) googleLoginCallback(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Failed to read body"))
    }
    params, err := url.ParseQuery(string(body))
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Malformed request body"))
        return
    }

    cookie, _ := c.Request.Cookie("g_csrf_token")
    csrf := params.Get("g_csrf_token")
    if cookie == nil || csrf == "" || cookie.Value != csrf {
        c.AbortWithError(http.StatusForbidden, fmt.Errorf("CSRF mismatch"))
    }

    token := params.Get("credential")
    if token == "" {
        c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Missing token"))
        return
    }
    claims, err := validateGoogleJWT(token)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid google JWT: %w", err))
        return
    }

    body, err = json.Marshal(claims)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    req, err := http.NewRequest("POST", authBaseUrl.String() + "/generate", bytes.NewReader(body))
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    req.Header.Add("Authorization", authServiceKey)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Failed to create token: %w", err))
        return
    }
    defer resp.Body.Close()

    body, err = io.ReadAll(resp.Body)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Failed to create token: %w", err))
        return
    }

    c.Data(200, "application/json", body)
}

