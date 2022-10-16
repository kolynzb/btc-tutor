# Hash Functions

## Introduction

- To understand the intricacies of bitcoin we need to understand these primitive concepts _Hash functions_ , _merkle trees_ and _Digital signitures_.

# Project init

- Create a go mod
  `go init mod github.com/username/bsv_examples`
- Get bsv packages

```bash
go get github.com/libsv/go-bt
go get github.com/libsv/go-bk
go get github.com/libsv/go-bc
```

# Hashing Algorithm

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FByw2otGdq4oQHAHLYG9s%2FBSVA-HashFunctions_Ch1L1_DA1.gif?alt=media&token=ca9d88bc-3315-464d-9f63-49fb115a31dd"/>

- A hash function is a mathematical process that maps a bitstring of any length to a bitstring of a defined length. A one-way hash function must be _preimage resistant_ and _second preimage resistant_ while a secure cryptographic or ideal hash function must be preimage resistant, second preimage resistant, and _collision resistant_
__NOTE__ - when you see the term hash function, it means secure cryptographic or ideal hash function.
- The output values of hash functions are often referred to as message digests, hash values, checksums, or hashes.