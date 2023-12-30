package main

import (
    "os"
    "fmt"
    "log"
    "bytes"
    "bufio"
    "net/http"
    "encoding/json"
    "github.com/urfave/cli/v2"
)

const apiUrl = "http://localhost:11434/api"
const model = "mistral"

type GenerationRequest struct {
    Model       string
    Prompt      string
}

type GenerationResponse struct {
    Model       string
    Response    string
    Done        bool
}

func translate(apiUrl, model, sourceLang string, targetLang string, text string) error {

    prompt := fmt.Sprintf("Translate text from %s to %s:\n%s\n", sourceLang, targetLang, text)
    generationRequest := GenerationRequest{Model: model, Prompt: prompt}

    data_json, err := json.Marshal(generationRequest)
    if err != nil {
        return err
    }

    resp, err := http.Post(apiUrl + "/generate/", "application/json", bytes.NewBuffer(data_json))
    if err != nil {
        return err
    }

    // read generation stream
    reader := bufio.NewReader(resp.Body)
    for {
        line, err := reader.ReadBytes('\n')
        if err != nil {
            break;
        }

        var resp_data GenerationResponse
        json.Unmarshal(line, &resp_data)

        if resp_data.Done == true {
            break;
        }

        fmt.Print(resp_data.Response)
    }
    defer resp.Body.Close()

    return nil
}

func main() {
    app := cli.App{
        Commands: []*cli.Command{
            {
                Name: "translate",
                Aliases: []string{"t"},
                Usage: "translate text from one language to another",
                Action: func(context *cli.Context) error {
                    return translate(apiUrl, model, "english", "french", "Hello world!")
                },
            },
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
