// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.


package base

import (
	"github.com/urfave/cli"
	"github.com/dipperin/dipperin-core/core/chain-config"
)

func NewApp(appName, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Author = chain_config.AppName
	app.Email = ""
	app.Version = chain_config.Version
	app.Usage = usage

	return app
}
