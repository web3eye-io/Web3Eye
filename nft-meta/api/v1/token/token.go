//nolint:nolintlint,dupl
package token

import (
	"context"
	"fmt"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/token"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/google/uuid"
)

func (s *Server) CreateToken(ctx context.Context, in *npool.CreateTokenRequest) (*npool.CreateTokenResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateToken", "error", err)
		return &npool.CreateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTokens(ctx context.Context, in *npool.CreateTokensRequest) (*npool.CreateTokensResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateTokens", "error", "Infos is empty")
		return &npool.CreateTokensResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTokens", "error", err)
		return &npool.CreateTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTokensResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateToken(ctx context.Context, in *npool.UpdateTokenRequest) (*npool.UpdateTokenResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateToken", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateToken", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateImageVector(ctx context.Context, in *npool.UpdateImageVectorRequest) (*npool.UpdateImageVectorResponse, error) {
	var err error

	id := in.GetID()
	vID := int64(0)
	vState := npool.ConvertState_Failed
	remark := in.GetRemark()
	if _, err := uuid.Parse(id); err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "ID", id, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	row, err := crud.Row(ctx, uuid.MustParse(id))
	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "ID", id, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if len(in.Vector) > 0 {
		milvusmgr := milvusdb.NewNFTConllectionMGR()

		if row.VectorID > 0 {
			err := milvusmgr.Delete(ctx, []int64{row.VectorID})
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

	info, err := crud.Update(ctx, &npool.TokenReq{
		ID:          &id,
		VectorID:    &vID,
		VectorState: &vState,
		Remark:      &remark,
	})

	if err != nil {
		logger.Sugar().Errorw("UpdateImageVector", "ID", id, "error", err)
		return &npool.UpdateImageVectorResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateImageVectorResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetToken(ctx context.Context, in *npool.GetTokenRequest) (*npool.GetTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetToken", "ID", in.GetID(), "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetToken", "ID", in.GetID(), "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTokenOnly(ctx context.Context, in *npool.GetTokenOnlyRequest) (*npool.GetTokenOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTokenOnly", "error", err)
		return &npool.GetTokenOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokenOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTokens(ctx context.Context, in *npool.GetTokensRequest) (*npool.GetTokensResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTokens", "error", err)
		return &npool.GetTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokensResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistToken(ctx context.Context, in *npool.ExistTokenRequest) (*npool.ExistTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistTokenConds(ctx context.Context, in *npool.ExistTokenCondsRequest) (*npool.ExistTokenCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTokenConds", "error", err)
		return &npool.ExistTokenCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountTokens(ctx context.Context, in *npool.CountTokensRequest) (*npool.CountTokensResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTokensResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteToken(ctx context.Context, in *npool.DeleteTokenRequest) (*npool.DeleteTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "ID", in.GetID(), "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "ID", in.GetID(), "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
