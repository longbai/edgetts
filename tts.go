package edgetts

import (
	"log"
	"os"
)

type TTS interface {
	NewConn() error
	CloseConn()
	GetAudio(ssml, format string) (audioData []byte, err error)
	GetAudioStream(ssml, format string, read func([]byte)) error
}

func TextToMp3(tts TTS, text string, ttsLang string, filePath string) error {
	ssml := CreateSSML(text, ttsLang)
	log.Println(ssml)
	b, err := tts.GetAudio(ssml, voiceFormat)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}
	err = os.WriteFile(filePath, b, 0644)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}
	return nil
}
