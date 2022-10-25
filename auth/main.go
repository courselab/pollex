package main

import (
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

type TokenPayload struct {
    Token string `json:"token"`
}

func parseAuthKey(tokenKey string) [keySize]byte {
    key, err := hex.DecodeString(tokenKey)
    if err != nil {
        log.Fatal(err)
    }
    if len(key) != keySize {
        log.Fatal("TOKEN_KEY needs to be 32 bytes/64 hex characters")
    }

    var out [keySize]byte
    copy(out[:], key)
    return out
}

func main() {
    // Key used for de/encrypting tokens, 64 hex characters
    tokenKey := os.Getenv("TOKEN_KEY")
    // Key used for authenticating services, needed for generating tokens
    // to prevent impersonation in case this service is exposed to the internet
    serviceKey := os.Getenv("SERVICE_KEY")

    auth, err := MakeAuth(parseAuthKey(tokenKey))
    if err != nil {
        log.Fatal(err)
    }

    if serviceKey == "" {
        log.Fatal("Empty SERVICE_KEY")
    }

    http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, `{"error": "Not found"}`, 404)
            return
        }

        bodyReader := http.MaxBytesReader(w, r.Body, 65536)
        defer bodyReader.Close()

        body, err := ioutil.ReadAll(bodyReader)
        if err != nil {
            http.Error(w, `{"error": "Failed to read body"}`, 400)
            return
        }
        var req TokenPayload
        if err := json.Unmarshal(body, &req); err != nil {
            http.Error(w, `{"error": "Malformed request body"}`, 400)
            return
        }
        data, err := auth.VerifyToken(req.Token)
        if err != nil {
            http.Error(w, `{"error": "Invalid token"}`, 401)
            return
        }
        w.Write(data)
    })

    http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, `{"error": "Not found"}`, 404)
            return
        }

        if r.Header.Get("Authorization") != serviceKey {
            http.Error(w, `{"error": "Unauthorized"}`, 401)
            return
        }

        log.Printf("Generating token for %v", r.RemoteAddr)

        bodyReader := http.MaxBytesReader(w, r.Body, 65536)
        defer bodyReader.Close()

        body, err := ioutil.ReadAll(bodyReader)
        if err != nil {
            http.Error(w, `{"error": "Failed to read body"}`, 400)
            return
        }

        data, err := json.Marshal(&TokenPayload {
            Token: auth.GenerateToken(body),
        })

        if err != nil {
            http.Error(w, `{"error": "Failed to generate token"}`, 500)
            return
        }

        w.Write(data)
    })

    log.Print("Listening on port 6969")
    log.Fatal(http.ListenAndServe(":6969", nil))
}

