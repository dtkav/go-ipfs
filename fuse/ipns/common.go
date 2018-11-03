package ipns

import (
	"context"

	"github.com/ipfs/go-ipfs/core"
	nsys "github.com/ipfs/go-ipfs/namesys"
	ft "gx/ipfs/QmTMiwq56r2vwz6v7Y6DNWaTLC1E1dWd5APXfB9o4snvp2/go-unixfs"
	path "gx/ipfs/Qmc3xBbgLrsqwFTn2eAdiMtRpesYVpMdhaQG75WEe6XfaT/go-path"
	ci "gx/ipfs/QmcgyFuP2SVMhgkTjFaXwnyAi9Fheg6A7SWgpYJKuajPnK/go-libp2p-crypto"
)

// InitializeKeyspace sets the ipns record for the given key to
// point to an empty directory.
func InitializeKeyspace(n *core.IpfsNode, key ci.PrivKey) error {
	ctx, cancel := context.WithCancel(n.Context())
	defer cancel()

	emptyDir := ft.EmptyDirNode()

	err := n.Pinning.Pin(ctx, emptyDir, false)
	if err != nil {
		return err
	}

	err = n.Pinning.Flush()
	if err != nil {
		return err
	}

	pub := nsys.NewIpnsPublisher(n.Routing, n.Repo.Datastore())

	return pub.Publish(ctx, key, path.FromCid(emptyDir.Cid()))
}
