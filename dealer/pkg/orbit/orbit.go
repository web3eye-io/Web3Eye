package orbit

import (
	"context"
	"encoding/binary"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
	coreiface "github.com/ipfs/boxo/coreiface"
	ipfscore "github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo"
	"github.com/ipfs/kubo/repo/fsrepo"
	"github.com/web3eye-io/Web3Eye/config"
	"go.uber.org/zap"
)

type _orbit struct {
	ipfsRepo        repo.Repo
	ipfsNode        *ipfscore.IpfsNode
	api             coreiface.CoreAPI
	db              orbitiface.OrbitDB
	kvSnapshotIndex orbitdb.KeyValueStore
	snapshotIndex   uint64
}

var _odb = &_orbit{}

func Initialize(ctx context.Context) error {
	cfg := config.GetConfig().Dealer

	plugins, err := loader.NewPluginLoader(cfg.IpfsRepo)
	if err != nil {
		return err
	}
	if err := plugins.Initialize(); err != nil {
		return err
	}
	if err := plugins.Inject(); err != nil {
		return err
	}

	_odb.ipfsRepo, err = fsrepo.Open(cfg.IpfsRepo)
	if err != nil {
		return err
	}

	_odb.ipfsNode, err = ipfscore.NewNode(ctx, &ipfscore.BuildCfg{
		Online: true,
		Repo:   _odb.ipfsRepo,
		ExtraOpts: map[string]bool{
			"pubsub": true,
		},
	})
	if err != nil {
		return err
	}

	_odb.api, err = coreapi.NewCoreAPI(_odb.ipfsNode)
	if err != nil {
		return err
	}

	_odb.db, err = orbitdb.NewOrbitDB(ctx, _odb.api, &orbitdb.NewOrbitDBOptions{
		Directory: &cfg.OrbitRepo,
		Logger:    zap.NewExample(),
	})
	if err != nil {
		return err
	}

	replicate := true
	_odb.kvSnapshotIndex, err = _odb.db.KeyValue(ctx, "snapshot-index", &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return err
	}

	b, err := _odb.kvSnapshotIndex.Get(ctx, "current-index")
	if err != nil {
		return err
	}
	_odb.snapshotIndex = binary.LittleEndian.Uint64(b)

	return nil
}

func Finalize() {
	_odb.kvSnapshotIndex.Close()
	_odb.db.Close()
	_odb.ipfsNode.Close()
	_odb.ipfsRepo.Close()
}
