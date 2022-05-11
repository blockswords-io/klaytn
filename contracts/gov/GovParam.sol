//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "./Ownable.sol";

contract GovParam is Ownable {
    mapping(uint => Param) private params;
    uint[] private paramIds;

    struct Param {
        string name;      // ex) "istanbul.epoch"
        uint64 activationBlock; // block number for last change
        bytes  prevValue; // RLP encoded value before nextBlock
        bytes  nextValue; // RLP encoded value after nextBlock
    }

    struct ParamView {
        string name;
        bytes  value;
    }

    event VoteContractUpdated(address);
    event VotableUpdated(bool);
    event ParamVotableUpdated(uint, bool);
    event SetParam(uint, string, bytes, uint64);
    event ValidatorAdded(address);
    event ValidatorRemoved(address);

    constructor(address _owner) {
        if (_owner != address(0)) {
            _transferOwnership(_owner);
        }
    }

    modifier onlyVotable(uint idx) {
        require(msg.sender == owner(), "permission denied");
        _;
    }

    function addParam(uint id, string memory _name, bytes memory value) public
    onlyOwner {
        require(bytes(_name).length > 0, "name cannot be empty");
        require(bytes(params[id].name).length == 0, "already existing id");
        params[id] = Param({
            name: _name,
            activationBlock: 0, // or block.number
            prevValue: "",  // unused
            nextValue: value
        });

        paramIds.push(id);
    }

    function setParam(uint id, bytes calldata value, uint64 _activationBlock) public
    onlyVotable(id) {
        require(params[id].activationBlock < block.number, "already have a pending change");
        require(block.number < _activationBlock, "cannot set activationBlock to past");
        require(bytes(params[id].name).length > 0, "no such parameter");

        params[id].prevValue = params[id].nextValue;
        params[id].activationBlock = _activationBlock;
        params[id].nextValue = value;

        emit SetParam(id, params[id].name, value, _activationBlock);
    }

    function getParam(uint n, uint id) public view returns (bytes memory) {
        if (n >= params[id].activationBlock) {
            return params[id].nextValue;
        }
        else {
            return params[id].prevValue;
        }
    }

    function getAllParams(uint n) external view returns (ParamView[] memory) {
        ParamView[] memory ret = new ParamView[](paramIds.length);
        for (uint i = 0; i < paramIds.length; i++) {
            Param memory p = params[paramIds[i]];
            bytes memory value = getParam(n, paramIds[i]);
            ParamView memory pv = ParamView({
                name: p.name,
                value: value
            });
            ret[i] = pv;
        }
        return ret;
    }
}
