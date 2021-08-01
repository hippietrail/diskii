// Copyright © 2016 Zellyn Hunter <zellyn@gmail.com>

package integer

import (
	"testing"

	"github.com/zellyn/diskii/basic"
)

// helloBinary is a simple basic program used for testing. Listing
// below.
var helloBinary = []byte{
	0x16, 0x0a, 0x00, 0x5d, 0xd4, 0xc8, 0xc9, 0xd3, 0xa0, 0xc9, 0xd3, 0xa0, 0xc1, 0xa0, 0xc3, 0xcf,
	0xcd, 0xcd, 0xc5, 0xce, 0xd4, 0x01, 0x17, 0x14, 0x00, 0x5d, 0xa0, 0xd4, 0xc8, 0xc9, 0xd3, 0xa0,
	0xc9, 0xd3, 0xa0, 0xc1, 0xa0, 0xc3, 0xcf, 0xcd, 0xcd, 0xc5, 0xce, 0xd4, 0x01, 0x29, 0x1e, 0x00,
	0x5d, 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7, 0xb8, 0xb9, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5,
	0xc6, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5,
	0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0x01, 0x11, 0x28, 0x00, 0xd8, 0xb0, 0x71, 0xb1, 0x64, 0x00, 0x03,
	0xd9, 0xb0, 0x71, 0xb2, 0xc8, 0x00, 0x01, 0x11, 0x32, 0x00, 0xd8, 0xb1, 0x71, 0x2f, 0x3f, 0xb2,
	0xc8, 0x00, 0x72, 0x12, 0xb1, 0x96, 0x00, 0x01, 0x0c, 0x3c, 0x00, 0x61, 0x28, 0xc8, 0xc5, 0xcc,
	0xcc, 0xcf, 0x29, 0x01, 0x0e, 0x46, 0x00, 0x61, 0x28, 0xc7, 0xcf, 0xcf, 0xc4, 0xc2, 0xd9, 0xc5,
	0x29, 0x01, 0x0f, 0x50, 0x00, 0x61, 0x28, 0x84, 0xc3, 0xc1, 0xd4, 0xc1, 0xcc, 0xcf, 0xc7, 0x29,
	0x01, 0x05, 0x5a, 0x00, 0x51, 0x01,
}

// helloListing is the text version of the basic program above. Note
// that there are trailing newlines on lines 20 and 60.
var helloListing = `   10 REM THIS IS A COMMENT
   20 REM  THIS IS A COMMENT
   30 REM 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ
   40 X0=100:Y0=200
   50 X1= RND (200)+150
   60 PRINT "HELLO"
   70 PRINT "GOODBYE"
   80 PRINT "«ctrl-D»CATALOG"
   90 END `

// TestParse tests the full parsing and output of a basic program from
// bytes.
func TestParse(t *testing.T) {
	t.Skip("ignoring for now")
	listing, err := Decode(helloBinary)
	if err != nil {
		t.Fatal(err)
	}
	text := basic.ChevronControlCodes(listing.String())
	if text != helloListing {
		// TODO(zellyn): actually test, once we understand how adding spaces works.
		t.Fatalf("Wrong listing; want:\n%s\ngot:\n%s", helloListing, text)
	}
}
