package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// For execution on FreePBX as flite
// https://github.com/FreePBX/tts/blob/release/17.0/agi-bin/propolys-tts.agi#L103C9-L103C14
// $enginebin." -f $textfile -o $tmpwavefile
func main() {
	wavOut := flag.String("o", "", "Output WAV file")
	textFile := flag.String("f", "", "Text file to read")
	flag.Parse()

	if *wavOut == "" || *textFile == "" {
		fmt.Println("Usage: go run script.go -o <output_wav> -f <text_file>")
		os.Exit(1)
	}

	text, err := ioutil.ReadFile(*textFile)
	if err != nil {
		log.Fatalf("error reading text file: %v", err)
		os.Exit(1)
	}

	//  --text TEXT           Text to generate speech.
	//  --out_path OUT_PATH   
	cmd := exec.Command("/usr/local/bin/tts", "--text", string(text), "--model_name", "tts_models/de/css10/vits-neon", "--out_path", *wavOut)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Cannot execute command: %v", err)
	}
}
