// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

//bootstrap nodes Ellaism formerly called bootstap in mainnet.json renaqmed to MainnetBootnodes
var MainnetBootnodes = []string{
	"enode://98b48cc7149326d00a57994ad55014a095a3e5cd4f0144cd7b034fb667d8e8017082bd90047d72c4403798b8ece11a33bd2e344d6500faba30889ebcfa5316fa@172.104.163.204:30303",
        "enode://834246cc2a7584df29ccdcf3b5366f118a0e291264980376769e809665a02c4caf0d68c43eecf8390dbeaf861823b05583807af0a62542a1f3f717046b958a76@45.77.106.33:30303",
        "enode://d1373d7187a20d695220fafb5746a289786bad92469d6bbe800f07572f8f444479f507cfcb8fcd8f5bee5dd44efb6185df4a1e4f1a7170408ada9876ef5b3fe4@178.79.189.58:30303",
        "enode://d8059dcb137cb52b8960ca82613eeba1d121105572decd8f1d3ea22b09070645eeab548d2a3cd2914f206e1331c7870bd2bd5a231ebac6b3d4886ec3b8e627e5@173.212.216.105:30303",
        "enode://5a89c8664d29a321fd4cb1b55b0f0be832ce376b5e7feb14a2073fdbd9bd7b7394169ed289dd991112b42ecfb74ea36e436bc72a1c99dcdb50d96eaf3b0ed254@213.136.91.42:30303",
        "enode://9215ad77bd081e35013cb42a8ceadff9d8e94a78fcc680dff1752a54e7484badff0904e331c4b40a68be593782e55acfd800f076d22f9d2832e8483733ade149@213.14.82.125:30303",
        "enode://07913818dafbadf44d4fc796fa414ec1d720ecfb087eff37efbe7134556658e92351559de788fa319c291e40b915cc26d902069d03bd935553d4efa688bdbbf8@45.32.19.37:30303",
        "enode://645a59b6e6e20ed8864767a1d0c731f89ae276ed4e04c4f10becce655532d95cbe1bc9e2da3f13a6564f9ca8fe46fab2781a380b3a89148bccac883d6068f684@45.77.159.123:30303",
        "enode://7c2f43b2e7fded9469917311574d267427e62cd12e1864effd15f31c1246e4e955463d843acaa37309693a515df7986cb6d160b7e85a4ca2779a798a72d90850@108.61.194.191:30303",
        "enode://5dd35866da95aea15211fb1f98684f6e8c4e355e6aa3cc17585680ed53fa164477b8c52cb6ca4b24ec4d80f3d48ff9212b53feb131d825c7945a3abaaf02d24d@178.79.189.58:60606",
        "enode://6c585c18024eb902ca093278af73b04863ac904caabc39ac2920c23532307c572ad92afd828a990c980d272b1f26307f2409cc97aec3ff9fe866732cae49a8c2@144.217.163.224:31337",
        "enode://edd90c4cc64528802ad52fd127d80b641ff80fd43fa5292fb111c8bd2914482dffee288fd1b0d26440c6b2c669b10a53cbcd37c895ba0d6194110e100a965b2d@188.166.179.159:30303",
        "enode://9d960373335c1cc38ca696dea8f2893e2a071c8f21524f21e8aae22be032acc3b67797b1d21e866f9d832943ae7d9555b8466c6ab34f473d21e547114952df37@213.32.53.183:30303"
}

/****commentiing out all other nodes until test of Ellaism*****
// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303",  // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303",  // IE
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
********end of comment out all other nodes test Ellaism******///
