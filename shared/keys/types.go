package keys

type KeyStore map[string]string

func (ks KeyStore) GetPrivateKey(chainID string) (string, bool) {
	privateKey, ok := ks[chainID]
	return privateKey, ok
}
