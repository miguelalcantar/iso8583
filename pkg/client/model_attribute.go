/*
 * ISO8583 API
 *
 * Package github.com/moov-io/iso8583 implements a file reader and writer written in Go decorated with a HTTP API for creating, parsing, and validating financial transaction card originated interchange messaging. User can seed iso8583's specification file as json file, iso8583 message with several formsts (mesage binary, json, xml)  | Input      | Output     |  |------------|------------|  | JSON       | JSON       |  | XML        | XML        |  | MESSAGE    | MESSAGE    |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Attribute a (key, attribute) map.
type Attribute struct {
	Default AttributeItem `json:"default,omitempty"`
}
