package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func loginPost(c *gin.Context) {
	password := c.PostForm("password")
	redirectURL := c.PostForm("originalURL")
	if redirectURL == "" {
		redirectURL = "/"
	}
	if password == configInstance.Password {
		c.SetCookie(configInstance.CookieName, configInstance.CookieValue, configInstance.CookieTTL, "/", "", false, true)
	}
	c.Redirect(http.StatusFound, redirectURL)
}

func proxyRequest(c *gin.Context) {
	auth, err := c.Cookie(configInstance.CookieName)
	if err != nil || auth != configInstance.CookieValue {
		// c.Redirect(http.StatusFound, configInstance.Prefix+"/login")
		originalURL := c.Request.URL.RequestURI()
		data := `<html>
			<head>
				<title>Login</title>
			</head>
			<body>
				<form action="` + configInstance.Prefix + `/login" method="post">
					<input type="hidden" name="originalURL" value="` + originalURL + `"/>
					<input type="password" name="password" placeholder="Enter your password"/>
					<button type="submit">Login</button>
				</form>
			</body>
		</html>`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, data)
		return
	}

	fullURL := configInstance.TargetURL + c.Request.URL.RequestURI()

	client := &http.Client{}
	req, err := http.NewRequest(c.Request.Method, fullURL, c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	req.Header = make(http.Header)
	for key, value := range c.Request.Header {
		req.Header[key] = value
	}
	req.Header.Del("Host")

	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// set the response status and headers
	c.Status(resp.StatusCode)
	for key, value := range resp.Header {
		for _, v := range value {
			c.Header(key, v)
		}
	}

	// stream the response body
	reader := bufio.NewReader(resp.Body)
	buf := make([]byte, 128)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if n == 0 {
			break
		}
		c.Writer.Write(buf[:n])
	}

	defer resp.Body.Close()
}
