package pkg

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Login(username, password string) error {
	url := fmt.Sprintf("%s%s%s?username=%s&password=%s", c.hostname, c.apiVer, "/auth/login", username, password)
	c.logger.Println(url)
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		c.logger.Printf("Failed to create request: %s", err)
		return err
	}
	req.Header.Add("Referer", c.hostname)

	resp, err := c.httpClient.Do(req)
	if err != nil || resp.Status != "200 OK" {
		c.logger.Printf("Failed to send request: %s", err)
		return err
	}

	var respbody string
	resp.Body.Read([]byte(respbody))
	c.logger.Println(respbody)

	var token string
	for _, cookie := range resp.Cookies() {
		c.logger.Printf("Cookie %s", cookie.String())
		if cookie.Name == "SID" {
			token = cookie.Value
			break
		}
		return errors.New("Failed to find cookie")
	}

	c.token = token

	return nil
}

func (c *Client) Logout() error {
	url := fmt.Sprintf("%s%s%s", c.hostname, c.apiVer, "/auth/logout")
	c.logger.Println(url)

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		c.logger.Printf("Failed to create logout request")
		return err
	}

	req.AddCookie(&http.Cookie{Name: "SID", Value: c.token})

	resp, err := c.httpClient.Do(req)
	if err != nil || resp.Status != "200 OK" {
		c.logger.Printf("Failed to send logout request")
		return err
	}

	return nil
}
