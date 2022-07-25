package utils

type Encryption struct {
	passphrase string
}

func NewEncryption(passphrase string) *Encryption {
	return &Encryption{
		passphrase: passphrase,
	}
}

func (e Encryption) Decrypt(ciphertext string) string {
	return Decrypt(e.passphrase, ciphertext)
}

func (e Encryption) Encrypt(plaintext string) string {
	return Encrypt(e.passphrase, plaintext)
}
