package token

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	entranceproto "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/token"

	"github.com/web3eye-io/Web3Eye/common/servermux"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/entrance/resource"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/token"
	"google.golang.org/grpc"

	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

// 8mb

const (
	MaxUploadFileSize = 1 << 10 << 10 << 3
	UploadFileFeild   = "UploadFile"
	LimitFeild        = "Limit"
)

type Img2VectorResp struct {
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

type SearchToken struct {
	nftmetaproto.Token
	Distance float32
}

var (
	pbJsonMarshaler jsonpb.Marshaler
)

func init() {
	pbJsonMarshaler = jsonpb.Marshaler{
		EmitDefaults: true,
	}
	mux := servermux.AppServerMux()
	mux.HandleFunc("/search/file", Search)

	pages, err := fs.Sub(resource.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(pages)))
}

// nolint
func Search(w http.ResponseWriter, r *http.Request) {
	// ret := `{"Infos":[{"ID":"7a3fff40-4f78-4de2-9f34-ca2e2b942e6e","ChainType":"Ethereum","ChainID":"1","Contract":"0x412039fF52f96D18570395C4dDcAa2cAC5707381","TokenType":"ERC721","TokenID":"1109","Owner":"","URI":"ipfs://QmQmrfkcuASkyAKPym2VmWPRjtpNSumTs8KDgLyMk6VPAH/1109","URIType":"ipfs","ImageURL":"ipfs://QmR9sexEQLMxVNzjDpYphXKmi2cACfzdCM1afXh6e6cDL4/1109.png","VideoURL":"","Description":"Cool Jpegs created by HPOP8I","Name":"Blip #1109","VectorState":"Success","VectorID":"444026885366617521","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"754e26ad-8ba6-4bef-a057-30ca2f39aea7","TokenID":"258","ImageURL":"ipfs://QmR9sexEQLMxVNzjDpYphXKmi2cACfzdCM1afXh6e6cDL4/258.png","IPFSImageURL":""},{"ID":"cda55d3c-c91e-45d9-b226-31726ea275b4","TokenID":"1829","ImageURL":"ipfs://QmR9sexEQLMxVNzjDpYphXKmi2cACfzdCM1afXh6e6cDL4/1829.png","IPFSImageURL":""},{"ID":"b8f5a5b2-7aad-4194-baf0-45db8a2ada03","TokenID":"419","ImageURL":"ipfs://QmR9sexEQLMxVNzjDpYphXKmi2cACfzdCM1afXh6e6cDL4/419.png","IPFSImageURL":""}],"SiblingsNum":4,"Distance":0},{"ID":"e3eb2a96-0009-471b-a9d1-a860df1f415e","ChainType":"Ethereum","ChainID":"1","Contract":"0xf30A9cd4Cd1Fd9EA3270afcecde5feCe34Bc4aCa","TokenType":"ERC721","TokenID":"8543","Owner":"","URI":"ipfs://bafybeid6otrd6a4ibafpteqcokvoyhuydpzzhmfqckbs3563plgfumruza/8543.json","URIType":"ipfs","ImageURL":"ipfs://bafybeicnzdrcbzvcqkp7qimmup5sizvevhxsn2a773e7xrzjzssbwqsppi/8543.png","VideoURL":"","Description":"","Name":"BoringNaka #8543","VectorState":"Success","VectorID":"444026885366615513","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.7466283},{"ID":"b6e4795d-c057-40fb-9df0-b126be40a1de","ChainType":"Ethereum","ChainID":"1","Contract":"0xD260c7aE9c8e425f02354E799cA2EB276410570A","TokenType":"ERC721","TokenID":"5313","Owner":"","URI":"https://gateway.pinata.cloud/ipfs/QmPJc5BvN2kezU4NGnxDdhPeEwRiniZT31GKdCqK4u8emt/5313","URIType":"ipfs-gateway","ImageURL":"https://gateway.pinata.cloud/ipfs/QmcQvTWn4cbUsKqxWgEiHBMyauvr6rJLPgRAhU7gYaoAn6/5313.png","VideoURL":"","Description":"Baby Doge Army is a collection of 10,000 adoptable baby doges. A unique digital art collection waiting to be rescued on the Ethereum Blockchain. Each one has been generated then hand-groomed by our team to be fit for adoption. Join us on our mission and have a good time. Having a Baby Doge grants you creative and commercial rights, as well as inclusion in the community, plus feel great knowing your NFT helped make a difference to save dogs in need.","Name":"Baby Doge 3D #5313","VectorState":"Success","VectorID":"444026885366618793","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.7569006},{"ID":"c24ebd6d-358f-48e3-a489-4a86df3f6ffe","ChainType":"Ethereum","ChainID":"1","Contract":"0x5Af0D9827E0c53E4799BB226655A1de152A425a5","TokenType":"ERC721","TokenID":"1687","Owner":"","URI":"https://www.miladymaker.net/milady/json/1687","URIType":"http","ImageURL":"https://www.miladymaker.net/milady/1687.png","VideoURL":"","Description":"Milady Maker is a collection of 10,000 generative pfpNFT's in a neochibi aesthetic inspired by Tokyo street style tribes.","Name":"Milady 1687","VectorState":"Success","VectorID":"444026885366616093","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":2,"SiblingTokens":[{"ID":"0bc5b85f-7f72-47ab-9528-cfa7cff1575a","TokenID":"5114","ImageURL":"https://miladymaker.net/milady/5114.png","IPFSImageURL":""},{"ID":"4733fcda-28aa-4723-a42f-34df73acf4ef","TokenID":"6760","ImageURL":"https://miladymaker.net/milady/6760.png","IPFSImageURL":""},{"ID":"2addb9c1-0615-4687-9db5-ff2a7ea40a5d","TokenID":"5128","ImageURL":"https://miladymaker.net/milady/5128.png","IPFSImageURL":""},{"ID":"ecb32db0-0715-47d2-bd8f-2e07c623bac6","TokenID":"2073","ImageURL":"https://www.miladymaker.net/milady/2073.png","IPFSImageURL":""}],"SiblingsNum":5,"Distance":0.757064},{"ID":"b18bf58f-1f6e-423f-ac32-66b6d8fd4f61","ChainType":"Ethereum","ChainID":"1","Contract":"0xDd63a22d90cADdDaA5ec36331F38724F415d2702","TokenType":"ERC721","TokenID":"4666","Owner":"","URI":"https://app.bueno.art/api/contract/gbxzzzxOKSDQoTiXwybdI/chain/1/metadata/4666","URIType":"http","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/4666?s=cbd53705f5b8852270a55280bb02d49c","VideoURL":"","Description":"6346 Sad Dead roam the blockchain to remind the living of one thing: Living without regret and not giving a damn!","Name":"Sad Dead Club #4666","VectorState":"Success","VectorID":"444026885366617075","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"3d915abe-7ff7-4bbf-bd1b-1cf62017e9e0","TokenID":"2751","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/2751?s=a59554e8ac80ff12e8a159f152f368d0","IPFSImageURL":""},{"ID":"72db5fa3-2e17-4ab4-a1c8-4a1e0f837c0c","TokenID":"6211","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/6211?s=ac7d1043ba4c943c46eb7724d655a921","IPFSImageURL":""},{"ID":"2927612b-9b52-435a-9234-bd92113f8d88","TokenID":"6215","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/6215?s=43311e3af7f26ed17ef674d80a206a37","IPFSImageURL":""},{"ID":"f6379794-5778-4605-a1ee-f199fb3d40a6","TokenID":"5621","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/5621?s=6853b8fc29549900a2d3c8b6adec913c","IPFSImageURL":""},{"ID":"9297eb74-cbc4-4c52-9188-02c87f71ef3a","TokenID":"6214","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/6214?s=68f456d49bcde876c58003badf4f56db","IPFSImageURL":""},{"ID":"40c4b253-d870-4002-a126-f39faffbff15","TokenID":"1972","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/1972?s=fcad23792bd5694a8709eed64fa41bfa","IPFSImageURL":""},{"ID":"39040e94-f0b1-4594-84b8-2385908fbaed","TokenID":"2146","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/2146?s=f216ef4b766c20993a30b992c1f0d731","IPFSImageURL":""},{"ID":"d79ae7e7-cf3a-4d97-9d3e-c90ba53a5e28","TokenID":"1917","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/1917?s=5b70fe4317872733215b1da43de5ff3d","IPFSImageURL":""},{"ID":"6322ed86-a375-42c8-8d8e-9b353a3f6424","TokenID":"4660","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/4660?s=d4ec78e5d5e032a96fe0b0093b5c817f","IPFSImageURL":""},{"ID":"780c70ef-fab0-469c-a017-2444b5c8ad8c","TokenID":"4633","ImageURL":"https://assets.bueno.art/images/29dfe06a-83cd-4e53-83cb-66921406a070/default/4633?s=a49abe3a0bfda353a1edb5e1ad57ec59","IPFSImageURL":""}],"SiblingsNum":31,"Distance":0.76034224},{"ID":"ef13af49-a822-488f-9134-faccf21ce56c","ChainType":"Ethereum","ChainID":"1","Contract":"0xCC845392C20a5836b8f5d2D3D88EF7b5B1820644","TokenType":"ERC721","TokenID":"3726","Owner":"","URI":"ipfs://Qmez3EWDrMT6oR6JTfRE4mskD7XFuZtRXhRshsc3tNdkwg/3726.json","URIType":"ipfs","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3726.png","VideoURL":"","Description":"Bored Punk is a unique and exciting NFT project that combines the distinctive features of Bored with the rebellious spirit of Punk","Name":"BoredPunks #3726","VectorState":"Success","VectorID":"444026885366616199","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"6ad5fb56-6299-45b2-b545-ad2844064d5a","TokenID":"3714","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3714.png","IPFSImageURL":""},{"ID":"833fcb28-a384-4c82-8d0a-fca5ac357953","TokenID":"3712","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3712.png","IPFSImageURL":""},{"ID":"10699b15-8404-4d48-aff9-0c9977e34021","TokenID":"3709","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3709.png","IPFSImageURL":""},{"ID":"ba07e380-d8e2-4401-b84d-f7eb947e9a3a","TokenID":"3710","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3710.png","IPFSImageURL":""},{"ID":"63de8c6d-ae1a-46ba-a135-4c3f55c68449","TokenID":"3707","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3707.png","IPFSImageURL":""},{"ID":"cea02cc9-f34d-4387-b5fa-eb693f762ff3","TokenID":"3723","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3723.png","IPFSImageURL":""},{"ID":"c67cd091-25dc-4422-a281-247190d4dd7e","TokenID":"3729","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3729.png","IPFSImageURL":""},{"ID":"892514e2-b394-4bc8-9e2e-28af353424ff","TokenID":"3728","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3728.png","IPFSImageURL":""},{"ID":"70020b45-d647-4ebb-8e7a-b00370ded7ee","TokenID":"3721","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3721.png","IPFSImageURL":""},{"ID":"1b261267-1f12-4a07-b853-a7d0e3d12b47","TokenID":"3722","ImageURL":"ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3722.png","IPFSImageURL":""}],"SiblingsNum":24,"Distance":0.7798524},{"ID":"446c2f62-8a37-4ced-b4b8-fc007bec3447","ChainType":"Ethereum","ChainID":"1","Contract":"0x5E4aAB148410DE1CB50cDCD5108e1260Cc36d266","TokenType":"ERC721","TokenID":"898","Owner":"","URI":"https://badbears.mypinata.cloud/ipfs/QmVLeFqrnDE8ZjZuVp4kXMDh7Uv18HxQXc98wPJV1x9DeR/898","URIType":"ipfs-gateway","ImageURL":"https://badbears.mypinata.cloud/ipfs/QmRx41npxcYrHZXAg5ozxpKcDaV9AdgJeeZbWjyCxB6eiU/898-bored-hazy-ink-bear.png","VideoURL":"","Description":"5,555 Bad Bears are taking over the Metaverse with a one-of-a-kind tokenized, ERC-721 based ecosystem built on the Ethereum blockchain. \n\n [View Bear](https://app.badbears.io/bear/898)","Name":"#898 • Bored Hazy Ink Bear","VectorState":"Success","VectorID":"444026885366616493","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.7852753},{"ID":"421bd2fc-e164-41b3-8d10-b1bb872e9c18","ChainType":"Ethereum","ChainID":"1","Contract":"0x2F073c4a897c615101Fe4dF00ea0869191c6FA8d","TokenType":"ERC721","TokenID":"156","Owner":"","URI":"https://drunkrobots.net/nft/metadata/156.json","URIType":"http","ImageURL":"https://drunkrobots.net/nft/robot/156.png","VideoURL":"","Description":"This is Karsten 🤖 Drunk Robot #156. Favorite Drink: Greyhound\n\n\r https://twitter.com/theDrunkRobots 👈 follow the Robots\n https://discord.gg/8xAqNKdbFR 👈 join our Discord","Name":"🤖 Drunk Robot #156 - Karsten","VectorState":"Success","VectorID":"444026885366615631","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"9f4ade5c-004b-45f2-90eb-17dacaad4999","TokenID":"3825","ImageURL":"https://drunkrobots.net/nft/robot/3825.png","IPFSImageURL":""},{"ID":"55006203-4977-4d8d-94e2-f5627f9f016c","TokenID":"8564","ImageURL":"https://drunkrobots.net/nft/robot/8564.png","IPFSImageURL":""},{"ID":"d364b587-7b39-4aa8-9883-42ff2c7de466","TokenID":"7408","ImageURL":"https://drunkrobots.net/nft/robot/7408.png","IPFSImageURL":""}],"SiblingsNum":4,"Distance":0.8462875},{"ID":"5b7846a0-301f-4cd2-b890-9234fe9a793c","ChainType":"Ethereum","ChainID":"1","Contract":"0xe41bEa6888f771A0C16d7188284522B76C135252","TokenType":"ERC721","TokenID":"7975","Owner":"","URI":"https://api.elysiumshell.xyz/es/7975","URIType":"http","ImageURL":"https://api.elysiumshell.xyz/es/7975/image/half","VideoURL":"https://api.elysiumshell.xyz/es/7975/image/full?format=html","Description":"Escape from the broken reality. Ascend to the above.\n\nLet go of the old me. Transform into a new shell.\n\nWhen all consciousness gathers, the ultimate consensus is formed.\n\n[Full Body](https://api.elysiumshell.xyz/es/7975/image/full)\n\n[Half Body](https://api.elysiumshell.xyz/es/7975/image/half)","Name":"E-Shell #1570","VectorState":"Success","VectorID":"444026885366617073","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.8463441},{"ID":"105c37e7-3eff-4660-98d2-6a0483f58a28","ChainType":"Ethereum","ChainID":"1","Contract":"0xD7B397eDad16ca8111CA4A3B832d0a5E3ae2438C","TokenType":"ERC1155","TokenID":"2174","Owner":"","URI":"ipfs://QmUfGyJx8phhTGbXSbTtjjX4x5UEytu5tVkSmf4DPF8WFe/2174","URIType":"ipfs","ImageURL":"ipfs://QmPoXb2zs9ehCfnUz5xU1mn3e4epekmm8YHLF8yDn6JyxG/2174.png","VideoURL":"","Description":"The year is 2050. Humans are an interplanetary species and have all but abandoned the post-apocalyptic shatters of society on earth. Cats have taken over. One crime-ridden, nondescript inner city is inhabited by a group of cats collectively known as the Gutter Cats. Far underground, in the bowels of the Autonomous Zone, the Gutter Rats roam, sustaining themselves on the decaying crumbs of the extravagant lifestyles above. Discontent with their subservient position in the Gutter, the conniving rats lie in wait, scheming to one day claim the throne from the Gutter Cat Gang.","Name":"Gutter Rat #2174","VectorState":"Success","VectorID":"444026885366616789","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.8506723},{"ID":"11a68a65-a0d6-491b-badf-c74c48c8bcf0","ChainType":"Ethereum","ChainID":"1","Contract":"0xcc51892FFFf49B3169D06229EB3ee6d00eA1E554","TokenType":"ERC721","TokenID":"457","Owner":"","URI":"ipfs://QmcDN2yCrgakKY6tAXd7D2cnHSF37R8s9SzviesARtA2r9/457","URIType":"ipfs","ImageURL":"https://togglesuniverse.s3.ap-northeast-2.amazonaws.com/image/505.jpg","VideoURL":"","Description":"Toggles is the character IP of web2.0 that has ventured into the world of 3D digital collectibles in web3.0. \nThere are a total of 6000 Toggles, which are virtual trainees aspiring to become singers. They make their debut through collaborations with web2.0 artists. \nEach Toggles character belongs to one of seven music genres: EDM, HIPHOP, JAZZ, REGGAE, ROCK, FOLK, and DISCO. They all reside in a place called NUT-HEXAGON.","Name":"Toggles #505","VectorState":"Success","VectorID":"444026885366616001","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"521fcee1-1ba1-4b78-9eba-aeef02412eee","TokenID":"496","ImageURL":"https://togglesuniverse.s3.ap-northeast-2.amazonaws.com/image/5178.jpg","IPFSImageURL":""},{"ID":"134a7a41-b209-4cd4-8278-6efab27b7e6d","TokenID":"446","ImageURL":"https://togglesuniverse.s3.ap-northeast-2.amazonaws.com/image/493.jpg","IPFSImageURL":""}],"SiblingsNum":3,"Distance":0.87897956},{"ID":"16f04b5b-6755-441e-b4e3-05070fa46051","ChainType":"Ethereum","ChainID":"1","Contract":"0x932261f9Fc8DA46C4a22e31B45c4De60623848bF","TokenType":"ERC721","TokenID":"179180","Owner":"","URI":"https://dna.zerion.io/api/v1/avatars/onepointo/179180","URIType":"http","ImageURL":"https://zerion-dna.s3.us-east-1.amazonaws.com/onepointo/32c32f197e36340893c970692729ca8b1450191a.png","VideoURL":"","Description":"Zerion DNA is a first of a kind Dynamic NFT Avatar – a generative, living NFT that evolves over time to reflect your digital identity. Every wallet action renders unique attributes so that no two NFTs are alike. Wallet actions that will cause the NFT to evolve include buying and selling tokens, interacting with different networks and exploring the Zerion app.\n\nTo claim your DNA, create a Zerion Wallet via our Zerion iOS and Android apps.","Name":"#179180","VectorState":"Success","VectorID":"444026885366618627","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"2615289f-f7a3-4d66-93e0-acb42875149c","TokenID":"179183","ImageURL":"https://zerion-dna.s3.us-east-1.amazonaws.com/onepointo/b193091ac361a24944ac949195f3a2f9d5cdf93f.png","IPFSImageURL":""},{"ID":"99d84133-4768-4446-94b2-772fb9af803c","TokenID":"179181","ImageURL":"https://zerion-dna.s3.us-east-1.amazonaws.com/onepointo/05a5e678c6eccdc5a568875a532286fd1975eb0a.png","IPFSImageURL":""},{"ID":"23633cd3-d445-47fa-a2f6-ce538b23b26b","TokenID":"179182","ImageURL":"https://zerion-dna.s3.us-east-1.amazonaws.com/onepointo/fada6585d451a446c581d23e16d7c8f3ad034b40.png","IPFSImageURL":""}],"SiblingsNum":4,"Distance":0.8905762},{"ID":"73b75db1-1076-4d48-8375-081a76a5c3fe","ChainType":"Ethereum","ChainID":"1","Contract":"0x947e50732366125811e131a116f5AaDa6315Af5C","TokenType":"ERC721","TokenID":"1120","Owner":"","URI":"ipfs://QmYg8XbSnNRCscgGvwk48KxqN7mys4TZ7VijLKWtUT89ok/1120","URIType":"ipfs","ImageURL":"ipfs://QmWdG76gDiVHpK5mDJRKKQr4PLtESqrMMBk4wYrXGfQ6GJ/1120.png","VideoURL":"","Description":"Sproto Cats are the embodiment of XCatge, serving as pets for the Sproto Gremlins.","Name":"Sproto Cats #1120","VectorState":"Success","VectorID":"444026885366616083","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.890836},{"ID":"7c9cec29-4b24-4dae-9f7c-ee23b43434fd","ChainType":"Ethereum","ChainID":"1","Contract":"0x33F4085C68546fAf28d8d0CE20BBbc2B3B59C2B4","TokenType":"ERC721","TokenID":"985","Owner":"","URI":"https://freemintlabs.xyz/meta2/985","URIType":"http","ImageURL":"https://freemintlabs.xyz/zombieape/985.png","VideoURL":"","Description":"ZombieApes","Name":"ZombieApes #985","VectorState":"Success","VectorID":"444026885366618155","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"71bdb02e-1626-4bf7-9105-9603aff359a5","TokenID":"987","ImageURL":"https://freemintlabs.xyz/zombieape/987.png","IPFSImageURL":""},{"ID":"b1689881-5c27-46d7-beb3-ef42dc9b17cf","TokenID":"984","ImageURL":"https://freemintlabs.xyz/zombieape/984.png","IPFSImageURL":""},{"ID":"81e45c9b-a3af-4d7c-aff6-ec91cbb1b5f4","TokenID":"981","ImageURL":"https://freemintlabs.xyz/zombieape/981.png","IPFSImageURL":""},{"ID":"0b36239f-2069-4fd2-ab63-fe6ba386ecc7","TokenID":"982","ImageURL":"https://freemintlabs.xyz/zombieape/982.png","IPFSImageURL":""},{"ID":"a0178eb7-b186-4916-af46-a4f3c6c916ea","TokenID":"983","ImageURL":"https://freemintlabs.xyz/zombieape/983.png","IPFSImageURL":""},{"ID":"3cd63170-9ba1-48a6-b6eb-8ad9060d8c58","TokenID":"986","ImageURL":"https://freemintlabs.xyz/zombieape/986.png","IPFSImageURL":""},{"ID":"411ae42d-ac72-4408-a763-2fb9b0629f5a","TokenID":"988","ImageURL":"https://freemintlabs.xyz/zombieape/988.png","IPFSImageURL":""},{"ID":"0ab610ff-8658-47ee-be27-38963517704e","TokenID":"989","ImageURL":"https://freemintlabs.xyz/zombieape/989.png","IPFSImageURL":""},{"ID":"579886bb-3b9e-4bfc-b245-dc460873140a","TokenID":"991","ImageURL":"https://freemintlabs.xyz/zombieape/991.png","IPFSImageURL":""},{"ID":"3b316150-9bf5-40d6-9b6b-2ce0673f9089","TokenID":"990","ImageURL":"https://freemintlabs.xyz/zombieape/990.png","IPFSImageURL":""}],"SiblingsNum":12,"Distance":0.90559924},{"ID":"6b958e00-a2e7-49d0-8529-1258a89b77e2","ChainType":"Ethereum","ChainID":"1","Contract":"0xc92cedDfb8dd984A89fb494c376f9A48b999aAFc","TokenType":"ERC721","TokenID":"3360","Owner":"","URI":"https://creature.mypinata.cloud/ipfs/QmVDNzQNuD5jBKHmJ2nmVP35HsXUqhGRX9V2KVHvRznLg8/3360","URIType":"ipfs-gateway","ImageURL":"https://creature.mypinata.cloud/ipfs/QmeZGc1CL3eb9QJatKXTGT7ekgLMq9FyZUWckQ4oWdc53a/3360.jpg","VideoURL":"","Description":"Welcome to The Creature World. You have arrived in a nearby magical dimension of love, divine intervention, and possibility. 10,000 unique Creatures are here to guide you on this journey. Follow their lead. Created with love by NYC-based artist Danny Cole. www.creature.world.","Name":"Creature #3360","VectorState":"Success","VectorID":"444026885366619075","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"383158cc-6475-438a-8d98-0614f6ece699","TokenID":"5376","ImageURL":"https://creature.mypinata.cloud/ipfs/QmeZGc1CL3eb9QJatKXTGT7ekgLMq9FyZUWckQ4oWdc53a/5376.jpg","IPFSImageURL":""}],"SiblingsNum":2,"Distance":0.90878034},{"ID":"ed7878fa-97be-4a4a-92db-11132df0bc6f","ChainType":"Ethereum","ChainID":"1","Contract":"0xfE4CB03F2A4Cf5A4074F6C8F77894A814782FdBA","TokenType":"ERC721","TokenID":"1779","Owner":"","URI":"ipfs://QmbPtEwTBVKbMV6TJuitTF95gdYfxChZfa6cXJjsYcyEPG/1779.json","URIType":"ipfs","ImageURL":"ipfs://QmSR6iGmxEDJYgy4LhCoSWUFQCFagPT8ZrTe3KRibrGJww/1779.png","VideoURL":"","Description":"","Name":"Football Ape Fan Club 1779","VectorState":"Success","VectorID":"444026885366615757","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.91910034},{"ID":"b356cc94-5e4e-4278-a579-ee6e5912ebbb","ChainType":"Ethereum","ChainID":"1","Contract":"0xedF6d3C3664606Fe9EE3a9796d5CC75E3B16e682","TokenType":"ERC721","TokenID":"4267","Owner":"","URI":"ipfs://QmeizFQc2x29rPLFFsXWVYNqqXRn7rHwK5YS9jbWfhcQ4m/454.json","URIType":"ipfs","ImageURL":"ipfs://QmYsgfkwQRnvxf55zGTDjmqE8K7PkYTaGSvu2oQKjvwktS/454.png","VideoURL":"","Description":"","Name":"Fat Cat #4475","VectorState":"Success","VectorID":"444026885366615785","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"b7964ec3-d360-4556-918d-397967ca077b","TokenID":"4224","ImageURL":"ipfs://QmYsgfkwQRnvxf55zGTDjmqE8K7PkYTaGSvu2oQKjvwktS/411.png","IPFSImageURL":""}],"SiblingsNum":2,"Distance":0.9265545},{"ID":"bae15f71-b6e9-426f-9b68-3f5e02dc0921","ChainType":"Ethereum","ChainID":"1","Contract":"0xB9D5551a31ceeB95E58180583262164012728C16","TokenType":"ERC721","TokenID":"1371","Owner":"","URI":"https://app.bueno.art/api/contract/ntGfdVCVWAHiybQxLhzXP/chain/1/metadata/1371","URIType":"http","ImageURL":"https://assets.bueno.art/images/e92aee60-5454-438c-9f8c-b23ff260c7bb/default/1371?s=6e3a6e92d139e577ab98171f6f0bb2cf","VideoURL":"","Description":"A lifestyle brand for mental wellness.","Name":"MONK #1371","VectorState":"Success","VectorID":"444026885366616481","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"7665384e-75bb-4880-80f9-79e3e489912f","TokenID":"5352","ImageURL":"https://assets.bueno.art/images/e92aee60-5454-438c-9f8c-b23ff260c7bb/default/5352?s=aac4991ef9883240cfbc0e3b749dc44b","IPFSImageURL":""},{"ID":"64dce656-b8d3-48eb-847c-e61b2055e115","TokenID":"4665","ImageURL":"https://assets.bueno.art/images/e92aee60-5454-438c-9f8c-b23ff260c7bb/default/4665?s=596c672b9776487eeacc93b9408baf66","IPFSImageURL":""},{"ID":"a692f57b-97c4-4643-89db-ad3a398b50f7","TokenID":"4244","ImageURL":"https://assets.bueno.art/images/e92aee60-5454-438c-9f8c-b23ff260c7bb/default/4244?s=1142ce9b3ad6e52ef772ea67a20af13e","IPFSImageURL":""}],"SiblingsNum":4,"Distance":0.9418286},{"ID":"0266b6c9-fdd2-4e0d-8911-766613489180","ChainType":"Ethereum","ChainID":"1","Contract":"0xc934e429360e2b55d5BA95751d31D3da1E3410cD","TokenType":"ERC721","TokenID":"7800","Owner":"","URI":"https://trippies.blob.core.windows.net/cosmics-metadata/7800","URIType":"http","ImageURL":"https://trippies.blob.core.windows.net/cosmics-images-large/7800.png","VideoURL":"","Description":"The Cosmic Trippies are part of the Paradise Trippies collection and are Paradise Island’s latest residents. What are their intentions in Paradise? They are visitors with a mission.","Name":"Cosmic Trippies #7800","VectorState":"Success","VectorID":"444026885366616451","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[{"ID":"7da711ae-ee73-460a-bc75-fd510249336e","TokenID":"8757","ImageURL":"https://trippies.blob.core.windows.net/cosmics-images-large/8757.png","IPFSImageURL":""}],"SiblingsNum":2,"Distance":0.9503592},{"ID":"11ba7672-72ce-41ac-bc81-bb24a928a0bc","ChainType":"Ethereum","ChainID":"1","Contract":"0xd9F092BdF2b6eaF303fc09cc952e94253AE32fae","TokenType":"ERC721","TokenID":"5313","Owner":"","URI":"https://ipfs.io/ipfs/QmTfVckiaAxjaEwZTa5PmLa1LCoaYgTZbLW38Cobs9xRDv/5313","URIType":"ipfs-gateway","ImageURL":"https://ipfs.io/ipfs/QmcJygK3B27ZnebsPjZ5DDJF4G5FhTPJcQHqc4mm743Su5/5313.png","VideoURL":"","Description":"Baby Doge Army is a collection of 10,000 adoptable baby doges. A unique digital art collection waiting to be rescued on the Ethereum Blockchain. Each one has been generated then hand-groomed by our team to be fit for adoption. Join us on our mission and have a good time. Having a Baby Doge grants you creative and commercial rights, as well as inclusion in the community, plus feel great knowing your NFT helped make a difference to save dogs in need.","Name":"Baby Doge #5313","VectorState":"Success","VectorID":"444026885366618797","Remark":"","IPFSImageURL":"","ImageSnapshotID":"","TransfersNum":1,"SiblingTokens":[],"SiblingsNum":1,"Distance":0.95439065}],"StorageKey":"4f3cba26-2e60-4b48-870c-171eafc09f0a","Page":1,"TotalPages":2,"TotalTokens":23,"Limit":20}`
	// w.Write([]byte(ret))
	// return

	startT := time.Now()
	respBody := []byte{}
	var err error
	var errMsg string
	defer func() {
		if errMsg != "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}

		_, err = w.Write(respBody)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("write response body fail, %v", err)))
		}
	}()

	_limit := r.FormValue(LimitFeild)
	limit, err := strconv.ParseUint(_limit, 10, 32)

	if err != nil {
		errMsg = fmt.Sprintf("failed to parse feild Limit %v, %v", _limit, err)
		return
	}

	// judge weather filesize exceed max-size
	err = r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		errMsg = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		return
	}

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	// convert to vector
	vector, err := ImgReqConvertVector(r)
	if err != nil {
		errMsg = fmt.Sprintf("image convert fail, %v", err)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish convert to vector %v ms", inT.UnixMilli()-startT.UnixMilli())

	token.UseCloudProxyCC()
	resp, err := token.Search(context.Background(), &rankerproto.SearchTokenRequest{
		Vector: vector,
		Limit:  uint32(limit),
	})

	if err != nil {
		errMsg = fmt.Sprintf("search fail, %v", err)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish query id %v ms", inT.UnixMilli()-startT.UnixMilli())

	buff := bytes.NewBuffer([]byte{})
	err = pbJsonMarshaler.Marshal(buff, resp)
	if err != nil {
		errMsg = fmt.Sprintf("marshal result fail, %v", err)
		return
	}

	respBody = buff.Bytes()
}

// TODO: this method from nft-meta/pkg/imageconvert/utils.go that will be reconstruct
// converte http request with image file to vector
func ImgReqConvertVector(r *http.Request) ([]float32, error) {
	// get file info
	file, handler, err := r.FormFile(UploadFileFeild)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// rewrite file to new request-body
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)
	fileWriter, err := bodyWriter.CreateFormFile(UploadFileFeild, handler.Filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	bodyWriter.Close()
	ICServer := fmt.Sprintf("%v:%v",
		config.GetConfig().Transform.Domain,
		config.GetConfig().Transform.HTTPPort,
	)
	icURL := fmt.Sprintf("http://%v/v1/transform/file", ICServer)

	res, err := http.Post(icURL, bodyWriter.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body1, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parse response
	resp := &Img2VectorResp{}
	err = json.Unmarshal(body1, resp)
	if err != nil {
		return nil, err
	}

	return resp.Vector, nil
}

type Server struct {
	entranceproto.UnimplementedManagerServer
}

func (s *Server) SearchPage(ctx context.Context, in *rankerproto.SearchPageRequest) (*rankerproto.SearchResponse, error) {
	// ret := `{
	// 	"Infos": [
	// 		{
	// 			"ID": "3b9a4210-1a78-46b4-9553-d6edcfcc22c2",
	// 			"ChainType": "Ethereum",
	// 			"ChainID": "1",
	// 			"Contract": "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D",
	// 			"TokenType": "ERC721",
	// 			"TokenID": "2783",
	// 			"Owner": "",
	// 			"URI": "ipfs://QmeSjSinHpPnmXmspMjwiXyN6zS4E9zccariGR3jxcaWtq/2783",
	// 			"URIType": "ipfs",
	// 			"ImageURL": "ipfs://QmNUUwVkKHjK4zgUZJ4YDrH48HoT3vahcJt2fS4A8QHcvj",
	// 			"VideoURL": "",
	// 			"Description": "",
	// 			"Name": "",
	// 			"VectorState": "Success",
	// 			"VectorID": "444026885366617813",
	// 			"Remark": "",
	// 			"IPFSImageURL": "",
	// 			"ImageSnapshotID": "",
	// 			"TransfersNum": 1,
	// 			"SiblingTokens": [
	// 				{
	// 					"ID": "8fd39cb5-aab7-403a-a8de-01ed9ca4e272",
	// 					"TokenID": "8970",
	// 					"ImageURL": "ipfs://QmfPwwT9uCMakYUBrHi4B8V3MyrwqUV5eSCD8C3ScCoqT6",
	// 					"IPFSImageURL": ""
	// 				}
	// 			],
	// 			"SiblingsNum": 2,
	// 			"Distance": 0.9635749
	// 		},
	// 		{
	// 			"ID": "de47ddc2-2e5c-411c-a192-9c9242cdc2eb",
	// 			"ChainType": "Ethereum",
	// 			"ChainID": "1",
	// 			"Contract": "0xBd3531dA5CF5857e7CfAA92426877b022e612cf8",
	// 			"TokenType": "ERC721",
	// 			"TokenID": "1509",
	// 			"Owner": "",
	// 			"URI": "ipfs://bafybeibc5sgo2plmjkq2tzmhrn54bk3crhnc23zd2msg4ea7a4pxrkgfna/1509",
	// 			"URIType": "ipfs",
	// 			"ImageURL": "ipfs://QmNf1UsmdGaMbpatQ6toXSkzDpizaGmC9zfunCyoz1enD5/penguin/1509.png",
	// 			"VideoURL": "",
	// 			"Description": "A collection 8888 Cute Chubby Pudgy Penquins sliding around on the freezing ETH blockchain.",
	// 			"Name": "Pudgy Penguin #1509",
	// 			"VectorState": "Success",
	// 			"VectorID": "444026885366618973",
	// 			"Remark": "",
	// 			"IPFSImageURL": "",
	// 			"ImageSnapshotID": "",
	// 			"TransfersNum": 1,
	// 			"SiblingTokens": [
	// 				{
	// 					"ID": "dcb0f4ae-e1da-4713-9007-c4b70e4a9db6",
	// 					"TokenID": "1213",
	// 					"ImageURL": "ipfs://QmNf1UsmdGaMbpatQ6toXSkzDpizaGmC9zfunCyoz1enD5/penguin/1213.png",
	// 					"IPFSImageURL": ""
	// 				}
	// 			],
	// 			"SiblingsNum": 2,
	// 			"Distance": 0.9694272
	// 		},
	// 		{
	// 			"ID": "ddb6de79-0707-4730-88e8-0fb4f1e59df6",
	// 			"ChainType": "Ethereum",
	// 			"ChainID": "1",
	// 			"Contract": "0xA3F7250306Dbb856D8d312f93029be73343939aF",
	// 			"TokenType": "ERC721",
	// 			"TokenID": "1609",
	// 			"Owner": "",
	// 			"URI": "https://thesadtimes.com/api/metadata/1609",
	// 			"URIType": "http",
	// 			"ImageURL": "https://thesadtimescdn.com/metadata/1609.gif",
	// 			"VideoURL": "https://thesadtimescdn.com/metadata/1609.mp4",
	// 			"Description": "https://thesadtimes.com/tokens/1609",
	// 			"Name": "Sheep #1609",
	// 			"VectorState": "Success",
	// 			"VectorID": "444026885366616455",
	// 			"Remark": "",
	// 			"IPFSImageURL": "",
	// 			"ImageSnapshotID": "",
	// 			"TransfersNum": 1,
	// 			"SiblingTokens": [],
	// 			"SiblingsNum": 1,
	// 			"Distance": 0.97734225
	// 		}
	// 	],
	// 	"StorageKey": "4f3cba26-2e60-4b48-870c-171eafc09f0a",
	// 	"Page": 2,
	// 	"TotalPages": 2,
	// 	"TotalTokens": 23,
	// 	"Limit": 20
	// }`
	// retIn := &rankerproto.SearchResponse{}
	// json.Unmarshal([]byte(ret), retIn)
	// return retIn, nil

	token.UseCloudProxyCC()
	return token.SearchPage(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entranceproto.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entranceproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
