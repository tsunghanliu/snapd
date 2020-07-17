// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2020 Canonical Ltd
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

package assets

type ForEditions = forEditions

var (
	RegisterInternal           = registerInternal
	RegisterSnippetForEditions = registerSnippetForEditions
)

func MockCleanState() (restore func()) {
	oldRegisteredAssets := registeredAssets
	oldRegisteredEditionAssets := registeredEditionAssets
	registeredAssets = map[string][]byte{}
	registeredEditionAssets = map[string][]forEditions{}
	return func() {
		registeredAssets = oldRegisteredAssets
		registeredEditionAssets = oldRegisteredEditionAssets
	}
}
