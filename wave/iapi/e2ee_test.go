package iapi

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDirectE2EE(t *testing.T) {
	_, werr := NewParsedEntitySecrets(context.Background(), &PNewEntity{})
	require.NoError(t, werr)
	dst, werr := NewParsedEntitySecrets(context.Background(), &PNewEntity{})
	require.NoError(t, werr)

	ctx := context.Background()
	msg := make([]byte, 512)
	rand.Read(msg)

	r, err := EncryptMessage(ctx, &PEncryptMessage{
		Subject: dst.Entity,
		Content: msg,
	})
	require.NoError(t, err)

	r2, err := DecryptMessage(ctx, &PDecryptMessage{
		Decryptor:  dst.EntitySecrets,
		Ciphertext: r.Ciphertext,
	})
	require.NoError(t, err)
	require.Equal(t, msg, r2.Content)
	//kpdc := NewKeyPoolDecryptionContext()
	//kpdc.AddEntity(source.EntitySecrets.Entity)

}

func TestOAQUEE2EE(t *testing.T) {
	source, werr := NewParsedEntitySecrets(context.Background(), &PNewEntity{})
	require.NoError(t, werr)
	dst, werr := NewParsedEntitySecrets(context.Background(), &PNewEntity{})
	require.NoError(t, werr)

	ctx := context.Background()
	msg := make([]byte, 512)
	rand.Read(msg)
	r, err := EncryptMessage(ctx, &PEncryptMessage{
		Namespace:         source.Entity,
		NamespaceLocation: NewLocationSchemeInstanceURL("test", 1),
		Content:           msg,
		Resource:          "foo",
		ValidAfter:        Time(time.Now()),
		ValidBefore:       Time(time.Now().Add(30 * 24 * time.Hour)),
	})
	require.NoError(t, err)

	kpdc := NewKeyPoolDecryptionContext()
	kpdc.AddEntity(source.EntitySecrets.Entity)
	kpdc.AddEntitySecret(source.EntitySecrets, true)
	kpdc.AddDomainVisibilityID([]byte(source.Entity.Keccak256HI().MultihashString()))

	r2, err := DecryptMessage(ctx, &PDecryptMessage{
		Decryptor:  dst.EntitySecrets,
		Ciphertext: r.Ciphertext,
		Dctx:       kpdc,
	})
	require.NoError(t, err)
	require.Equal(t, msg, r2.Content)

}
