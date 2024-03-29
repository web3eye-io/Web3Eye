package orbit

import (
	"context"
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

	"github.com/web3eye-io/Web3Eye/dealer/pkg/orbit/backup"
	"github.com/web3eye-io/Web3Eye/dealer/pkg/orbit/filestate"
	"github.com/web3eye-io/Web3Eye/dealer/pkg/orbit/snapshot"
)

type _orbit struct {
	ipfsRepo    repo.Repo
	ipfsNode    *ipfscore.IpfsNode
	api         coreiface.CoreAPI
	db          orbitiface.OrbitDB
	kvSnapshot  *snapshot.SnapshotKV
	kvFileState *filestate.FileStateKV
	kvBackup    *backup.BackupKV
}

var _odb = &_orbit{}

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

	logger, _ := zap.NewProduction()
	_odb.db, err = orbitdb.NewOrbitDB(ctx, _odb.api, &orbitdb.NewOrbitDBOptions{
		Directory: &cfg.OrbitRepo,
		Logger:    logger,
	})
	if err != nil {
		return err
	}

	_odb.kvSnapshot, err = snapshot.NewSnapshotKV(ctx, _odb.db)
	if err != nil {
		return err
	}

	_odb.kvFileState, err = filestate.NewFileStateKV(ctx, _odb.db)
	if err != nil {
		return err
	}

	_odb.kvBackup, err = backup.NewBackupKV(ctx, _odb.db)
	if err != nil {
		return err
	}

	return nil
}

func Snapshot() *snapshot.SnapshotKV {
	return _odb.kvSnapshot
}

func FileState() *filestate.FileStateKV {
	return _odb.kvFileState
}

func Backup() *backup.BackupKV {
	return _odb.kvBackup
}

func Finalize() {
	if _odb.kvBackup != nil {
		_odb.kvBackup.Close()
	}
	if _odb.kvFileState != nil {
		_odb.kvFileState.Close()
	}
	if _odb.kvSnapshot != nil {
		_odb.kvSnapshot.Close()
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
