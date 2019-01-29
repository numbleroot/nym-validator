// main_sample.go - sample usage for tendermint client
// Copyright (C) 2019  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"

	"0xacab.org/jstuczyn/CoconutGo/constants"

	"0xacab.org/jstuczyn/CoconutGo/crypto/bpgroup"
	"0xacab.org/jstuczyn/CoconutGo/logger"
	"0xacab.org/jstuczyn/CoconutGo/tendermint/client"
	"0xacab.org/jstuczyn/CoconutGo/tendermint/nymabci/transaction"
	Curve "github.com/jstuczyn/amcl/version3/go/amcl/BLS381"
)

// currently used entirely for debug purposes
func mainOld() {

	log, err := logger.New("", "DEBUG", false)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a logger: %v", err))
	}

	client, err := client.New("tcp://0.0.0.0:46667", log)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a client: %v", err))
	}

	bpgroup := bpgroup.New()

	z1 := Curve.G1mul(bpgroup.Gen1(), Curve.Randomnum(bpgroup.Order(), bpgroup.Rng()))
	z2 := Curve.G1mul(bpgroup.Gen1(), Curve.Randomnum(bpgroup.Order(), bpgroup.Rng()))

	// basically lazy way to convert to bytes
	t := transaction.NewLookUpZetaTx(z1)
	client.Broadcast(t[1:])

	// shouldn be
	isPresent1 := client.LookUpZeta(z1)

	// shouldn't be
	isPresent2 := client.LookUpZeta(z2)

	fmt.Println(isPresent1, isPresent2)

	client.SendAsync(t)
	client.SendAsync(t[1:])
	client.SendAsync(t[2:])

	// error
	res, err := client.Broadcast([]byte{'a'})

	if err != nil {
		fmt.Printf("Error response: %v", err)
	}
	fmt.Println(res)

	client.Stop()
}

func main() {
	// priv := ed25519.GenPrivKey()
	// pub := priv.PubKey()
	// fmt.Println(priv)
	// fmt.Println(priv.Bytes())
	// fmt.Println(pub)
	// fmt.Println(pub.Bytes())

	bpgroup := bpgroup.New()

	z1 := Curve.G1mul(bpgroup.Gen1(), Curve.Randomnum(bpgroup.Order(), bpgroup.Rng()))

	b1 := make([]byte, constants.ECPLen)
	b2 := make([]byte, constants.ECPLenUC)

	z1.ToBytes(b1, true)
	z1.ToBytes(b2, false)

	fmt.Println(b1)

	fmt.Println(b2)
}
