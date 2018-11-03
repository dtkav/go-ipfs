package commands

import (
	"testing"

	"github.com/ipfs/go-ipfs/namesys"

	ipns "gx/ipfs/QmVecLVqBPwV4LR4dyBXE9DeBWg1fsdTzBVFVtQnFAmGWA/go-ipns"
	tu "gx/ipfs/Qmcxbk5F1rJ9KeJaKRoZNG7qrexftEM7Rz9y15hi5FHAuk/go-testutil"
)

func TestKeyTranslation(t *testing.T) {
	pid := tu.RandPeerIDFatal(t)
	pkname := namesys.PkKeyForID(pid)
	ipnsname := ipns.RecordKey(pid)

	pkk, err := escapeDhtKey("/pk/" + pid.Pretty())
	if err != nil {
		t.Fatal(err)
	}

	ipnsk, err := escapeDhtKey("/ipns/" + pid.Pretty())
	if err != nil {
		t.Fatal(err)
	}

	if pkk != pkname {
		t.Fatal("keys didnt match!")
	}

	if ipnsk != ipnsname {
		t.Fatal("keys didnt match!")
	}
}
