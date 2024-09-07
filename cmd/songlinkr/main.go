package main

import (
	"log/slog"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	services "github.com/ChausseBenjamin/songlinkr/internal/service"
	"github.com/ChausseBenjamin/songlinkr/internal/urls"
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

func AppAction(ctx *cli.Context) error {
	token, err := getSecret(ctx.String("secrets-path"), "token")
	if err != nil {
		return err
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		slog.Error("Unable to launch slog", "error", err)
		os.Exit(1)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		services := services.GetServices()

		links := urls.Find(m.Content)
		for _, link := range links {
			for _, srvc := range services {
				if srvc.Owns(link) {

					link = srvc.Resolve(link)
					src, err := urls.Resolve(link)
					if err != nil {
						src = link // An error occured, use the less elegant link...
					}

					msg := srvc.Name() + " link detected!\n"
					msg += "[Here is a universal link for everyone to enjoy](" + src + ")! 🎶\n"
					msg += "Beep boop! Have a nice one! 🤖"
					s.ChannelMessageSend(m.ChannelID, msg)
					break
				}
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		return err
	}
	defer sess.Close()

	slog.Info("Bot is now running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	return nil
}

func getSecret(secretsPath, secretName string) (string, error) {
	secretValue, err := os.ReadFile(path.Join(secretsPath, secretName))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(secretValue)), nil
}

func main() {
	loggr := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(loggr)

	app := &cli.App{
		Name:   "Songlinkr",
		Usage:  "A Discord bot that converts song links to Universal Song.link",
		Action: AppAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "secrets-path",
				EnvVars: []string{"SECRETS_PATH"},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		slog.Error("An error occured during runtime", "error", err)
		os.Exit(1)
	}
}
