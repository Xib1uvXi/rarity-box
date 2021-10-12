#!/usr/bin/env bash

mkdir -p ./pkg/contractlib/rarity-contract/rarity
mkdir -p ./pkg/contractlib/rarity-contract/attributes
mkdir -p ./pkg/contractlib/rarity-contract/gold
mkdir -p ./pkg/contractlib/rarity-contract/skills
mkdir -p ./pkg/contractlib/rarity-contract/craft-i
mkdir -p ./pkg/contractlib/rarity-contract/crafting-i


abigen --abi=./pkg/contractlib/rarity-abi/rarity.abi.json --pkg=rarity --out=./pkg/contractlib/rarity-contract/rarity/contract.go
abigen --abi=./pkg/contractlib/rarity-abi/attributes.abi.json --pkg=attributes --out=./pkg/contractlib/rarity-contract/attributes/contract.go
abigen --abi=./pkg/contractlib/rarity-abi/gold.abi.json --pkg=gold --out=./pkg/contractlib/rarity-contract/gold/contract.go
abigen --abi=./pkg/contractlib/rarity-abi/skills.abi.json --pkg=skills --out=./pkg/contractlib/rarity-contract/skills/contract.go
abigen --abi=./pkg/contractlib/rarity-abi/Craft.I.abi.json --pkg=craft_i --out=./pkg/contractlib/rarity-contract/craft-i/contract.go
abigen --abi=./pkg/contractlib/rarity-abi/Crafting.I.abi.json --pkg=crafting_i --out=./pkg/contractlib/rarity-contract/crafting-i/contract.go

mkdir -p ./pkg/contractlib/caller-contract/caller
abigen --abi=./pkg/contractlib/caller-abi/caller.abi.json --pkg=caller --out=./pkg/contractlib/caller-contract/caller/contract.go