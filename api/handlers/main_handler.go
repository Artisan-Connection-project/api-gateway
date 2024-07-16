package handlers

import "github.com/sirupsen/logrus"

type MainHandler interface {
	Auth() AuthHandler
	Product() ProductHandler
}

type mainHandler struct {
	auth    AuthHandler
	product ProductHandler
	log     *logrus.Logger
}

func (m *mainHandler) Auth() AuthHandler {
	return m.auth
}

func (m *mainHandler) Product() ProductHandler {
	return m.product
}

func NewMainHandler(auth AuthHandler, product ProductHandler, log *logrus.Logger) MainHandler {
	return &mainHandler{auth: auth, product: product, log: log}
}
