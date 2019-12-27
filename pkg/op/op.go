package op

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Client for 1Password using https://1password.com/downloads/command-line/
type Client struct {
	session string
	vault   string
}

// NewClient returns a new op wrapper
func NewClient() *Client {
	return &Client{}
}

// SignIn https://support.1password.com/command-line/#sign-in-or-out
func (c *Client) SignIn(ctx context.Context, shorthand string) error {
	output, err := runOp(ctx, "signin", shorthand, "--output=raw")
	if err != nil {
		return err
	}

	c.session = strings.TrimSpace(string(output))

	return nil
}

// GetDocuments https://support.1password.com/command-line/#list-objects
func (c *Client) GetDocuments(ctx context.Context, vault string) ([]Document, error) {
	output, err := runOp(ctx, "list", "documents", fmt.Sprintf("--vault=%s", vault), fmt.Sprintf("--session=%s", c.session))
	if err != nil {
		return nil, err
	}

	var documents []Document
	if err := json.Unmarshal(output, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}

// GetItem
func (c *Client) GetItem(ctx context.Context, vault string, document string) (*Item, error) {
	output, err := runOp(ctx, "get", "item", document, fmt.Sprintf("--vault=%s", vault), fmt.Sprintf("--session=%s", c.session))
	if err != nil {
		return nil, err
	}

	var item Item
	if err := json.Unmarshal(output, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

// GetDocument https://support.1password.com/command-line/#work-with-documents
func (c *Client) GetDocument(ctx context.Context, vault string, document string, filename string) error {
	output, err := runOp(ctx, "get", "document", document, fmt.Sprintf("--vault=%s", vault), fmt.Sprintf("--session=%s", c.session))
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, output, 0644); err != nil {
		return err
	}

	return nil
}

// CreateDocument https://support.1password.com/command-line/#work-with-documents
func (c *Client) CreateDocument(ctx context.Context, vault string, filename string, overview Overview) error {
	_, err := runOp(ctx, "create", "document", filename, fmt.Sprintf("--title=%s", overview.Title), fmt.Sprintf("--tags=%s", strings.Join(overview.Tags, ",")), fmt.Sprintf("--vault=%s", vault), fmt.Sprintf("--session=%s", c.session))
	if err != nil {
		return err
	}

	return nil
}

func runOp(ctx context.Context, args ...string) ([]byte, error) {
	opCmd := exec.CommandContext(ctx, "op", args...)
	opCmd.Stdin = os.Stdin
	output, err := opCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%w; %s", err, string(output))
	}

	return output, nil
}
