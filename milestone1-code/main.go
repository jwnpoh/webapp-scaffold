package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type opts struct {
	nameOfProject  string
	locationOnDisk string
	remoteURL      string
	website        bool
}

func main() {
	_, err := setupFlags(os.Stderr, os.Args[1:])
	if err != nil {
	}
}

func setupFlags(w io.Writer, args []string) (opts, error) {
	var projectName, locationOnDisk, remoteURL string
	var s bool
	var o opts

	fs := flag.NewFlagSet("opts", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.BoolVar(&o.website, "s", s, "Setup project for a website along with the HTTP backend API?")
	fs.StringVar(&o.nameOfProject, "n", projectName, "Name of the project")
	fs.StringVar(&o.locationOnDisk, "d", locationOnDisk, "Location of the project on disk")
	fs.StringVar(&o.remoteURL, "r", remoteURL, "URL of remote repository")
	err := fs.Parse(args)
	if err != nil {
		return o, errors.New("project not set up correctly. Try -h for flag usage")
	}

	err = validateOpts(o)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return o, err
	}

	fmt.Fprintf(
		w,
		"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
		o.nameOfProject,
		o.locationOnDisk,
		o.remoteURL,
		o.website,
	)
	return o, nil
}

func validateOpts(o opts) error {
	var errMsg string

	if o.nameOfProject == "" {
		errMsg += "Project name cannot be unspecified.\n"
	}

	if o.locationOnDisk == "" {
		errMsg += "Project path cannot be unspecified.\n"
	}

	if o.remoteURL == "" {
		errMsg += "Project repository URL cannot be unspecified.\n"
	}

	if errMsg != "" {
		err := errors.New(errMsg)
		return err
	}

	return nil
}
