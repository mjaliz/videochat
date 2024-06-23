package server

import (
	"flag"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"time"
	"videochat/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

var (
	addr = flag.String("addr", fmt.Sprintf(":%s", os.Getenv("PORT")), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()
	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket")
	app.Get("/room/:uuid/chat", hanlders.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	}))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/:ssuid/chat/websocket")
	app.Get("/stream/:ssuid/viewer/websocket")
}
