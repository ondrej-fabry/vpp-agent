//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp

import (
	"fmt"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

	"github.com/ligato/vpp-agent/api/models"
	vpp_l3 "github.com/ligato/vpp-agent/api/models/vpp/l3"
)

func init() {
	r := &vpp_l3.Route{
		Type: vpp_l3.Route_INTER_VRF,
	}

	fd, md := descriptor.ForMessage(r)
	fmt.Printf("FILE: %+v\n", proto.MarshalTextString(fd))
	fmt.Printf("-----------------\nMSG: %+v\n", proto.MarshalTextString(md))

	// returns specific descriptor
	x, err := proto.GetExtension(md.Options, models.E_Spec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("models.Spec: %+v\n", x)
	// output: models.Spec: module:"vpp.l3" version:"v2" type:"route"

	// returns extensions descriptors
	ed, err := proto.ExtensionDescs(md.Options)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExtensionDescs:")
	for _, ee := range ed {
		fmt.Printf(" - %+v\n", ee)
	}
	// output: &{ExtendedType:<nil> ExtensionType:<nil> Field:50222 Name:models.spec Tag:bytes,50222,opt,name=spec Filename:models/options.proto}

	// returns all registered including ones from gogoproto.*
	e := proto.RegisteredExtensions(md.Options)
	fmt.Println("RegisteredExtensions:")
	for _, ee := range e {
		fmt.Printf(" - %+v\n", ee)
	}

	os.Exit(1)
}
