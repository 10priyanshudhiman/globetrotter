package controller

type AppController struct {
	UserRegistry interface{ UserRegistryController }
	Game         interface{ GameController }
}
