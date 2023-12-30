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
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name: "source-lang",
                        Value: "english",
                        Aliases: []string{"sl"},
                        Usage: "language of source text",
                    },
                    &cli.StringFlag{
                        Name: "target-lang",
                        Value: "english",
                        Aliases: []string{"tl"},
                        Usage: "language to which translate",
                    },
                    &cli.StringFlag{
                        Name: "text",
                        Value: "Hello world!",
                        Aliases: []string{"t"},
                        Usage: "text to translate",
                    },
                    &cli.StringFlag{
                        Name: "model",
                        Value: "mistral",
                        Aliases: []string{"m"},
                        Usage: "model to use for translation",
                    },
                },
                Action: func(context *cli.Context) error {
                    model := context.String("model")
                    sourceLang := context.String("source-lang")
                    targetLang := context.String("target-lang")
                    text := context.String("text")
                    return translate(apiUrl, model, sourceLang, targetLang, text)
                },
            },
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
