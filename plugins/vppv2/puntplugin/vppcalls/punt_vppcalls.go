// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"strings"

	"github.com/go-errors/errors"
	punt "github.com/ligato/vpp-agent/api/models/vpp/punt"
	api_ip "github.com/ligato/vpp-binapi/binapi/ip"
	api_punt "github.com/ligato/vpp-binapi/binapi/punt"
)

// AddPunt configures new punt entry
func (h *PuntVppHandler) AddPunt(puntCfg *punt.ToHost) error {
	return h.handlePuntToHost(puntCfg, true)
}

// DeletePunt removes punt entry
func (h *PuntVppHandler) DeletePunt(puntCfg *punt.ToHost) error {
	return h.handlePuntToHost(puntCfg, false)
}

func (h *PuntVppHandler) handlePuntToHost(toHost *punt.ToHost, isAdd bool) error {
	req := &api_punt.SetPunt{
		IsAdd: boolToUint(isAdd),
		Punt: api_punt.Punt{
			IPv:        resolveL3Proto(toHost.L3Protocol),
			L4Protocol: resolveL4Proto(toHost.L4Protocol),
			L4Port:     uint16(toHost.Port),
		},
	}
	reply := &api_punt.SetPuntReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

// RegisterPuntSocket registers new punt to socket
func (h *PuntVppHandler) RegisterPuntSocket(puntCfg *punt.ToHost) error {
	if puntCfg.L3Protocol == punt.L3Protocol_IPv4 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.registerPuntWithSocketIPv4(puntCfg); err != nil {
			return err
		}
	}
	if puntCfg.L3Protocol == punt.L3Protocol_IPv6 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.registerPuntWithSocketIPv6(puntCfg); err != nil {
			return err
		}
	}
	return nil
}

func (h *PuntVppHandler) registerPuntWithSocketIPv4(punt *punt.ToHost) error {
	return h.registerPuntWithSocket(punt, true)
}

func (h *PuntVppHandler) registerPuntWithSocketIPv6(punt *punt.ToHost) error {
	return h.registerPuntWithSocket(punt, false)
}

func (h *PuntVppHandler) registerPuntWithSocket(toHost *punt.ToHost, isIPv4 bool) error {
	//pathName := []byte(toHost.SocketPath)
	//pathByte := make([]byte, 108) // linux sun_path defined to 108 bytes as by unix(7)
	//for i, c := range pathName {
	//	pathByte[i] = c
	//}

	req := &api_punt.PuntSocketRegister{
		HeaderVersion: 1,
		Punt: api_punt.Punt{
			IPv:        resolveL3Proto(toHost.L3Protocol),
			L4Protocol: resolveL4Proto(toHost.L4Protocol),
			L4Port:     uint16(toHost.Port),
		},
		Pathname: []byte(toHost.SocketPath),
	}
	reply := &api_punt.PuntSocketRegisterReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	p := *toHost
	p.SocketPath = strings.SplitN(string(reply.Pathname), "\x00", 2)[0]
	socketPathMap[toHost.Port] = &p

	return nil
}

// DeregisterPuntSocket removes existing punt to socket sogistration
func (h *PuntVppHandler) DeregisterPuntSocket(puntCfg *punt.ToHost) error {
	if puntCfg.L3Protocol == punt.L3Protocol_IPv4 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.unregisterPuntWithSocketIPv4(puntCfg); err != nil {
			return err
		}
	}
	if puntCfg.L3Protocol == punt.L3Protocol_IPv6 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.unregisterPuntWithSocketIPv6(puntCfg); err != nil {
			return err
		}
	}
	return nil
}

func (h *PuntVppHandler) unregisterPuntWithSocketIPv4(punt *punt.ToHost) error {
	return h.unregisterPuntWithSocket(punt, true)
}

func (h *PuntVppHandler) unregisterPuntWithSocketIPv6(punt *punt.ToHost) error {
	return h.unregisterPuntWithSocket(punt, false)
}

