//nolint:nolintlint,dupl
package token

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/google/uuid"
)

const (
	MaxPutTaskNumOnce = 100
	ReportInterval    = 100
	RedisLockTimeout  = time.Second * 10
)

type PulsarProducer struct {
	client   pulsar.Client
	producer pulsar.Producer
}

var pulsarProducer *PulsarProducer

func getPulsar() (*PulsarProducer, error) {
	if pulsarProducer != nil {
		return pulsarProducer, nil
	}
	var err error
	pulsarCli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}

	producer, err := pulsarCli.CreateProducer(pulsar.ProducerOptions{
		Topic: config.GetConfig().Pulsar.TopicTransformImage,
	})
	if err != nil {
		pulsarCli.Close()
		return nil, err
	}

	pulsarProducer = &PulsarProducer{
		client:   pulsarCli,
		producer: producer,
	}
	return pulsarProducer, nil
}

// if the VectorState is waiting,then will auto to transform imageUrl
func (s *Server) CreateToken(ctx context.Context, in *npool.CreateTokenRequest) (*npool.CreateTokenResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateToken",
			"In", in,
		)
		return &npool.CreateTokenResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	inInfo := in.GetInfo()
	entID := uuid.New().String()
	inInfo.EntID = &entID

	h, err := handler.NewHandler(ctx,
		handler.WithEntID(inInfo.EntID, true),
		handler.WithChainType(inInfo.ChainType, true),
		handler.WithChainID(inInfo.ChainID, true),
		handler.WithContract(inInfo.Contract, true),
		handler.WithTokenType(inInfo.TokenType, true),
		handler.WithTokenID(inInfo.TokenID, true),
		handler.WithOwner(inInfo.Owner, false),
		handler.WithURI(inInfo.URI, false),
		handler.WithURIState(inInfo.URIState, false),
		handler.WithURIType(inInfo.URIType, false),
		handler.WithImageURL(inInfo.ImageURL, false),
		handler.WithVideoURL(inInfo.VideoURL, false),
		handler.WithName(inInfo.Name, false),
		handler.WithDescription(inInfo.Description, false),
		handler.WithVectorState(inInfo.VectorState, false),
		handler.WithVectorID(inInfo.VectorID, false),
		handler.WithIPFSImageURL(inInfo.IPFSImageURL, false),
		handler.WithImageSnapshotID(inInfo.ImageSnapshotID, false),
		handler.WithRemark(inInfo.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateToken", "error", err)
		return &npool.CreateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.CreateToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateToken", "error", err)
		return &npool.CreateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = TransformImage(ctx, inInfo)
	if err != nil {
		logger.Sugar().Errorw("CreateToken", "action", "publish imageurl to pulsar", "error", err)
	}

	return &npool.CreateTokenResponse{
		Info: info,
	}, nil
}

// if the VectorState is waiting,then will auto to transform imageUrl
func (s *Server) UpsertToken(ctx context.Context, in *npool.UpsertTokenRequest) (*npool.UpsertTokenResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpsertToken",
			"In", in,
		)
		return &npool.UpsertTokenResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithContract(in.Info.Contract, true),
		handler.WithTokenType(in.Info.TokenType, true),
		handler.WithTokenID(in.Info.TokenID, true),
		handler.WithOwner(in.Info.Owner, false),
		handler.WithURI(in.Info.URI, false),
		handler.WithURIState(in.Info.URIState, false),
		handler.WithURIType(in.Info.URIType, false),
		handler.WithImageURL(in.Info.ImageURL, false),
		handler.WithVideoURL(in.Info.VideoURL, false),
		handler.WithName(in.Info.Name, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithVectorState(in.Info.VectorState, false),
		handler.WithVectorID(in.Info.VectorID, false),
		handler.WithIPFSImageURL(in.Info.IPFSImageURL, false),
		handler.WithImageSnapshotID(in.Info.ImageSnapshotID, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertToken", "error", err)
		return &npool.UpsertTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpsertToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertToken", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	in.Info.EntID = &info.EntID
	err = TransformImage(ctx, in.Info)
	if err != nil {
		logger.Sugar().Errorw("UpsertToken", "action", "publish imageurl to pulsar", "error", err)
	}

	return &npool.UpsertTokenResponse{
		Info: info,
	}, err
}

// if the VectorState is waiting,then will auto to transform imageUrl
func (s *Server) CreateTokens(ctx context.Context, in *npool.CreateTokensRequest) (*npool.CreateTokensResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateTokens", "error", "Infos is empty")
		return &npool.CreateTokensResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}
	inInfos := in.GetInfos()

	for i := 0; i < len(inInfos); i++ {
		entID := uuid.New().String()
		inInfos[i].EntID = &entID
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(inInfos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateTokens", "error", err)
		return &npool.CreateTokensResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.CreateTokens(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateTokens", "error", err)
		return &npool.CreateTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	for i := 0; i < len(inInfos); i++ {
		err := TransformImage(ctx, inInfos[i])
		if err != nil {
			logger.Sugar().Errorw("CreateToken", "action", "publish imageurl to pulsar", "error", err)
		}
	}

	return &npool.CreateTokensResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateToken(ctx context.Context, in *npool.UpdateTokenRequest) (*npool.UpdateTokenResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateToken",
			"In", in,
		)
		return &npool.UpdateTokenResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithContract(in.Info.Contract, false),
		handler.WithTokenType(in.Info.TokenType, false),
		handler.WithTokenID(in.Info.TokenID, false),
		handler.WithOwner(in.Info.Owner, false),
		handler.WithURI(in.Info.URI, false),
		handler.WithURIState(in.Info.URIState, false),
		handler.WithURIType(in.Info.URIType, false),
		handler.WithImageURL(in.Info.ImageURL, false),
		handler.WithVideoURL(in.Info.VideoURL, false),
		handler.WithName(in.Info.Name, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithVectorState(in.Info.VectorState, false),
		handler.WithVectorID(in.Info.VectorID, false),
		handler.WithIPFSImageURL(in.Info.IPFSImageURL, false),
		handler.WithImageSnapshotID(in.Info.ImageSnapshotID, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateBlock", "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.UpdateToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateToken", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTokenResponse{
		Info: info,
	}, nil
}

// if the VectorState is waiting,then will auto to transform imageUrl
func TransformImage(ctx context.Context, inInfo *npool.TokenReq) error {
	if inInfo.EntID == nil {
		return fmt.Errorf("not set entID")
	}

	if inInfo.VectorState.String() != npool.ConvertState_Waiting.String() {
		return nil
	}

	inInfo.VectorState = npool.ConvertState_Failed.Enum()
	if inInfo.ImageURL == nil {
		return nil
	}

	pProducer, err := getPulsar()
	if err != nil {
		return err
	}

	_, err = pProducer.producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte(*inInfo.ImageURL),
		Key:     *inInfo.EntID,
	})

	if err != nil {
		return err
	}
	inInfo.VectorState = npool.ConvertState_Processing.Enum()
	return nil
}

func (s *Server) UpdateImageVector(ctx context.Context, in *npool.UpdateImageVectorRequest) (*npool.UpdateImageVectorResponse, error) {
	var err error

	vID := int64(0)
	vState := npool.ConvertState_Failed
	remark := in.GetRemark()
	h, err := handler.NewHandler(
		ctx,
		handler.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "EntID", in.EntID, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := h.GetToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "EntID", in.EntID, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if info == nil {
		return nil, nil
	}

	if len(in.Vector) > 0 {
		milvusmgr := milvusdb.NewNFTConllectionMGR()

		if info.VectorID > 0 {
			err := milvusmgr.Delete(ctx, []int64{info.VectorID})
			if err != nil {
				remark = fmt.Sprintf("%v,%v", remark, err)
			}
		}

		ids, err := milvusmgr.Create(ctx, [][milvusdb.VectorDim]float32{imageconvert.ToArrayVector(in.Vector)})
		if err == nil {
			vState = npool.ConvertState_Success
			vID = ids[0]
		} else {
			remark = fmt.Sprintf("%v,%v", remark, err)
		}
	}

	h, err = handler.NewHandler(
		ctx,
		handler.WithID(&info.ID, true),
		handler.WithVectorID(&vID, true),
		handler.WithVectorState(&vState, true),
		handler.WithRemark(&remark, true),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "EntID", in.EntID, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err = h.UpdateToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "EntID", in.EntID, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateImageVectorResponse{
		Info: info,
	}, nil
}

func (s *Server) GetToken(ctx context.Context, in *npool.GetTokenRequest) (*npool.GetTokenResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetToken", "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetToken", "ID", in.GetID(), "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokenResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTokenOnly(ctx context.Context, in *npool.GetTokenOnlyRequest) (*npool.GetTokenOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetTokenOnly", "error", err)
		return &npool.GetTokenOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetTokens(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTokenOnly", "error", err)
		return &npool.GetTokenOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		return &npool.GetTokenOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetTokenOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetTokens(ctx context.Context, in *npool.GetTokensRequest) (*npool.GetTokensResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetTokens", "error", err)
		return &npool.GetTokensResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetTokens(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTokens", "error", err)
		return &npool.GetTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokensResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistToken(ctx context.Context, in *npool.ExistTokenRequest) (*npool.ExistTokenResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := h.ExistToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistTokenConds(ctx context.Context, in *npool.ExistTokenCondsRequest) (*npool.ExistTokenCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistTokenConds", "error", err)
		return &npool.ExistTokenCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistTokenConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistTokenConds", "error", err)
		return &npool.ExistTokenCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteToken(ctx context.Context, in *npool.DeleteTokenRequest) (*npool.DeleteTokenResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := h.DeleteToken(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "ID", in.GetID(), "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTokenResponse{
		Info: info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) TriggerTokenTransform(ctx context.Context, conds *npool.Conds) error {
	// lock
	lockKey := "TriggerTokenTransform_Lock"
	lockID, err := ctredis.TryLock(lockKey, RedisLockTimeout)
	if err != nil {
		logger.Sugar().Warn("TriggerTokenTransform", "warning", err)
		return nil
	}

	defer func() {
		err := ctredis.Unlock(lockKey, lockID)
		if err != nil {
			logger.Sugar().Warn("TriggerTokenTransform", "warning", err)
		}
	}()

	h, err := handler.NewHandler(ctx,
		handler.WithConds(conds),
		handler.WithOffset(0),
		handler.WithLimit(MaxPutTaskNumOnce),
	)

	if err != nil {
		logger.Sugar().Errorw("TriggerTokenTransform", "error", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := h.GetTokens(ctx)
	if err != nil {
		logger.Sugar().Errorw("TriggerTokenTransform", "error", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	pProducer, err := getPulsar()
	if err != nil {
		logger.Sugar().Errorw("TriggerTokenTransform", "error", err)
		return status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("put the task of transforming token image to vector,put %v tokens", total)
	for _, info := range infos {
		_, err = pProducer.producer.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte(info.ImageURL),
			Key:     info.EntID,
		})
		if err != nil {
			logger.Sugar().Errorw("TriggerTokenTransform", "msg", "faild to put task to pulsar", "error", err)
			return err
		}
	}
	return nil
}
