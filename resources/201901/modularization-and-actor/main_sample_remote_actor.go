package main

import (
	"sampleactor/actor"
	"context"
	"sampleactor/samplestruct"
)

func main() {

	ctxMain, _ := context.WithCancel(context.Background())

	actor.StartWebServer(ctxMain,"9999")

	training := samplestruct.NewTraining("Training")

	_ = actor.StartActor(training)

	for {
		select {
		case <-ctxMain.Done():
		}
	}
}