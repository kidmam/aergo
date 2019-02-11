/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package audit

type AuditCategory int32

const (
	ShortTerm AuditCategory = iota
	LongTerm
	Permanent
)