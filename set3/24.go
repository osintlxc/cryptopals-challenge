/*
 * Create the MT19937 stream cipher and break it
 *
 * You can create a trivial stream cipher out of any PRNG; use it to generate a
 * sequence of 8 bit outputs and call those outputs a keystream. XOR each byte
 * of plaintext with each successive byte of keystream.
 *
 * Write the function that does this for MT19937 using a 16-bit seed. Verify
 * that you can encrypt and decrypt properly. This code should look similar to
 * your CTR code.
 *
 * Use your function to encrypt a known plaintext (say, 14 consecutive 'A'
 * characters) prefixed by a random number of random characters.
 *
 * From the ciphertext, recover the "key" (the 16 bit seed).
 *
 * Use the same idea to generate a random "password reset token" using MT19937
 * seeded from the current time.
 *
 * Write a function to check if any given password token is actually the product
 * of an MT19937 PRNG seeded with the current time.
 *
 */

package set_three

func (mt *mersenneTwister) CryptBlocks(dst, src []byte) {
	if len(dst) < len(src) {
		panic("mt19937Cipher: output smaller than input")
	}

	for i, b := range src {
		// XOR plaintext with 1 byte of output from MT
		dst[i] = b ^ byte(mt.Extract()&0xFF)
	}
}

// The keyspace here is only 2^16, so we can pretty easily brute force it
// Find the random prefix length with len(encrypted) - len(plaintext) and then
// brute force those bytes in the ciphertext until the XOR matches plaintext.
