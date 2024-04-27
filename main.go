package main

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)
import "google.golang.org/api/option"

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("너의 이름은?"))
	if err != nil {
		log.Fatal(err)
	}
	PrintModelResp(resp)

	resp, err = model.GenerateContent(ctx, genai.Text("Gemini의 대단함을 칭송하는 3줄의 짧은 시를 작성해줘"))
	if err != nil {
		log.Fatal(err)
	}
	PrintModelResp(resp)

	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("안녕 나는 올에 일본, 코타키나발루 여행을 다녀왔어 다른 휴양지도 가보고 싶어"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("만나서 반갑습니다, 저는 여행 전문 가이드입니다."),
			},
			Role: "model",
		},
	}
	resp, err = cs.SendMessage(ctx, genai.Text("내가 가보지 않고 좋아할만한 여행지를 추천해줘"))
	if err != nil {
		log.Fatal(err)
	}
	PrintModelResp(resp)
}

func PrintModelResp(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}
