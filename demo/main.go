package main

import (
	"context"
	"log"
	"time"

	"github.com/themakers/mactts"
)

func main() {
	nvoices, err := mactts.NumVoices()
	if err != nil {
		panic(err)
	}

	log.Println("num voices", nvoices)

	var (
		voice *mactts.VoiceSpec
		text  string
	)

	for n := 1; n <= nvoices; n++ {
		v, err := mactts.GetVoice(n)
		if err != nil {
			panic(err)
		}

		vd, err := v.Description()
		if err != nil {
			panic(err)
		}

		log.Println("voice# ", vd.Name(), vd.Age())

		voice = v
		switch vd.Name() {
		case "Samantha":
			text = "Hi Corey, I am Samantha! I can speak your language better than these assholes around you! Ha ha ha!"
			//n = nvoices
		case "Zarvox":
			text = "Everything that lives is designed to end."
			text = "Everything that leaves is designed to end."
			n = nvoices
		case "Albert":
			text = "I have a frog in my fruit!"
			//n = nvoices
		}
	}

	chn, err := mactts.NewChannel(voice)
	if err != nil {
		panic(err)
	}

	if err := chn.SpeakString(text); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := chn.SetDone(func() {
		log.Println("done")
		cancel()
	}); err != nil {
		panic(err)
	}

	select {
	case <-ctx.Done():
		log.Println("really done")
	case <-time.After(10 * time.Second):
		log.Println("timeout")
	}
}
