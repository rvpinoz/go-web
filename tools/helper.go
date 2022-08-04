package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func UID() string {
	now := time.Now()
	uniq := now.Format(time.RFC3339Nano)
	hash := md5.Sum([]byte(uniq))
	return hex.EncodeToString(hash[:])
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Base64Encode(value string) string {
	var a = base64.StdEncoding.EncodeToString([]byte(value))
	return a
}

func Base64Decode(value string) string {
	var a, _ = base64.StdEncoding.DecodeString(value)
	return string(a)
}

func Sha256(value string) string {
	a := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", a)
}

func Now() time.Time {
	ina, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(ina)
}

func DateNow(format string) string {
	ina, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(ina).Format(format)
}

func DateNowHour(format string, hour int) string {
	ina, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().Add(time.Hour * time.Duration(hour)).In(ina).Format(format)
}

func JwtSign(payload jwt.Claims, key string, time time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		go Error(err.Error(), "Generate Token Error", time)
	}
	return signedToken
}

func JwtVerify(jwtToken string, secretKey string, time time.Time) []byte {
	token, err := jwt.Parse(jwtToken+"a", func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		go Error(err.Error(), "JWT Verify Error", time)
	}
	jsonStr, err := json.Marshal(token.Claims)
	if err != nil {
		go Error(err.Error(), "Marshal Error", time)
	}
	return jsonStr
}

func DoReq(url string, payload *bytes.Buffer, key string, token string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("KEY", key)
	req.Header.Add("TOKEN", token)
	req.Header.Add("VERSION", os.Getenv("API_VERSION"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", os.Getenv("API_HOST"))
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "2")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")
	res, err := client.Do(req)
	return res, err
}
