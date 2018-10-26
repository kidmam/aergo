/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2p

import (
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMustParseBytes(t *testing.T) {
	sampleUUID := uuid.Must(uuid.NewRandom())
	tests := []struct {
		name string
		in []byte
		expectErr bool
	}{
		{"TSucc", sampleUUID[:], false},
		{"TWrongSize", sampleUUID[:15], true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseBytesToMsgID(test.in)
			assert.Equal(t, test.expectErr, err != nil)

		})
	}
}

