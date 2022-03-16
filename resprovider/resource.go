package resprovider

import "github.com/sonkt1210/tiki-vas/logger"

type Logger logger.Logger

type ResourceProvider interface {

	// GetAppRoot() string

	Logger(prefix string) Logger

	Flogger(prefix string) Logger

	// GetConfig() *appConfService

	// GetServiceHost(string) struct {
	// 	Host string
	// }

	// //DB() *gorm.DB

	// MgoDb() *mgo.Session
}

var rp ResourceProvider

// func InitResourceProvider(app sdms.Application) {
// 	if rp != nil {
// 		app := rp.(*myResourceProvider).app
// 		if !app.IsShutdown() {
// 			log := rp.Logger("resprovider")
// 			log.Fatal("InitResourceProvider can not called twice!")
// 		}
// 	}
// 	rp = newRP(app)
// }
