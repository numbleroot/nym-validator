// main_sample.go - sample usage for coconut/tendermint client
// Copyright (C) 2018-2019  Jedrzej Stuczynski.
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
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"0xacab.org/jstuczyn/CoconutGo/common/utils"

	cclient "0xacab.org/jstuczyn/CoconutGo/client"
	"0xacab.org/jstuczyn/CoconutGo/client/config"
	"0xacab.org/jstuczyn/CoconutGo/crypto/bpgroup"
	"0xacab.org/jstuczyn/CoconutGo/crypto/coconut/scheme"
	"0xacab.org/jstuczyn/CoconutGo/logger"
	"0xacab.org/jstuczyn/CoconutGo/nym/token"
	"0xacab.org/jstuczyn/CoconutGo/tendermint/account"
	tmclient "0xacab.org/jstuczyn/CoconutGo/tendermint/client"
	"0xacab.org/jstuczyn/CoconutGo/tendermint/nymabci/code"
	"0xacab.org/jstuczyn/CoconutGo/tendermint/nymabci/transaction"
	Curve "github.com/jstuczyn/amcl/version3/go/amcl/BLS381"
)

const providerAddress = "127.0.0.1:4000"
const providerAddressGrpc = "127.0.0.1:5000"

const providerAcc = "AwYXtM4pa4WV47TozIi1gf6t/jdRQyQkPv6mAC0S/fyzdPP4Pr3DAtOP0h0BYcHQIQ=="

var tendermintABCIAddresses = []string{
	// "tcp://0.0.0.0:12345", // does not exist
	"tcp://0.0.0.0:26657",
	"tcp://0.0.0.0:26659",
	"tcp://0.0.0.0:26661",
	"tcp://0.0.0.0:26663",
}

// const tendermintABCIAddress = "tcp://0.0.0.0:26657"

func getRandomAttributes(G *bpgroup.BpGroup, n int) []*Curve.BIG {
	attrs := make([]*Curve.BIG, n)
	for i := 0; i < n; i++ {
		attrs[i] = Curve.Randomnum(G.Order(), G.Rng())
	}
	return attrs
}

// nolint: gosec, lll, errcheck
func main() {
	cfgFile := flag.String("f", "config.toml", "Path to the client config file.")
	flag.Parse()

	syscall.Umask(0077)

	haltCh := make(chan os.Signal)
	signal.Notify(haltCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			<-haltCh
			fmt.Println("Received SIGTERM...")
			os.Exit(0)
		}
	}()

	cfg, err := config.LoadFile(*cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file '%v': %v\n", *cfgFile, err)
		os.Exit(-1)
	}

	// Start up the coconut client.
	cc, err := cclient.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to spawn client instance: %v\n", err)
		os.Exit(-1)
	}

	wholeSystem(cc)
}

