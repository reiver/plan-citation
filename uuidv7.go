package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

// generateUUIDv7 generates a UUIDv7 using the given time for the embedded timestamp.
//
// UUIDv7 layout (RFC 9562):
//
//	0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                         unix_ts_ms (48 bits)                  |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|          unix_ts_ms           | ver(0111) |    rand_a (12)    |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|var(10)|              rand_b (62 bits)                          |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                          rand_b                               |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
func generateUUIDv7(t time.Time) (string, error) {
	var buf [16]byte

	_, err := rand.Read(buf[:])
	if nil != err {
		return "", fmt.Errorf("could not read random bytes: %w", err)
	}

	ms := uint64(t.UnixMilli())

	// bytes 0-5: 48-bit unix timestamp in milliseconds (big-endian)
	buf[0] = byte(ms >> 40)
	buf[1] = byte(ms >> 32)
	buf[2] = byte(ms >> 24)
	buf[3] = byte(ms >> 16)
	buf[4] = byte(ms >> 8)
	buf[5] = byte(ms)

	// byte 6: version 7 (0111) in high nibble, keep low nibble random
	buf[6] = (buf[6] & 0x0F) | 0x70

	// byte 8: variant 10 in high 2 bits, keep low 6 bits random
	buf[8] = (buf[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16],
	), nil
}
