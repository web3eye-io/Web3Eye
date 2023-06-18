# Web3Eye

[![Test](https://github.com/web3eye-io/Web3Eye/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/web3eye-io/Web3Eye/actions/workflows/main.yml)

![web3eye](doc/picture/web3eye.png)

Currently in the world of NFT, many methods for obtaining information about blockchain are complex and difficult for users to use, making it difficult for the public to obtain information.
Furthermore, various blockchain project data is in a fragmented state currently, making it even more difficult to obtain or organize information.
Web3Eye is a search engine that aggregates historical NFT transaction records, and provides multi-chain aggregation search for NFT assets.

## Architecture

Middleware:
**MySql** stores task information and NFT data relationships
**Kafka** is mainly used for task assignment
**Redis** is used to cache hot information queried from MySql (a local principle in computers, which reduces the pressure on MySql)
**Milvus** is used to store vector data and provide vector search

Microservice modules:
**NFT-Meta** maintains block dumping tasks, stores NFT transactions, NFT assets, and NFT corresponding Contract information
**Block-ETL** is responsible for interacting with the blockchain node, obtaining NFT transfer logs, analyzing corresponding Token information and Contract information
**Image-Converter** converts images into vectors.

![Architecture](doc/picture/archi.jpg)

Among the three main microservice modules, NFT-Meta is responsible for providing search, information storage query, task distribution and other functions, while the other two modules are more focused on obtaining and processing tasks. Image-Converter not only processes tasks sent by NFT-Meta from Kafka, but also provides HTTP service support to directly request vectors, mainly used to provide image search services. Block-ETL does not provide external interfaces, only receives tasks and submits tasks.

### Module Design

#### Image-Converter

Currently, the service mainly provides vector conversion operations for common image formats such as JPG, JPEG, and PNG. Other image resources such as GIF and Base64 are currently not supported.
After the service is started, there are two threads. One is responsible for providing vector conversion through HTTP interface, providing synchronous vector conversion method, supporting both URL and file methods. The other is responsible for getting vector conversion tasks from Kafka, converting them, and putting them into Kafka for NFT-Meta to obtain and store in Milvus and the database.

#### Block-ETL

Currently, only Ethereum is supported, so the following description is based on the context of ETH. Currently, only standard ERC721 and ERC1155 are supported, while support for other NFT gameplay will be provided in the future.
Pull transfer information from the logs in the full node of the blockchain (which holds all block data, referred to here as a wallet node), parse out the NFT transaction, Token, and Contract information. However, some Token information cannot be parsed into asset information (such as image descriptions and image addresses) due to the existence of Swap contracts and non-standard NFT contracts.
The granularity of the tasks obtained from NFT-Meta is the block height. All transfer logs from a block height are obtained, and each transfer is recorded. Then, Token information is searched from the transfer information, because multiple transfers may correspond to the same Token. Therefore, the database is first queried to see if the Token exists. If it does not exist, the TokenURI will be requested from the wallet node, and the corresponding Contract will also be checked. When checking whether the Token and Contract exist here, Redis is actually checked first for records. If there is no record, a query is made to the database, and a record is created in Redis when the information is found.

![Data Relationship](doc/picture/transfer-token-contract.jpg)

When parsing the transfer logs of a block, most of the information can be obtained from the wallet node. However, the information carried by the TokenURI needs to be obtained from the Internet or IPFS, or directly stored on the blockchain as Base64, SVG, and so on. Currently, the work of parsing the TokenURI belongs to this module, and it is planned to be independently developed into a separate module in the future. Because such parsing work is time-consuming and labor-intensive, Block-ETL tries to interact only with the wallet node and only do the job of storing on-chain data as much as possible.

#### NFT-Meta

Tasks assigned to the other two modules are issued by NFT-Meta, which stores the data processed by the other two modules and provides search capabilities externally. Currently, this module may be a bit bloated, for example, the search functionality can be separated out, and issues related to other modules directly interacting with the database will be further considered. However, the search function has been initially placed in this module because of its limited functionality.

NFT-Meta mainly maintains four tables:
1 **Transfers** - NFT transaction records
2 **Tokens** - NFT asset information
3 **Contracts** - NFT contract information
4 **SyncTasks** - synchronization tasks

NFT-Meta provides both GRPC and HTTP API interfaces, with GRPC primarily provided to internal microservice modules and HTTP provided externally. Vector data is mainly stored in Milvus, while relational data is mainly stored in MySql. Data in Milvus and MySql are linked by ID provided by Milvus.
The structure in Milvus is:

{
    ID: 13125
    Vector: [0.234,2.923,...]
}

The ID field in Milvus is linked in MySql, for example:

{
...
    ID: 29aa144d-beb0-4d25-b7bb-95587fe06ba4
    VectorID: 13125
    VectorState: Success
...
}

### Workflow

How to search for NFT assets? The current search methods are mainly through the contract and TokenID. However, Web3Eye adopts similarity search. NFT assets have diverse forms, mostly presented as unstructured data such as images, audios, videos, etc., which can be transformed into vectors that can be searched. Currently, only image-to-image search is supported, and follow-up support for other forms of NFT assets will be provided.
With the search of unstructured data in place, the next step is to aggregate NFT data. The transfer logs (logs of NFT transactions) are obtained from the wallet node, and Token and Contract information are analyzed from them. Among the three data related to NFTs, transfer and Contract information can be obtained relatively easily, while parsing Token is a bit more complex.

#### Search NFT by Image

Image-to-image search is the process of calculating the distance between feature vectors, which involves calculating the distance between a given feature vector and a set of existing feature vectors and then selecting the top N results based on the smallest distances.
The existing feature vectors are derived from the image data of NFTs and stored in the vector database.
The image used for searching also needs to be converted into a vector for distance calculation with the data in the vector database.

![Image to vector conversion (the sample image does not represent the actual converted vector result, only for conceptual demonstration purposes)](doc/picture/image-to-vector.jpg)

Searching for an image in Web3Eye generally goes through four stages.

![Search Image by Image](doc/picture/pictrue-search.jpg)

The process of searching for an image in Web3Eye typically involves four stages:

1. The user sends a request with a file (image file) to NFT-Meta.
2. NFT-Meta forwards the request to Image-Converter to convert the file into a vector.
3. NFT-Meta uses the vector to search for similar vectors in Milvus and returns the vector ID.
4. NFT-Meta retrieves the Token information corresponding to the ID from MySql and returns it to the user.

#### Retrieving Token Information
The main fields of a Token are as follows:

{
"Contract"
"TokenID"
"TokenType"
"URI"
"ImageURL"
"VectorID"
}

Among these fields, Contract, TokenID, and URI can be obtained directly from the wallet node. TokenType and ImageURL can be derived from URI analysis. VectorID, on the other hand, needs to be inserted into a Token record in NFT-Meta after all other information has been obtained. The vector conversion task is then added to a queue and waits for Image-Converter to process it. Once the conversion is completed, the result is placed in a queue and waits for NFT-Meta to update the VectorID field.

#### Task Dispatch

There are currently two places where Kafka is used. The first is to send vector conversion tasks to Image-Converter, and the second is to send block heights that need to be parsed to Block-ETL.
The figure below shows the process of handling a large number of vector conversion tasks with low time requirements using asynchronous processing, as network bandwidth and computing resources are consumed during vector conversion.

![Image2Vector task dispatch](doc/picture/to-vector-task.jpg)

However, when searching, the HTTP vector conversion method provided by Image-Converter is directly requested to improve response speed.
Block-ETL is mainly responsible for analyzing data from each block height and placing it in NFT-Meta.
There are two processes for the tasks that Block-ETL obtains (block heights to be synchronized):
The administrator requests NFT-Meta to establish a synchronization task, which includes the start block, end block, and current block.
Block-ETL periodically checks (requests from NFT-Meta) whether there are topics that need to be synchronized, listens, and consumes if there are.
The first process is simple, just adding a task record to the database. The second process requires Block-ETL to actively query for tasks to be synchronized, and to trigger NFT-Meta to send data to Kafka. Although NFT-Meta sends the data, the initiative of consumption is given to Block-ETL, while making NFT-Meta more stateless.

![Image2Vector task dispatch](doc/picture/block-etl-task.jpg)

## Deployment

[k8s-deploy](doc/deploy/k8s-deploy.md)

## Version Plan

[0.1.0](doc/feature/0.1.0.md)

[100.0.0](doc/feature/100.0.0.md)
