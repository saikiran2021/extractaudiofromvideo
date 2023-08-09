package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/extract-audio", ExtractAudio)

	log.Fatal(app.Listen(":3000"))
}

func ExtractAudio(c *fiber.Ctx) error {

	inputFilePath := c.Query("input")
	fmt.Println("inputFilePath:", inputFilePath)
	outputFilePath := "output_audio.mp3" // Change this to your desired output file name

	cmd := exec.Command("ffmpeg", "-i", inputFilePath, "-vn", "-acodec", "libmp3lame", outputFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Error extracting audio: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error extracting audio")
	}

	return c.SendString("Audio extracted successfully")
}
