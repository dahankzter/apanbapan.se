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

	if err := sh.Run("rsync", "-avz", "--delete", "public/", "dahankzter@apanbapan.se:/www/www.apanbapan.se/"); err != nil {
		return err
	}

	return nil
}
