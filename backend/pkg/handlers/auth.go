package handlers

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
    "path"

	"github.com/gin-gonic/gin"
)

type errMsg struct {
    Message string `json:"message"`
}

func checkAuth(authServiceBase *url.URL) func(*gin.Context) {
    //copy it
    u := *authServiceBase
    u.Path = path.Join(u.Path, "verify")

    verifyUrl := u.String()

    return func(context *gin.Context) {
        token := context.Request.Header.Get("Authorization")
        if token == "" {
            context.AbortWithStatusJSON(401, errMsg { "Missing 'Authorization' header" })
            return
        }

        verifyReq, err := json.Marshal(&struct { Token string `json:"token"` } { token })
        if err != nil {
            //should never happen
            panic(err)
        }

        resp, err := http.Post(verifyUrl, "application/json", bytes.NewReader(verifyReq))
        if err != nil {
            context.AbortWithStatusJSON(500, errMsg { "Failed to verify authentication token" })
            return
        }

        //not 2xx
        if resp.StatusCode / 100 != 2 {
            context.AbortWithStatusJSON(401, errMsg { "Invalid or malformed authentication token" })
            return
        }

        tokenData, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            context.AbortWithStatusJSON(500, errMsg { "Failed to verify authentication token" })
            return
        }

        context.Set("auth", tokenData)
    }
}

