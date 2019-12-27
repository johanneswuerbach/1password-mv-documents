package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/johanneswuerbach/1password-mv-documents/pkg/op"
)

func main() {
	var originShorthand, targetShorthand, originVault, targetVault string

	flag.StringVar(&originShorthand, "origin-shorthand", "", "Origin shorthand (required)")
	flag.StringVar(&targetShorthand, "target-shorthand", "", "Target shorthand (required)")

	flag.StringVar(&originVault, "origin-vault", "", "Origin vault by UUID or name (required)")
	flag.StringVar(&targetVault, "target-vault", "", "Target vault by UUID or name (required)")

	flag.Parse()

	if originShorthand == "" {
		flag.Usage()
		panic(errors.New("origin shorthand is required"))
	}

	if targetShorthand == "" {
		flag.Usage()
		panic(errors.New("target shorthand is required"))
	}

	if originVault == "" {
		flag.Usage()
		panic(errors.New("origin vault is required"))
	}

	if targetVault == "" {
		flag.Usage()
		panic(errors.New("target vault is required"))
	}

	originClient := op.NewClient()
	targetClient := op.NewClient()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		cancel()
	}()

	if err := originClient.SignIn(ctx, originShorthand); err != nil {
		panic(err)
	}

	if err := targetClient.SignIn(ctx, targetShorthand); err != nil {
		panic(err)
	}

	documents, err := originClient.GetDocuments(ctx, originVault)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copying %d document(s)\n", len(documents))

	for _, document := range documents {
		item, err := originClient.GetItem(ctx, originVault, document.UUID)
		if err != nil {
			panic(fmt.Errorf("%w; error getting document details %s", err, document.Overview.Title))
		}

		fileName := path.Join("documents", item.Details.DocumentAttributes.FileName)

		if err := originClient.GetDocument(ctx, originVault, document.UUID, fileName); err != nil {
			panic(fmt.Errorf("%w; error getting document %s", err, document.Overview.Title))
		}

		if err := targetClient.CreateDocument(ctx, targetVault, fileName, document.Overview); err != nil {
			panic(fmt.Errorf("%w; error creating document %s", err, document.Overview.Title))
		}
	}

	fmt.Printf("Documents copied\n")
}
