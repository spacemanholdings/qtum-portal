package portal

type method struct {
	// This method does not require user authorization
	NoAuth bool
}

var rpcMethods = map[string]method{
	// == Blockchain ==
	"callcontract": {true},
	"getaccountinfo": {true},
	"getbestblockhash": {true},
	"getblock": {true},
	"getblockchaininfo": {true},
	"getblockcount": {true},
	"getblockhash": {true},
	"getblockheader": {true},
	"getchaintips": {true},
	"getdifficulty": {true},
	"getmempoolancestors": {true},
	"getmempooldescendants": {true},
	"getmempoolentry": {true},
	"getmempoolinfo": {true},
	"getrawmempool": {true},
	"gettransactionreceipt": {true},
	"gettxout": {true},
	"gettxoutproof": {true},
	"gettxoutsetinfo": {true},
	"listcontracts": {true},
	"preciousblock": {true},
	"pruneblockchain": {true},
	"verifychain": {true},
	"verifytxoutproof": {true},

	// == Control ==
	"getinfo": {true},
	"getinfo": {true},
	"getmemoryinfo": {true},
	"help": {true},
	"stop": {true},

	// == Generating ==
	"generate": {true},
	"generatetoaddress": {true},

	// == Mining ==
	"getblocktemplate": {true},
	"getmininginfo": {true},
	"getnetworkhashps": {true},
	"getstakinginfo": {true},
	"getsubsidy": {true},
	"prioritisetransaction": {true},
	"submitblock": {true},

	// == Network ==
	"addnode": {true},
	"clearbanned": {true},
	"disconnectnode": {true},
	"getaddednodeinfo": {true},
	"getconnectioncount": {true},
	"getnettotals": {true},
	"getnetworkinfo": {true},
	"getpeerinfo": {true},
	"listbanned": {true},
	"ping": {true},
	"setban": {true},
	"setnetworkactive": {true},

	// == Rawtransactions ==
	"createrawtransaction": {true},
	"decoderawtransaction": {true},
	"decodescript": {true},
	"fromhexaddress": {true},
	"fundrawtransaction": {true},
	"gethexaddress": {true},
	"getrawtransaction": {true},
	"sendrawtransaction": {true},
	"signrawtransaction": {true},

	// == Util ==
	"createmultisig": {true},
	"estimatefee": {true},
	"estimatepriority": {true},
	"estimatesmartfee": {true},
	"estimatesmartpriority": {true},
	"signmessagewithprivkey": {true},
	"validateaddress": {true},
	"verifymessage": {true},

	// == Wallet ==
	"abandontransaction": {true},
	"addmultisigaddress": {true},
	"addwitnessaddress": {true},
	"backupwallet": {true},
	"bumpfee": {true},
	"createcontract": {true},
	"dumpprivkey": {true},
	"dumpwallet": {true},
	"encryptwallet": {true},
	"getaccount": {true},
	"getbalance": {true},
	"getnewaddress": {true},
	"getrawchangeaddress": {true},
	"getreceivedbyaddress": {true},
	"gettransaction": {true},
	"getunconfirmedbalance": {true},
	"getwalletinfo": {true},
	"importaddress": {true},
	"importmulti": {true},
	"importprivkey": {true},
	"importprunedfunds": {true},
	"importpubkey": {true},
	"importwallet": {true},
	"keypoolrefill": {true},
	"listaddressgroupings": {true},
	"listlockunspent": {true},
	"listreceivedbyaccount": {true},
	"listreceivedbyaddress": {true},
	"listsinceblock": {true},
	"listtransactions": {true},
	"listunspent": {true},
	"lockunspent": {true},
	"removeprunedfunds": {true},
	"reservebalance": {true},
	"sendmany": {true},
	"sendmanywithdupes": {true},
	"sendtoaddress": {true},
	"sendtocontract": {true},
	"settxfee": {true},
	"signmessage": {true},
	"waitforlogs": {true},
}
