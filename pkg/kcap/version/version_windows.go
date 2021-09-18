/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

const (
	// KevtSecV1 is the v1 of the kernel event section
	KevtSecV1 Version = iota + 1
)

const (
	// ProcessSecV1 is the v1 of the process section
	ProcessSecV1 Version = iota + 1
)

const (
	// HandleSecV1 is the v1 of the handle section
	HandleSecV1 Version = iota + 1
)

const (
	// PESecV1 is the v1 of the PE section
	PESecV1 Version = iota + 1
)
