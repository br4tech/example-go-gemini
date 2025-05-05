package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_API_KEY") // Certifique-se de definir a variável de ambiente

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Println("Erro ao criar o cliente Gemini:", err)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	prompt := "Qual é a capital do Brasil?"

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		fmt.Println("Erro ao gerar conteúdo:", err)
		return
	}

	fmt.Println("Resposta do Gemini:")
	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Println(part)
	}
}
