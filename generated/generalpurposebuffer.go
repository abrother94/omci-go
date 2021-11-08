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

// GeneralPurposeBufferClassID is the 16-bit ID for the OMCI
// Managed entity General purpose buffer
const GeneralPurposeBufferClassID ClassID = ClassID(308)

var generalpurposebufferBME *ManagedEntityDefinition

// GeneralPurposeBuffer (class ID #308)
//	This ME is created by the OLT when needed to store the results of an operation, such as a test
//	command, that needs to return a block of data of indeterminate size. The buffer is retrieved
//	with get next operations, since its size is not known a priori. An instance of this ME is
//	created and deleted by the OLT, and typically made known to an ONU ME or to an action through a
//	pointer.
//
//	The ME is defined as generically as possible, such that it can be used for other applications
//	that may not initially be apparent, such as logging. The format of its content is specific to
//	each application, and is documented there.
//
//	The general purpose buffer is neither captured in an MIB upload, nor retained in a non-volatile
//	ONU memory.
//
//	Relationships
//		Through a pointer, the OLT may associate a general purpose buffer with an ME or an operation
//		that has a need to create large or indeterminate blocks of data for subsequent upload.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Maximum Size
//			Maximum size: The ONU determines the actual size of the buffer table in the process of capturing
//			the data directed to it. The maximum size attribute permits the OLT to restrict the maximum size
//			of the buffer table. The value 0 indicates that the OLT imposes no limit on the size; it is
//			recognized that ONU implementations will impose their own limits. The ONU will not create a
//			buffer table larger than the value of this attribute. If the ONU cannot allocate enough memory
//			to accommodate this size, it should deny the ME create action or a write operation that attempts
//			to expand an existing ME. (R,-W, setbycreate) (optional) (4-bytes)
//
//		Buffer Table
//			Buffer table:	This attribute is an octet string that contains the result of some operation
//			performed on the ONU. The exact content depends on the operation, and is documented with the
//			definition of each operation. (R) (mandatory) (N bytes)
//
type GeneralPurposeBuffer struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	generalpurposebufferBME = &ManagedEntityDefinition{
		Name:    "GeneralPurposeBuffer",
		ClassID: 308,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: Uint32Field("MaximumSize", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 1),
			2: TableField("BufferTable", TableAttributeType, 0x4000, TableInfo{nil, -1}, mapset.NewSetWith(Read), true, false, false, 2),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewGeneralPurposeBuffer (class ID 308) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewGeneralPurposeBuffer(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*generalpurposebufferBME, params...)
}
