// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_uint256)49_storage\"}],\"types\":{\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c80637f46ddb2116100505780637f46ddb2146100f1578063927ede2d1461013d578063aa5574521461016457600080fd5b80633687011a1461007757806354fd4d501461008c578063761f4493146100de575b600080fd5b61008a610085366004610fc2565b610177565b005b6100c86040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d591906110b0565b60405180910390f35b61008a6100ec3660046110c3565b610223565b6101187f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d5565b6101187f000000000000000000000000000000000000000000000000000000000000000081565b61008a61017236600461115b565b61078a565b333b1561020b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b61021b8686333388888888610846565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561034157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610305573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032991906111d2565b73ffffffffffffffffffffffffffffffffffffffff16145b6103cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f7468657220627269646765006064820152608401610202565b3073ffffffffffffffffffffffffffffffffffffffff881603610472576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c66000000000000000000000000000000000000000000006064820152608401610202565b61049c877f74259ebf00000000000000000000000000000000000000000000000000000000610de4565b610528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e74000000000000000000006064820152608401610202565b8673ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610573573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059791906111d2565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610677576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204b726f6d61204d696e7461626c6520455243373231206c6f6360648201527f616c20746f6b656e000000000000000000000000000000000000000000000000608482015260a401610202565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b1580156106e757600080fd5b505af11580156106fb573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107799493929190611238565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff851661082d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f742062652061646472657373283029000000000000000000000000000000006064820152608401610202565b61083d8787338888888888610846565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff87166108e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f7420626520616464726573732830290000000000000000000000000000006064820152608401610202565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa158015610954573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097891906111d2565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610a32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e657200006064820152608401610202565b60008873ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa391906111d2565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b60576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c75650000000000000000006064820152608401610202565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610bd057600080fd5b505af1158015610be4573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610c0f9796959493929190611278565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610d24907f00000000000000000000000000000000000000000000000000000000000000009085908a906004016112d5565b600060405180830381600087803b158015610d3e57600080fd5b505af1158015610d52573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610dd09493929190611238565b60405180910390a450505050505050505050565b6000610def83610e07565b8015610e005750610e008383610e6c565b9392505050565b6000610e33827f01ffc9a700000000000000000000000000000000000000000000000000000000610e6c565b8015610e665750610e64827fffffffff00000000000000000000000000000000000000000000000000000000610e6c565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015610f24575060208210155b8015610f305750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f5d57600080fd5b50565b803563ffffffff81168114610f7457600080fd5b919050565b60008083601f840112610f8b57600080fd5b50813567ffffffffffffffff811115610fa357600080fd5b602083019150836020828501011115610fbb57600080fd5b9250929050565b60008060008060008060a08789031215610fdb57600080fd5b8635610fe681610f3b565b95506020870135610ff681610f3b565b94506040870135935061100b60608801610f60565b9250608087013567ffffffffffffffff81111561102757600080fd5b61103389828a01610f79565b979a9699509497509295939492505050565b6000815180845260005b8181101561106b5760208185018101518683018201520161104f565b8181111561107d576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e006020830184611045565b600080600080600080600060c0888a0312156110de57600080fd5b87356110e981610f3b565b965060208801356110f981610f3b565b9550604088013561110981610f3b565b9450606088013561111981610f3b565b93506080880135925060a088013567ffffffffffffffff81111561113c57600080fd5b6111488a828b01610f79565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561117657600080fd5b873561118181610f3b565b9650602088013561119181610f3b565b955060408801356111a181610f3b565b9450606088013593506111b660808901610f60565b925060a088013567ffffffffffffffff81111561113c57600080fd5b6000602082840312156111e457600080fd5b8151610e0081610f3b565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061126e6060830184866111ef565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526112c860c0830184866111ef565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006113046060830185611045565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
	immutableReferences["L2ERC721Bridge"] = true
}
