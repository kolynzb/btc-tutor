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
  **NOTE** - when you see the term hash function, it means secure cryptographic or ideal hash function.
- The output values of hash functions are often referred to as message digests, hash values, checksums, or hashes.

- You can think of a message digest as analogous to a fingerprint for digital information. In the physical world, a fingerprint is an identifier unique to each individual person such that no two people ever share the same fingerprint. A hash function takes some input message (which can be almost anything as all digital information is essentially a binary string of ones and zeros), and produces a unique, compressed, fixed-length, digital fingerprint for that input message.

- There applied in

  - Hash tables
  - Digital Signatures
  - Keyed Hashes
  - One way transformations
  - Key derivation

- These operations and techniques form the foundation of most digital storage and communication technologies. However, the two fundamental uses of hash functions are: maintaining _data integrity_, and _data storage_ & _access efficiency_. Since message digests are digital fingerprints, they are very useful for making sure data hasn’t been tampered with.

  - [Click Me for an introduction to hashing and checksums in Linux and How to make use of checksums](https://www.redhat.com/sysadmin/hashing-checksums).
  - **Effiecient Data Storage and Rerieval**

    - Using hash functions as a tagging or identification mechanism in data storage or in a data structure can provide efficiency gains. Message digests of hash functions are a fixed length but still unique, and hash functions can take almost anything as an input value. Many data structures found in information systems utilize hash functions to identify or index data. Fundamentally, all computer systems use [hash tables](https://youtu.be/KyUTuwz_b7Q) (also called associative arrays, maps, or dictionaries) within their memory storage architectures to make memory access as efficient as possible. Hash tables are such an important data structure that all other data structures can be constructed using them; there are even some programming languages like AWK that use them exclusively. Hash tables work a lot like arrays; the difference being hash tables allow for anything that can be hashed to be used as an index value creating a key-value mapping.
    - For example, given an array of size 10, the contents of the array are stored by accessing the array at indexes from 0..9
    - [Click me 😭 to see array example](/0x00-hash_func/arrayex.go)
    - Whereas hash tables use a hash of data modulus the length of the array (i.e. the remainder of the message digest divided by the length of the array) to derive the index the data will be stored in the array.
    - [Click me 😭 to see hash table example](/0x00-hash_func/hashtableex.go)
    - Notice the output of the hash table is not in order like the array was, and that's because the index of the value is determined by the hash of its key. This is what makes hash tables so efficient. Because the index of the value is provided when fetching the data, retrieval is fast, yet almost any kind of input can be used to derive the index.

## Merkle Trees

- Bitcoin makes extensive use of Merkle Trees. A Merkle Tree is a binary tree structure (meaning its branches terminate in two nodes) that's made entirely of hashes. For example, the leaf nodes at the bottom layer of the Merkle trees used in Bitcoin are the hash values of the raw data from bitcoin transactions. And the successive layers of the tree leading back to the root node are constructed by concatenating the two hashes from the previous layer, and hashing them together to form a new hash. This has the effect of connecting the tree together to be represented securely by one single root node value. This structure makes it very easy to find data at the bottom of the tree, as well as to prove specific data exists in the tree which is particularly important for scaling the Bitcoin system efficiently.
  <img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less1_VA3%20updated.jpg" alt="Merkle Tree"/>

## Digital Signatures

- Although hash tables and data structures like Merkle Trees are where hash functions find their greatest use, the most widely recognized use of hash functions is in cryptography. In addition to encryption and key-exchanges, digital signatures make extensive use of hash functions. Like a physical signature, a digital signature associates an identity to a document or message. To ensure the security of digital signatures, well tested algorithms are used to both construct and verify them.
- In Bitcoin, digital signatures are most commonly used for Pay-to-Public-Key-Hash (P2PKH) transactions – the most common type of Bitcoin transaction template -- where the ownership of funds are transferred from the sending party to the receiving party. The two other common uses of digital signatures in Bitcoin are for signing arbitrary messages and using a Bitcoin address as part of the verification process, and MinerID which is a message included in the first transaction of a block that acts as an authentication mechanism in case another miner (Bitcoin node) tries to act dishonestly while using the ID of a competing miner.
- With that said, it's helpful to recognize that even though hash functions and cryptographic digital signatures are used in Bitcoin, it is not a cryptographic system, and it's actually incorrect to use the term “cryptocurrency” when referring to Bitcoin or a system that follows the Bitcoin architectural design like BSV

## Differences between Hashing and Encryption

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2F908ZCORCWaRrXH9UHMDR%2FBSVA-HashFunctions_Ch1L2_DA1.gif?alt=media&token=6cfb2d3a-06d2-47cc-bc32-7f43c70e385d" alt="Decryption"/>

1. **Hash functions don't care about data loss**: A hash function produces a fixed-length output in clear text because it’s nothing more than a unique digital fingerprint and data loss is not a concern. In contrast, encryption encodes and preserves input messages.
2. **Hash functions are not reversible**: A hash function is a one-way function that does not need to be reversed. In fact, it's important it's not reversible. In contrast, it's crucial the intended recipient of an encrypted message is able to reverse the encryption process in order to decrypt the message so they can read it.
3. **Hash functions have a collision risk**: A hash function has a collision risk. Due to the fixed length of message digests, it's possible for two different input values to produce the same digest. Bitcoin uses hash functions like SHA-256 and RIPEMD-160 because they greatly reduce the risk of collisions, yet, although they're probabilistically insignificant, they can still theoretically happen. In contrast, since encryption preserves 100% of the input message, each encrypted message produces a unique output so there's no collision risk.

## Three Important Properties of Hash Functions.

1. **Preimage Resistance: One-way**
   <img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FB0wsRpq9bChNJr1i6IGj%2FBSVA-HashFunctions_Ch1L3_DA2.gif?alt=media&token=75773213-8eb6-4951-9820-d65734caebee" alt="Preimage resistance"/>
   - Hash functions are computationally [infeasible](https://www.merriam-webster.com/dictionary/infeasible) to reverse -- in the same sense as trying to unscramble an egg or put toothpaste back into its tube. Said in another way, hash functions are one-way 'trap-door' functions; they make it easy to find a message digest from an input, but so difficult to find an input from a message digest, that it would either take longer than human beings are capable of waiting (sometimes longer than the age of the universe) or the cost of marshalling the necessary computing power to reverse them is so high, it's not worth doing it.

- To be considered preimage resistant, it should take
  2^m
  hash attempts (where m is the length of the output in bits) to find the corresponding input message for a particular message digest. If a preimage can be found using fewer than
  2^m
  hash attempts, the hash function is no longer considered preimage resistant and is no longer secure.

2. **Second Preimage Resistance: Deterministic**

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FTt1pWx5W4GZI4ACdRCK8%2FBSVA-HashFunctions_Ch1L3_DA3.gif?alt=media&token=c6552c18-4dc5-45b3-8914-09f281f3b78a" alt="2nd preimage resistance"/> 
- The message digest of a hash function is consistent given the same input value. In other words, hash functions are deterministic -- the same input produces the same output hash every time. Furthermore, if even a single bit of the input message is changed, the resulting output hash is significantly different, unique, and preimage resistant.

3.  **Collision Resistance**

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2Fp67oxC5KPI94htiIDYqo%2FBSVA-HashFunctions_Ch1L3_DA4.gif?alt=media&token=002874b4-766f-4433-a634-6bbcd818d950" alt="collition resistance"/>

- It’s computationally infeasible to find two different input values that produce the same output value for a hash function. Although the fixed length of their output means a hash function has a collision risk, an ideal hash function produces message digests with a length large enough that the risk of finding a collision is so small, it’s safe to ignore in most situations.
- Ideally, a hash function with an output value of bitlength `m`
  should require at least
  2^m/2
  operations to find a collision. In a _birthday attack_, given a set of `m `equally probable values, roughly
  2^m/2
  random samples are needed to find a single collision. This means if a collision is found in fewer than
  2^m/2
  random samples, the collision resistance property of the hash function is “broken”, and it’s no longer considered secure.
- [Collition Resistance](https://youtu.be/amEX2EdLIcs)
- [Merkle damgard Parardigm](https://youtu.be/VCOinPlsThw?list=PLPnhHDyS1o44pBF_HyGzWOzxQ9tEk_XIn)
- This is particularly important with respect to Bitcoin's _proof-of-work process_.

## The Hash Functions Found In Bitcoin

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FyASwEq5gwo3LOEm8gcXh%2FBSVA-HashFunctions_Ch1L4_DA1.gif?alt=media&token=10626f39-60d1-478b-bafb-60a5924d46fb" alt="hash and bitcoin"/>

## Merkle-Damgard Hash Functions

- The fundamental issue with constructing a hash function is the fact that it must be able to take an input of arbitrary length and compress it into an output of a fixed length. If input is an arbitrary length, it's difficult to achieve second preimage resistance. Also, for a compression function to be collision resistant, it must take fixed length inputs that are longer than its output. One way to remedy this issue is to use a preprocessing step to pad the input to a fixed length rather than compress it directly. Once the input has been preprocessed, it can be compressed consistently.
- In 1979, both [Ralph Merkle](https://en.wikipedia.org/wiki/Ralph_Merkle) and [Ivan Damgård](https://en.wikipedia.org/wiki/Ivan_Damg%C3%A5rd) independently proved that so long as an appropriate padding scheme is used in the preprocessing step, and the compression function used is collision resistant (implying preiamge and second preimage resistance), it follows that the hash function itself is also collision resistant. The most popular hash functions in use today are based on this Merkle-Damgård method of construction.
- Merkle-Damgård hash functions have three parts: input and preprocessing, compression, and final value construction and output. The preprocessing step pads the input message so its [congruent](https://dictionary.cambridge.org/dictionary/english/congruent) to a fixed bitlength, e.g. 512 or 1024, and then appends the length of the input itself (this is called Merkle-Damgård strengthening). The compression step uses bitwise logical functions to compute and mutate the processed input in rounds using chaining variables to link each round. Finally, the chaining variables are concatenated and outputted most commonly in hexadecimal format.

# MD4 and MD5

- The first two Merkle-Damgård hash functions to gain widespread adoption were Ron Rivest's MD4 and MD5. First specified in 1992 in the Internet Engineering Task Force (IETF) Request For Comments (RFCs) 1320 and 1321, respectively, MD4 and MD5 were the first hash functions to use a 448 congruent to 512 message block padding scheme with the remaining 64 bits left for Merkle-Damgård strengthening.
- MD4 enjoyed short-lived success, but Rivest realized shortly after its release that it was likely to be insecure, so he designed MD5 shortly thereafter. Despite Hans Dobbertin's announcement of a collision found for MD5's compression function in 1996, MD5 enjoyed long-lived success. While cryptographers were recommending SHA-1 as a replacement to MD5 as early as 1996, MD5 continued to be used well into the early 2000s.
- However, in 2004, Xiaoyun Wang, Dengguo Feng, Xuejia Lai, and Hongbo Yu developed an analytical collision attack that could reportedly be performed within an hour on an IBM p690 cluster. This marked the end for MD5 for use in key derviation and digital signatures. In 2008, CMU Software Engineering Institute announced MD5 was "cryptographically broken and unsuitable for further use".
- Even so, the groundwork laid by MD4 and MD5 provided a practical framework for newer generations of hash functions to follow; including Bitcoin's SHA-256 and RIPEMD-160 which both use 512 bit message blocks and 32 bit words.

## Bitcoin's Hash Functions

- As already mentioned, the two hash functions found in the Bitcoin system are SHA-256 and RIPEMD-160. For reasons explored further in Chapter 7, Bitcoin uses SHA-256 and RIPEMD-160 in double hash forms: SHA-256(SHA-256) and RIPEMD-160(SHA-256) – often abstracted and referred to as HASH-256 and HASH-160, respectively, following the convention set out in the [original bitcoin code](https://github.com/trottier/original-bitcoin).

- However, the following table displays the full range of hash functions commonly found in the greater Bitcoin ecosystem; including SHA-512 and the HMACs of SHA-256 and SHA-512 which are often utilized by Bitcoin wallet implementations:

![hash functions](/images/hashfunctable.jpg)

## What is Base58 and Why Does Bitcoin Use It?

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FmskOpCwSISbsbvGXIjYs%2FBSVA-HashFunctions_Ch2L1_DA1.gif?alt=media&token=bdbe954b-e938-4ae3-88ac-02a050ac7297" alt=""/>

- While Base58 & Base58Check are not hash functions, they are important topics to cover.
- Base58 is a binary-to-text encoding of data to plain text; i.e., printable characters.
- Like Base64, Base58 is used to represent 8-bit bytes of binary data in an **ASCII** string format. Bitcoin uses a binary-to-text encoding to enable data stored in binary formats to be transmitted across communication channels that only reliably support text values like email and the internet.

```go

// Why base58 instead of standard base-64 encoding?
// - Don't want 0OIl characters that look the same in some fonts and
// could be used to create visually identical looking account numbers.
// - A string with non-alphanumeric characters is not as easily accepted as an account number.
// - E-mail usually won't line-break if there's no punctuation to break at.
// - Doubleclicking selects the whole number as one word if it's all alphanumeric.-
```

- In other words, since Base58 avoids using non-alphanumeric characters (+ and /) and ambiguous letters (0 - zero, I - capital i, O - capital o, and l - lower case L), it helps ensure Bitcoin addresses are readable, transmittable, and more secure.

<!-- - TODO: add table  https://5thwork.com/courses/course-v1:BIP_Uganda+BIPU0001+2022Q4/courseware/c75b47f9f18b462494b9d67ecc7edd92/00bba1592dcc4cd7b11eda620a8ef7a2/?child=first -->

- Using [hash calculator](https://bitcoinsv.academy/hash-calculator) encode the following hash (displayed in HEX) in Base58:
  `094d77441cfead50daa8e9ce893698962dbcbbce`

<img src="https://1089794075-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FHeYD5HNM81vcf8bAxUyk%2Fuploads%2FlPCIiAzU11plvspnOyED%2FBSVA-HashFunctions_Ch2L2_DA1%20(1).gif?alt=media&token=66b99488-c37f-48a7-ba59-6353cf0c7cff" alt="BASE58"/>
