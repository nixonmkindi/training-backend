package services

import (
	"training-backend/package/client"
	"training-backend/package/config"
	"training-backend/package/log"
	"training-backend/package/util"
)

const (
	AuthBaseUrl = "http://localhost:4322/auth/api/v1"
	SomaBaseUrl = "http://localhost:4325/soma/api/v1"
)

var AuthClient *client.Client
var SomaClient *client.Client

func InitBackendClients(cfg *config.Config) {

	efileSystem := "notification"
	efileKey, err := cfg.GetSystemPrivateKey(efileSystem)
	if util.CheckError(err) {
		log.Errorf("error getting notification private key: %v", err)
		panic("notification system is not well started")
	}

	SomaClient, err = client.New(SomaBaseUrl, efileKey, efileSystem)
	if util.CheckError(err) {
		log.Errorf("error initiating the soma public key: %v", err)

		panic("soma system client could not be initiated")
	}
}
