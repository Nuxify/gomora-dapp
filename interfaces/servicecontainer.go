/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"gomora-dapp/infrastructures/database/mysql"
	smartcontract "gomora-dapp/infrastructures/smartcontracts"
	nftService "gomora-dapp/module/nft/infrastructure/service"
	nftREST "gomora-dapp/module/nft/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterNFTRESTCommandController() nftREST.NFTCommandController
	RegisterNFTRESTQueryController() nftREST.NFTQueryController
}

type kernel struct{}

var (
	m                              sync.Mutex
	k                              *kernel
	containerOnce                  sync.Once
	mysqlDBHandler                 *mysql.MySQLDBHandler
	EthHttpClient                  *ethclient.Client
	EthWsClient                    *ethclient.Client
	SampleContractContractInstance *smartcontract.Smartcontracts
)

//================================= REST ===================================
// RegisterNFTRESTCommandController performs dependency injection to the RegisterNFTRESTCommandController
func (k *kernel) RegisterNFTRESTCommandController() nftREST.NFTCommandController {
	service := k.nftCommandServiceContainer()

	controller := nftREST.NFTCommandController{
		NFTCommandServiceInterface: service,
	}

	return controller
}

// RegisterNFTRESTQueryController performs dependency injection to the RegisterNFTRESTQueryController
func (k *kernel) RegisterNFTRESTQueryController() nftREST.NFTQueryController {
	service := k.nftQueryServiceContainer()

	controller := nftREST.NFTQueryController{
		NFTQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) nftCommandServiceContainer() *nftService.NFTCommandService {
	service := &nftService.NFTCommandService{}

	return service
}

func (k *kernel) nftQueryServiceContainer() *nftService.NFTQueryService {
	service := &nftService.NFTQueryService{
		SampleContractContractInstance: SampleContractContractInstance,
	}

	return service
}

func registerHandlers() {
	var err error

	// connect to polygon
	EthHttpClient, err = ethclient.Dial(os.Getenv("ETH_MAINNET_HTTP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	EthWsClient, err = ethclient.Dial(os.Getenv("ETH_MAINNET_WS_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	// legend nft smart contract instance
	SampleContractContractAddress = common.HexToAddress(os.Getenv("ETH_MAINNET_SAMPLE_CONTRACT_ADDRESS"))
	SampleContractContractInstance, err = smartcontract.NewSmartcontracts(SampleContractContractAddress, EthWsClient)
	if err != nil {
		log.Fatal(err)
	}

	// load abi json
	SampleContractContractABI, err = abi.JSON(strings.NewReader(string(smartcontract.SmartcontractsABI)))
	if err != nil {
		log.Fatal(err)
	}

	// run event watcher
	go SampleContractEventWatcher()
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
