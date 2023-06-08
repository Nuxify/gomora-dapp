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
	greeter "gomora-dapp/infrastructures/smartcontracts/greeter"
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
	m                       sync.Mutex
	k                       *kernel
	containerOnce           sync.Once
	mysqlDBHandler          *mysql.MySQLDBHandler
	EthHttpClient           *ethclient.Client
	EthWsClient             *ethclient.Client
	GreeterContractInstance *greeter.Greeter
)

// ================================= REST ===================================
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

// ==========================================================================
func NFTCommandServiceDI() *nftService.NFTCommandService {
	service := &nftService.NFTCommandService{
		GreeterContractInstance: GreeterContractInstance,
	}

	return service
}

func (k *kernel) nftCommandServiceContainer() *nftService.NFTCommandService {
	return NFTCommandServiceDI()
}

func (k *kernel) nftQueryServiceContainer() *nftService.NFTQueryService {
	service := &nftService.NFTQueryService{
		GreeterContractInstance: GreeterContractInstance,
	}

	return service
}

func registerHandlers() {
	var err error

	// connect to blockchain
	EthHttpClient, err = ethclient.Dial(os.Getenv("ETH_MAINNET_HTTP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	EthWsClient, err = ethclient.Dial(os.Getenv("ETH_MAINNET_WS_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	// load smart contracts
	GreeterContractAddress = common.HexToAddress(os.Getenv("ETH_MAINNET_GREETER_CONTRACT_ADDRESS"))
	GreeterContractABI, err = abi.JSON(strings.NewReader(string(greeter.GreeterABI)))
	if err != nil {
		log.Fatal(err)
	}
	GreeterContractInstance, err = greeter.NewGreeter(GreeterContractAddress, EthWsClient)
	if err != nil {
		log.Fatal(err)
	}

	// run event listener
	go GreeterEventListener()
	go GreeterPollFilter(os.Getenv("ETH_MAINNET_HTTP_ENDPOINT")) // TODO: check if chain rpc supports eth_filterChanges
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
