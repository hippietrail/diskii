// Copyright © 2016 Zellyn Hunter <zellyn@gmail.com>

package cmd

import (
	"fmt"

	"github.com/zellyn/diskii/disk"
	"github.com/zellyn/diskii/helpers"
	"github.com/zellyn/diskii/types"
)

type PutCmd struct {
	Order        types.DiskOrder `kong:"default='auto',enum='auto,do,po',help='Logical-to-physical sector order.'"`
	System       string          `kong:"default='auto',enum='auto,dos3',help='DOS system used for image.'"`
	FiletypeName string          `kong:"default='B',help='Type of file (“diskii filetypes” to list).'"`
	Overwrite    bool            `kong:"short='f',help='Overwrite existing file?'"`

	DiskImage      string `kong:"arg,required,type='existingfile',help='Disk image to modify.'"`
	TargetFilename string `kong:"arg,required,help='Filename to use on disk.'"`
	SourceFilename string `kong:"arg,required,type='existingfile',help='Name of file containing data to put.'"`
}

func (p PutCmd) Help() string {
	return `Examples:
	# Put file gremlins.o onto disk image games.dsk, using the filename GREMLINS.
	diskii put games.dsk GREMLINS gremlins.o`
}

func (p *PutCmd) Run(globals *types.Globals) error {
	if p.DiskImage == "-" {
		if p.SourceFilename == "-" {
			return fmt.Errorf("cannot get both disk image and file contents from stdin")
		}
	}
	op, order, err := disk.OpenFilename(p.DiskImage, p.Order, p.System, globals.DiskOperatorFactories, globals.Debug)
	if err != nil {
		return err
	}

	contents, err := helpers.FileContentsOrStdIn(p.SourceFilename)
	if err != nil {
		return err
	}

	filetype, err := types.FiletypeForName(p.FiletypeName)
	if err != nil {
		return err
	}
	fileInfo := types.FileInfo{
		Descriptor: types.Descriptor{
			Name:   p.TargetFilename,
			Length: len(contents),
			Type:   filetype,
		},
		Data: contents,
	}
	_, err = op.PutFile(fileInfo, p.Overwrite)
	if err != nil {
		return err
	}

	return disk.WriteBack(p.DiskImage, op, order, true)
}
