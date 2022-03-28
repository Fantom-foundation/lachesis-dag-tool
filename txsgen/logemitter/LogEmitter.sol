pragma solidity >=0.4.22 <0.7.0;

/// @title Generating a lot of logs.
contract LogEmitter {

    /// Create a new log emitter.
    constructor() public {
    }

    event Logging(address indexed sender, bytes32 indexed str1, bytes32 indexed str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8);

    function emitLogs(bytes32 str1, bytes32 str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8) public payable {
	emit Logging(msg.sender, str1, str2, str3, str4, str5, str6, str7, str8);
    }

}

