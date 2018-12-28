package codefresh

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
)

type (
	Codefresh interface {
		requestAPI(*requestOptions) (*gentleman.Response, error)
		ITokenAPI
		IPipelineAPI
	}
)

func New(opt *ClietOptions) Codefresh {
	client := gentleman.New()
	client.URL(opt.Host)
	return &codefresh{
		token:  opt.Auth.Token,
		client: client,
	}
}

func (c *codefresh) requestAPI(opt *requestOptions) (*gentleman.Response, error) {
	req := c.client.Request()
	req.Path(opt.path)
	req.Method(opt.method)
	req.AddHeader("Authorization", c.token)
	if opt.body != nil {
		req.Use(body.JSON(opt.body))
	}
	if opt.qs != nil {
		for k, v := range opt.qs {
			req.Use(query.Set(k, v))
		}
	}
	res, _ := req.Send()
	if res.StatusCode > 400 {
		return res, fmt.Errorf("Error occured during API invocation\nError: %s", res.String())
	}
	return res, nil
}
