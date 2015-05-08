/*
 * Copyright (C) 2014-2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"fmt"

	"launchpad.net/snappy/logger"
	"launchpad.net/snappy/snappy"
)

type cmdInternalFirstBootOemConfig struct{}

func init() {
	_, err := parser.AddCommand("firstboot",
		"internal",
		"internal",
		&cmdInternalFirstBootOemConfig{})
	if err != nil {
		logger.Panic("unable to first_boot: %v", err)
	}
}

func (x *cmdInternalFirstBootOemConfig) Execute(args []string) error {
	err := snappy.OemConfig()
	if err == snappy.ErrNotFirstBoot {
		fmt.Println("First boot has already run")
		return nil
	}

	return err
}
