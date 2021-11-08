/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// RtpProfileDataClassID is the 16-bit ID for the OMCI
// Managed entity RTP profile data
const RtpProfileDataClassID ClassID = ClassID(143)

var rtpprofiledataBME *ManagedEntityDefinition

// RtpProfileData (class ID #143)
//	This ME configures RTP. It is conditionally required for ONUs that offer VoIP service. If a non-
//	OMCI interface is used to manage VoIP, this ME is unnecessary.
//
//	An instance of this ME is created and deleted by the OLT. An RTP profile is needed for each
//	unique set of attributes.
//
//	Relationships
//		An instance of this ME may be associated with one or more VoIP media profile MEs.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Local Port Min
//			Local port min: This attribute defines the base UDP port that should be used by RTP for voice
//			traffic. The recommended default is 50000 (R,-W, set-by-create) (mandatory) (2-bytes)
//
//		Local Port Max
//			Local port max: This attribute defines the highest UDP port used by RTP for voice traffic. The
//			value must be greater than the local port minimum. The value 0 specifies that the local port
//			maximum be equal to the local port minimum. (R,-W, set-by-create) (optional) (2-bytes)
//
//		Dscp Mark
//			DSCP mark:	Diffserv code point to be used for outgoing RTP packets for this profile. The
//			recommended default value is expedited forwarding (EF)-= 0x2E. (R,-W, setbycreate) (mandatory)
//			(1-byte)
//
//		Piggyback Events
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Tone Events
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Dtmf Events
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Cas Events
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Ip Host Config Pointer
//			IP host config pointer: This optional pointer associates the bearer (voice) flow with an IP host
//			config data or IPv6 host config data ME. If this attribute is not present or is not populated
//			with a valid pointer value, the bearer flow uses the same IP stack that is used for signalling,
//			indicated by the TCP/UDP pointer in the associated SIP agent or MGC config data. The default
//			value is 0xFFFF, a null pointer. (R,-W) (optional) (2-bytes)
//
type RtpProfileData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	rtpprofiledataBME = &ManagedEntityDefinition{
		Name:    "RtpProfileData",
		ClassID: 143,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: Uint16Field("LocalPortMin", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2: Uint16Field("LocalPortMax", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 2),
			3: ByteField("DscpMark", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 3),
			4: ByteField("PiggybackEvents", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 4),
			5: ByteField("ToneEvents", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 5),
			6: ByteField("DtmfEvents", UnsignedIntegerAttributeType, 0x0400, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 6),
			7: ByteField("CasEvents", UnsignedIntegerAttributeType, 0x0200, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 7),
			8: Uint16Field("IpHostConfigPointer", UnsignedIntegerAttributeType, 0x0100, 0, mapset.NewSetWith(Read, Write), false, true, false, 8),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewRtpProfileData (class ID 143) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewRtpProfileData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*rtpprofiledataBME, params...)
}
