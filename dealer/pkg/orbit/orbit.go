package orbit

import (
	"context"
	"encoding/binary"
	"os"
	"os/exec"

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

const (
	CurrentSnapshotIndex = "current-index"
	KVStoreSnapshotIndex = "snapshot-index"
)

func Initialize(ctx context.Context) error {
	cfg := config.GetConfig().Dealer
	os.Setenv("IPFS_PATH", cfg.IpfsRepo)

	cmd := exec.Command("ipfs", "stats", "repo")
	if err := cmd.Run(); err != nil {
		cmd = exec.Command("ipfs", "init")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

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
	_odb.kvSnapshotIndex, err = _odb.db.KeyValue(ctx, KVStoreSnapshotIndex, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return err
	}

	if err := _odb.kvSnapshotIndex.Load(ctx, -1); err != nil { //nolint
		return err
	}

	b, err := _odb.kvSnapshotIndex.Get(ctx, CurrentSnapshotIndex)
	if err != nil {
		return err
	}
	_odb.snapshotIndex, _ = binary.Uvarint(b)

	if _odb.snapshotIndex == 0 {
		b := make([]byte, 8)
		_odb.snapshotIndex += 1
		binary.PutUvarint(b, _odb.snapshotIndex)
		if _, err := _odb.kvSnapshotIndex.Put(ctx, CurrentSnapshotIndex, b); err != nil {
			return err
		}
	}

	return nil
}

func Finalize() {
	if _odb.kvSnapshotIndex != nil {
		_odb.kvSnapshotIndex.Close()
	}
	if _odb.db != nil {
		_odb.db.Close()
	}
	if _odb.ipfsNode != nil {
		_odb.ipfsNode.Close()
	}
	if _odb.ipfsRepo != nil {
		_odb.ipfsRepo.Close()
	}
}
