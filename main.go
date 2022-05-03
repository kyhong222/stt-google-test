// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	speech "cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
        ctx := context.Background()

        // Creates a client.
        client, err := speech.NewClient(ctx, option.WithCredentialsFile("./defaultConfig.json"))
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }
        defer client.Close()

        // The path to the remote audio file to transcribe.
        fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"
        
        // Detects speech in the audio file.
        t1:= time.Now()
        fmt.Println(time.Now(), "speech sent")
        resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
                Config: &speechpb.RecognitionConfig{
                        Encoding:        speechpb.RecognitionConfig_LINEAR16,
                        SampleRateHertz: 16000,
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
                        AudioSource: &speechpb.RecognitionAudio_Uri{Uri: fileURI},
                },
        })
        if err != nil {
                log.Fatalf("failed to recognize: %v", err)
        }
        t2:=time.Since(t1)
        fmt.Println(time.Now(), "text received")
        // Prints the results.
        for _, result := range resp.Results {
                for _, alt := range result.Alternatives {
                        fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }
        fmt.Println("time diff:", t2)

}