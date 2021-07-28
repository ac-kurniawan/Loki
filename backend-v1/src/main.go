package main

import (
	apiServer "antriin/src/app"
	"antriin/src/business/attendee"
	"antriin/src/business/event"
	"antriin/src/modules/AMQ"
	AMQ2 "antriin/src/modules/AMQ/attendee"
	AMQ3 "antriin/src/modules/AMQ/event"
	attendeeController "antriin/src/modules/apiControllerV1/anttendee"
	authController "antriin/src/modules/apiControllerV1/auth"
	eventController "antriin/src/modules/apiControllerV1/event"
	keycloak "antriin/src/modules/keycloak/auth"
	"antriin/src/modules/mongoRepo"
	AttendeeMongoRepository "antriin/src/modules/mongoRepo/attendee"
	EventMongoRepository "antriin/src/modules/mongoRepo/event"
	"antriin/src/util/debug"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

const Banner = `
░█████╗░███╗░░██╗████████╗██████╗░██╗░░░██╗███╗░░██╗  ░░░░░░  ░█████╗░██████╗░██╗
██╔══██╗████╗░██║╚══██╔══╝██╔══██╗██║░░░██║████╗░██║  ░░░░░░  ██╔══██╗██╔══██╗██║
███████║██╔██╗██║░░░██║░░░██████╔╝██║░░░██║██╔██╗██║  █████╗  ███████║██████╔╝██║
██╔══██║██║╚████║░░░██║░░░██╔══██╗██║░░░██║██║╚████║  ╚════╝  ██╔══██║██╔═══╝░██║
██║░░██║██║░╚███║░░░██║░░░██║░░██║██║██╗██║██║░╚███║  ░░░░░░  ██║░░██║██║░░░░░██║
╚═╝░░╚═╝╚═╝░░╚══╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝╚═╝╚═╝╚═╝░░╚══╝  ░░░░░░  ╚═╝░░╚═╝╚═╝░░░░░╚═╝

██╗░░░██╗░░███╗░░
██║░░░██║░████║░░
╚██╗░██╔╝██╔██║░░
░╚████╔╝░╚═╝██║░░
░░╚██╔╝░░███████╗
░░░╚═╝░░░╚══════╝`

func main() {
	fmt.Printf("%s\n", Banner)

	// load env
	err := godotenv.Load("./src/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// define driver
	kc := keycloak.NewKeycloak(os.Getenv("KEYCLOAK_HOST"),
		os.Getenv("KEYCLOAK_REALM_NAME"), os.Getenv("KEYCLOAK_CLIENT_ID"),
		os.Getenv("KEYCLOAK_CLIENT_SECRET"))

	mongoConnection, err := mongoRepo.Connect()
	if err != nil {
		debug.Error("MongoConnection", err.Error())
	}

	pub, sub, err := AMQ.NewAMQP(os.Getenv("AMQ_SERVER"))
	if err != nil {
		debug.Error("AMQConnection", err.Error())
	}

	// Define infrastructure layer
	attendeePublisher := AMQ2.NewAttendeePublisher(pub)
	attendeeRepository := AttendeeMongoRepository.NewAttendeeRepository(mongoConnection)
	eventRepository := EventMongoRepository.NewEventRepository(mongoConnection)

	// Define service layer
	attendeeService := attendee.NewAttendee(attendeeRepository, attendeePublisher)
	eventService := event.NewEvent(eventRepository)

	// define echo instance
	e := echo.New()

	// Define presentation layer
	eventController.EventController(e, eventService, kc)
	attendeeController.AttendeeController(e, attendeeService, kc)
	authController.AuthController(e, kc)

	// event Subscriber
	eventRouter := AMQ3.NewEventConsumer(eventService, sub)
	eventRouter.OnAttendeeCreated()


	// run http server
	apiServer.RunHttpServer(e)
}