func wholeSystem(cc *cclient.Client) {
	log, err := logger.New("", "DEBUG", false)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a logger: %v", err))
	}

	tmclient, err := tmclient.New(tendermintABCIAddresses, log)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a tmclient: %v", err))
	}

	// create new account
	acc := account.NewAccount()
	newAccReq, err := transaction.CreateNewAccountRequest(acc, []byte("foo"))
	if err != nil {
		panic(err)
	}

	debugAcc := &account.Account{}
	debugAcc.FromJSONFile("../tendermint/debugAccount.json")

	// transfer some funds to the new account
	transferReq, err := transaction.CreateNewTransferRequest(*debugAcc, acc.PublicKey, 20)
	if err != nil {
		panic(err)
	}

	params, _ := coconut.Setup(5)
	G := params.G
	privM := getRandomAttributes(G, 2) // sequence and the key
	token := token.New(privM[0], privM[1], int32(1))
	_ = token

	res, err := tmclient.Broadcast(newAccReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created new account. Code: %v, additional data: %v\n", code.ToString(res.DeliverTx.Code), string(res.DeliverTx.Data))
	// add some funds
	res, err = tmclient.Broadcast(transferReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transferred funds from debug to new account. Code: %v, additional data: %v\n", code.ToString(res.DeliverTx.Code), string(res.DeliverTx.Data))

	b, err := utils.GenerateRandomBytes(10)
	if err != nil {
		panic(err)
	}
	tmclient.SendAsync(append([]byte{transaction.TxAdvanceBlock, 0x01}, b...))
	tmclient.SendAsync(append([]byte{transaction.TxAdvanceBlock, 0x02}, b...))
	tmclient.SendAsync(append([]byte{transaction.TxAdvanceBlock, 0x03}, b...))
	tmclient.SendAsync(append([]byte{transaction.TxAdvanceBlock, 0x04}, b...))
	tmclient.SendAsync(append([]byte{transaction.TxAdvanceBlock, 0x05}, b...))

	fmt.Printf("Send some dummy transactions to advance block")

	cred, err := cc.GetCredential(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transfered %v to the holding account\n", token.Value())
	fmt.Printf("Obtained Credential: %v %v\n", cred.Sig1().ToString(), cred.Sig2().ToString())

	addr, err := base64.StdEncoding.DecodeString(providerAcc)
	if err != nil {
		panic(err)
	}
	// spend credential:
	didSucceed, err := cc.SpendCredential(token, cred, providerAddress, addr, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Was credential spent: ", didSucceed)
}

func basicIA(cc *cclient.Client) {
	useGRPC := false

	params, _ := coconut.Setup(5)
	G := params.G
	pubM := getRandomAttributes(G, 3)
	privM := getRandomAttributes(G, 2)
	// all possible interactions with the IAs/SPs

	if useGRPC {
		sigGrpc, _ := cc.SignAttributesGrpc(pubM)
		sigBlindGrpc, _ := cc.BlindSignAttributesGrpc(pubM, privM)
		vkGrpc, _ := cc.GetAggregateVerificationKeyGrpc()

		isValidGrpc, _ := cc.SendCredentialsForVerificationGrpc(pubM, sigGrpc, providerAddressGrpc)
		isValidBlind1Grpc, _ := cc.SendCredentialsForBlindVerificationGrpc(pubM, privM, sigBlindGrpc, providerAddressGrpc, nil)
		isValidBlind2Grpc, _ := cc.SendCredentialsForBlindVerificationGrpc(pubM, privM, sigBlindGrpc, providerAddressGrpc, vkGrpc)
		isValidBlind3Grpc, _ := cc.SendCredentialsForVerificationGrpc(append(privM, pubM...), sigBlindGrpc, providerAddressGrpc)

		fmt.Println("Is validGrpc:", isValidGrpc)
		fmt.Println("Is valid localGrpc:", coconut.Verify(params, vkGrpc, pubM, sigGrpc))

		fmt.Println("Is validBlind1Grpc:", isValidBlind1Grpc)
		fmt.Println("Is validBlind2Grpc:", isValidBlind2Grpc)
		fmt.Println("Is validBlind3Grpc:", isValidBlind3Grpc)
	} else {
		sig, _ := cc.SignAttributes(pubM)
		sigBlind, _ := cc.BlindSignAttributes(pubM, privM)
		vk, _ := cc.GetAggregateVerificationKey()

		isValid, _ := cc.SendCredentialsForVerification(pubM, sig, providerAddress)
		isValidBlind1, _ := cc.SendCredentialsForBlindVerification(pubM, privM, sigBlind, providerAddress, nil)
		isValidBlind2, _ := cc.SendCredentialsForBlindVerification(pubM, privM, sigBlind, providerAddress, vk)
		isValidBlind3, _ := cc.SendCredentialsForVerification(append(privM, pubM...), sigBlind, providerAddress)

		fmt.Println("Is valid", isValid)
		fmt.Println("Is valid local:", coconut.Verify(params, vk, pubM, sig))

		fmt.Println("Is validBlind1:", isValidBlind1)
		fmt.Println("Is validBlind2:", isValidBlind2)
		fmt.Println("Is validBlind3:", isValidBlind3)
	}

}
