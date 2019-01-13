package main

import (
	"github.com/containers/libpod/libpod/adapter"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	tagDescription = "Adds one or more additional names to locally-stored image"
	tagCommand     = cli.Command{
		Name:         "tag",
		Usage:        "Add an additional name to a local image",
		Description:  tagDescription,
		Action:       tagCmd,
		ArgsUsage:    "IMAGE-NAME [IMAGE-NAME ...]",
		OnUsageError: usageErrorHandler,
	}
)

func tagCmd(c *cli.Context) error {
	args := c.Args()
	if len(args) < 2 {
		return errors.Errorf("image name and at least one new name must be specified")
	}
	localRuntime, err := adapter.GetRuntime(c)
	if err != nil {
		return errors.Wrapf(err, "could not create runtime")
	}
	defer localRuntime.Runtime.Shutdown(false)

	newImage, err := localRuntime.NewImageFromLocal(args[0])
	if err != nil {
		return err
	}

	for _, tagName := range args[1:] {
		if err := newImage.TagImage(tagName); err != nil {
			return errors.Wrapf(err, "error adding '%s' to image %q", tagName, newImage.InputName)
		}
	}
	return nil
}
