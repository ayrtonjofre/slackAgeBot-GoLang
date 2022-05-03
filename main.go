package main

import (
	"context"
	"fmt"     //utilizado para printar cosas
	"log"     //utilizado para la gestion de logs (errores)
	"os"      //utilizado para setear los environments
	"strconv" //utilizado para converter el year(parametro pasado) en string (este pkg no es nativo de GoLang)

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// Declaro los tokens generados en slack en estos environments y genero el nuevo cliente asignando los tokens seteados
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3498881890144-3468519504870-07sD2jHOz6vPkvVVvAayeMI0")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03EBNMFG57-3475141053891-71f3569b23de93e5b8e4d5c537e2aafa2066d630ed8fa034cbefe12a789bf055")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Commands events para siempre que pase un comando al bot sea la forma que lide el bot con el evento hara con que suceda un print out.
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")  // Seteo el parametro que pasamos, el pkg lo convertira en string
			yob, err := strconv.Atoi(year) //Si no lo convierto a string, la funcion no sera capaz de subtrair
			if err != nil {
				println("error")
			}
			age := 2022 - yob // Aca calculara el year(parametro) pasado y lo va subtraer al ano actual
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
