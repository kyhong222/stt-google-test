// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	speech "cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
        ctx := context.Background()

        // audio config input from json
        audioConfig := make(map[string]int32)
        audioConfigPath := "./audioConfig.json"
        jsonFile, _ := os.Open(audioConfigPath)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
        json.Unmarshal([]byte(byteValue), &audioConfig)

        // Creates a client.
        googleCredentialPath := "./googleCredential.json"
        client, err := speech.NewClient(ctx, option.WithCredentialsFile(googleCredentialPath))
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
                        SampleRateHertz: (audioConfig["SampleRateHertz"]),
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
                        AudioSource: &speechpb.RecognitionAudio_Uri{Uri: fileURI},
                },
        })
        if err != nil {
                log.Fatalf("failed to recognize: %v", err)
        }
        fmt.Println(time.Now(), "text received")
        // Prints the results.
        for _, result := range resp.Results {
                for _, alt := range result.Alternatives {
                        fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }
        t2:=time.Since(t1)
        fmt.Println("time diff:", t2)

}