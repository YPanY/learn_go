package rfc4226

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"math"
)

// EncryptSHA1
/* Step 1 :
Generate an HMAC-SHA-1 value Let HS = HMAC-SHA-1(K,C)
HS is a 20-byte string
*/
func EncryptSHA1(secret string, count uint64) []byte {
	hash := hmac.New(sha1.New, []byte(secret))
	eightByteCount := make([]byte, 8)
	binary.BigEndian.PutUint64(eightByteCount, count)
	hash.Write(eightByteCount)
	return hash.Sum(nil)
}

/* Step 2: Generate a 4-byte string (Dynamic Truncation)
Let Sbits = DT(HS)   //  DT, defined below,
  returns a 31-bit string
*/

//DT
func DT(hs []byte, digit int) int {
	last4Bit := hs[len(hs)-1] & 0x0F
	truncatedHex := hs[last4Bit : last4Bit+4]
	truncated := int(binary.BigEndian.Uint32(truncatedHex) & 0x7FFFFFFF)
	return truncated % (int(math.Pow10(digit)))
}

// Step 3: Compute an HOTP value
