// Copyright 2023 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// SignatureType SignatureType is the type of a cryptographic signature. * ecdsa: `r (32-bytes) || s
// (32-bytes)` - `64 bytes` * ecdsa_recovery: `r (32-bytes) || s (32-bytes) || v (1-byte)` - `65
// bytes` * ed25519: `R (32-byte) || s (32-bytes)` - `64 bytes` * schnorr_1: `r (32-bytes) || s
// (32-bytes)` - `64 bytes`  (schnorr signature implemented by Zilliqa where both `r` and `s` are
// scalars encoded as `32-bytes` values, most significant byte first.) * schnorr_poseidon: `r
// (32-bytes) || s (32-bytes)` where s = Hash(1st pk || 2nd pk || r) - `64 bytes`  (schnorr
// signature w/ Poseidon hash function implemented by O(1) Labs where both `r` and `s` are scalars
// encoded as `32-bytes` values, least significant byte first.
// https://github.com/CodaProtocol/signer-reference/blob/master/schnorr.ml )
type SignatureType string

// List of SignatureType
const (
	Ecdsa           SignatureType = "ecdsa"
	EcdsaRecovery   SignatureType = "ecdsa_recovery"
	Ed25519         SignatureType = "ed25519"
	Schnorr1        SignatureType = "schnorr_1"
	SchnorrPoseidon SignatureType = "schnorr_poseidon"
)
