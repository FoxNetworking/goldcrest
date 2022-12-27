package main

import (
	tempest "github.com/Amatsagu/Tempest"
	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
	"log"
	"os"
)

// ensureEnvVar panics if an environmental variable is unavailable.
func ensureEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Variable %s was not present!", key)
	}
	return value
}

func main() {
	// First, we need to determine the port specified by App Engine.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Next, we use the Discord bot information provided via Berglas.
	token := "Bot " + ensureEnvVar("DISCORD_TOKEN")
	appId := ensureEnvVar("DISCORD_BOT_ID")
	publicKey := ensureEnvVar("DISCORD_PUBLIC_KEY")

	// Finally, we can start.
	client := tempest.CreateClient(tempest.ClientOptions{
		ApplicationId: tempest.StringToSnowflake(appId),
		PublicKey:     publicKey,
		Token:         token,
		PreCommandExecutionHandler: func(itx tempest.CommandInteraction) *tempest.ResponseData {
			log.Println("Processing command: " + itx.Data.Name)
			return nil
		},
	})

	client.RegisterCommand(colorCmd)
	client.SyncCommands(nil, nil, false)

	log.Printf("Listening to :%s", port)
	if err := client.ListenAndServe("/", ":"+port); err != nil {
		log.Fatal(err)
	}
}
