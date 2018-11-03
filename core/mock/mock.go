package coremock

import (
	"context"

	commands "github.com/ipfs/go-ipfs/commands"
	core "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo"

	config "gx/ipfs/QmQjXkAchh62ST5rkmWAFHmcnkyLiZgmDgTFmNt8MYiKKw/go-ipfs-config"
	libp2p "gx/ipfs/QmRKGtF5iRsA1kNgagiKZGo2VeJxizF2tXPD3sysGX2vcc/go-libp2p"
	mocknet "gx/ipfs/QmRKGtF5iRsA1kNgagiKZGo2VeJxizF2tXPD3sysGX2vcc/go-libp2p/p2p/net/mock"
	host "gx/ipfs/QmYWxr8cuXMeisJ69SK5HCg3S7FVuARCi8JZWxDyb9Aatw/go-libp2p-host"
	pstore "gx/ipfs/QmYrNDh2z1sFB7gxYU6kX2f9N8LcWJkAzAKwZqUzvf6d1P/go-libp2p-peerstore"
	datastore "gx/ipfs/QmaRb5yNXKonhbkpNxNawoydk4N6es6b4fPj19sjEKsh5D/go-datastore"
	syncds "gx/ipfs/QmaRb5yNXKonhbkpNxNawoydk4N6es6b4fPj19sjEKsh5D/go-datastore/sync"
	testutil "gx/ipfs/Qmcxbk5F1rJ9KeJaKRoZNG7qrexftEM7Rz9y15hi5FHAuk/go-testutil"
	peer "gx/ipfs/Qme2jJHLnAfULnkrSxmqAn7pGTp3HWVmFYLnuLtxFvH9nA/go-libp2p-peer"
)

// NewMockNode constructs an IpfsNode for use in tests.
func NewMockNode() (*core.IpfsNode, error) {
	ctx := context.Background()

	// effectively offline, only peer in its network
	return core.NewNode(ctx, &core.BuildCfg{
		Online: true,
		Host:   MockHostOption(mocknet.New(ctx)),
	})
}

func MockHostOption(mn mocknet.Mocknet) core.HostOption {
	return func(ctx context.Context, id peer.ID, ps pstore.Peerstore, _ ...libp2p.Option) (host.Host, error) {
		return mn.AddPeerWithPeerstore(id, ps)
	}
}

func MockCmdsCtx() (commands.Context, error) {
	// Generate Identity
	ident, err := testutil.RandIdentity()
	if err != nil {
		return commands.Context{}, err
	}
	p := ident.ID()

	conf := config.Config{
		Identity: config.Identity{
			PeerID: p.String(),
		},
	}

	r := &repo.Mock{
		D: syncds.MutexWrap(datastore.NewMapDatastore()),
		C: conf,
	}

	node, err := core.NewNode(context.Background(), &core.BuildCfg{
		Repo: r,
	})
	if err != nil {
		return commands.Context{}, err
	}

	return commands.Context{
		Online:     true,
		ConfigRoot: "/tmp/.mockipfsconfig",
		LoadConfig: func(path string) (*config.Config, error) {
			return &conf, nil
		},
		ConstructNode: func() (*core.IpfsNode, error) {
			return node, nil
		},
	}, nil
}
