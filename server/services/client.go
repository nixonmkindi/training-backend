package services

import (
	"training-backend/package/client"
	"training-backend/package/config"
	"training-backend/package/log"
	"training-backend/package/util"
)

const (
	AuthBaseUrl = "http://localhost:4333/auth/api/v1"
	SomaBaseUrl = "http://localhost:4333/soma/api/v1"
)

var AuthClient *client.Client
var SomaClient *client.Client

func InitBackendClients(cfg *config.Config) {

	trainingBackendSystem := "training_backend"
	trainingBackendKey, err := cfg.GetSystemPrivateKey(trainingBackendSystem)
	if util.CheckError(err) {
		log.Errorf("error getting training backend private key: %v", err)
		panic("training backend system is not well started")
	}

	SomaClient, err = client.New(SomaBaseUrl, trainingBackendKey, trainingBackendSystem)
	if util.CheckError(err) {
		log.Errorf("error initiating the soma public key: %v", err)

		panic("soma system client could not be initiated")
	}
}