func (h *PuntVppHandler) unregisterPuntWithSocket(toHost *punt.ToHost, isIPv4 bool) error {
	req := &api_punt.PuntSocketDeregister{
		Punt: api_punt.Punt{
			IPv:        resolveL3Proto(toHost.L3Protocol),
			L4Protocol: resolveL4Proto(toHost.L4Protocol),
			L4Port:     uint16(toHost.Port),
		},
	}
	reply := &api_punt.PuntSocketDeregisterReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	delete(socketPathMap, toHost.Port)

	return nil
}

// AddPuntRedirect adds new redirect entry
func (h *PuntVppHandler) AddPuntRedirect(puntCfg *punt.IPRedirect) error {
	if puntCfg.L3Protocol == punt.L3Protocol_IPv4 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.handlePuntRedirectIPv4(puntCfg, true); err != nil {
			return err
		}
	}
	if puntCfg.L3Protocol == punt.L3Protocol_IPv6 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.handlePuntRedirectIPv6(puntCfg, true); err != nil {
			return err
		}
	}
	return nil
}

// DeletePuntRedirect removes existing redirect entry
func (h *PuntVppHandler) DeletePuntRedirect(puntCfg *punt.IPRedirect) error {
	if puntCfg.L3Protocol == punt.L3Protocol_IPv4 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.handlePuntRedirectIPv4(puntCfg, false); err != nil {
			return err
		}
	}
	if puntCfg.L3Protocol == punt.L3Protocol_IPv6 || puntCfg.L3Protocol == punt.L3Protocol_ALL {
		if err := h.handlePuntRedirectIPv6(puntCfg, false); err != nil {
			return err
		}
	}
	return nil
}

func (h *PuntVppHandler) handlePuntRedirectIPv4(punt *punt.IPRedirect, isAdd bool) error {
	return h.handlePuntRedirect(punt, true, isAdd)
}

func (h *PuntVppHandler) handlePuntRedirectIPv6(punt *punt.IPRedirect, isAdd bool) error {
	return h.handlePuntRedirect(punt, false, isAdd)
}

func (h *PuntVppHandler) handlePuntRedirect(punt *punt.IPRedirect, isIPv4, isAdd bool) error {
	// rx interface
	var rxIfIdx uint32
	if punt.RxInterface == "" {
		rxIfIdx = ^uint32(0)
	} else {
		rxMetadata, exists := h.ifIndexes.LookupByName(punt.RxInterface)
		if !exists {
			return errors.Errorf("index not found for interface %s", punt.RxInterface)
		}
		rxIfIdx = rxMetadata.SwIfIndex
	}

	// tx interface
	txMetadata, exists := h.ifIndexes.LookupByName(punt.TxInterface)
	if !exists {
		return errors.Errorf("index not found for interface %s", punt.TxInterface)
	}

	// next hop address
	nextHop, err := ipToAddress(punt.NextHop)
	if err != nil {
		return err
	}

	req := &api_ip.IPPuntRedirect{
		IsAdd: boolToUint(isAdd),
		Punt: api_ip.PuntRedirect{
			RxSwIfIndex: rxIfIdx,
			TxSwIfIndex: txMetadata.SwIfIndex,
			Nh:          nextHop,
		},
	}
	reply := &api_ip.IPPuntRedirectReply{}
	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

func resolveL3Proto(protocol punt.L3Protocol) uint8 {
	switch protocol {
	case punt.L3Protocol_IPv4:
		return uint8(punt.L3Protocol_IPv4)
	case punt.L3Protocol_IPv6:
		return uint8(punt.L3Protocol_IPv6)
	case punt.L3Protocol_ALL:
		return ^uint8(0) // binary API representation for both protocols
	}
	return uint8(punt.L3Protocol_UNDEFINED_L3)
}

func resolveL4Proto(protocol punt.L4Protocol) uint8 {
	switch protocol {
	case punt.L4Protocol_TCP:
		return uint8(punt.L4Protocol_TCP)
	case punt.L4Protocol_UDP:
		return uint8(punt.L4Protocol_UDP)
	}
	return uint8(punt.L4Protocol_UNDEFINED_L4)
}

func boolToUint(input bool) uint8 {
	if input {
		return 1
	}
	return 0
}
