// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() error {
	mg.Deps(Clean)
	return sh.Run("hugo")
}

func Clean() error {
	return sh.Rm("public")
}

func Deploy() error {
	mg.Deps(Build)

	if err := sh.Run("scp", "-r", "public", "dahankzter@apanbapan.se:."); err != nil {
		return err
	}

	if err := sh.Run("ssh", "dahankzter@apanbapan.se", "cp ", "-r", "public/*", "/www/www.apanbapan.se/."); err != nil {
		return err
	}

	return nil
}
