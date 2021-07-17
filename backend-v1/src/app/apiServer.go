package apiServer

import (
	"antriin/src/business/attendee"
	"antriin/src/business/event"
	attendeeController "antriin/src/modules/apiControllerV1/anttendee"
	authController "antriin/src/modules/apiControllerV1/auth"
	eventController "antriin/src/modules/apiControllerV1/event"
	keycloak "antriin/src/modules/keycloak/auth"
	"antriin/src/modules/mongoRepo"
	AttendeeMongoRepository "antriin/src/modules/mongoRepo/attendee"
	EventMongoRepository "antriin/src/modules/mongoRepo/event"
	"antriin/src/util/debug"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func httpServer(e *echo.Echo, address string, mongoConnection *mongo.Client) {
	// Define infrastructure layer
	kc := keycloak.NewKeycloak(os.Getenv("KEYCLOAK_HOST"),
		os.Getenv("KEYCLOAK_REALM_NAME"), os.Getenv("KEYCLOAK_CLIENT_ID"),
		os.Getenv("KEYCLOAK_CLIENT_SECRET"))
	attendeeRepository := AttendeeMongoRepository.NewAttendeeRepository(mongoConnection)
	eventRepository := EventMongoRepository.NewEventRepository(mongoConnection)

	// Define service layer
	attendeeService := attendee.NewAttendee(attendeeRepository)
	eventService := event.NewEvent(eventRepository)

	// Define presentation layer
	eventController.EventController(e, eventService, kc)
	attendeeController.AttendeeController(e, attendeeService, kc)
	authController.AuthController(e, kc)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})


	// Start HTTP Server
	e.HideBanner = true
	go func() {
		if err := e.Start(address); err != nil {
			fmt.Printf("[HTTP Server] - %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		debug.Error("MongoConnection", err.Error())
	}
}

func RunHttpServer() {
	// init instance and configuration
	e := echo.New()
	address := fmt.Sprintf("0.0.0.0:8030")

	mongoConnection,err := mongoRepo.Connect()
	if err != nil {
		debug.Error("MongoConnection", err.Error())
	}

	// start server
	httpServer(e, address, mongoConnection)

}
