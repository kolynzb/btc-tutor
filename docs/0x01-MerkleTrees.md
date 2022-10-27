# Merkle Trees

## What are Merkle Trees

<img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less1_VA1%20updated.jpg" alt="" />
- A Merkle tree is a data structure that leverages the properties of hash functions to create an efficient process for verifying the inclusion of a single data element within a large set of data. Hash functions are used to convert any type of digital input data, from strings to files, into a unique, fixed length, digital fingerprint.

<img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less1_VA2%20updated.jpg" alt="" />

<img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less1_VA3%20updated.jpg" alt="" />

- Merkle trees are formed from multiple data elements (A-P) by first putting each element through a hash function to generate the leaf nodes (HA - HP) at the bottom of the Merkle tree. To generate the upper layers of the tree back to the root node, these leaf node values are then concatenated together as ordered pairs and hashed to form another fixed length string on the next layer up of the tree toward the single root value (top of the image above illustrated as H(HA-H || HI-P). Each leaf node consists of a hashed data entry, so n data entries would result in n leaf nodes. As we move up the tree layers, we combine ordered interior node values together by hashing the concatenation of the two individual values provided at the previous layer.

- By convention , bottom leaf is refered to as `layer 0` with layers about being numbered naturally.

- As (for any layer) Merkle trees must have an even number of values to pair, when a dataset has an uneven number of entries to begin with or where a layer reveals an odd number of values to be paired, the last value (widow) is repeated, concatenated with itself and hashed to become paired with another node value at a higher layer of the Merkle tree.

- In this respect, the 16-leaf tree above has 8 nodes on layer 1, whereas a 17-leaf tree will have 9 layer 1 nodes with the ninth being HQQ. We can say that in a Merkle tree with n leaf nodes, if n/2 does not equal an integer, then the number of layer 1 nodes will be the next integer after n/2. This is the case for any layer of the tree. Another layer of the Merkle tree is required after every 2x entries are included, where a tree for 4 values will have 3 layers, 8 having 4 layers, 16 having 5 layers, and so on and so forth.

## Why Use A merkle Tree ?

<img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less2_VA4%20updated.jpg" alt="" />

- Due to the deterministic output of hash functions, if a single bit changes anywhere throughout the values or data that constitute any layer of the Merkle tree, every subsequent value closer to the Merkle root on the same side’s branch will be entirely different. As such, Merkle trees are used as an extremely efficient data verification tool since it can be safely assumed that only the exact data entries that were hashed to create the leaf nodes of a Merkle tree in the same specific order, could formulate its specific root.

- Therefore, we can verify that a data entry exists in Merkle tree by following the sequence of hashes from that data entry's index, up the layers of the tree, until we arrive at the root. For example, to verify the existence of element K in a data set represented by a Merkle root R, we need to know the Merkle path for K, which is the minimum list of node values required to reconstruct R from K. The Merkle path for K includes node values H(L), H(IJ), H(MNOP), H(ABCDEFGH). We first calculate the hash for K and then using the node values in its Merkle path to obtain Merkle root R' which is compared to R. This process is known as a Merkle Proof. Merkle Proofs are efficient as they only require log2(n) hash values to perform the proof. Further, they can be performed without revealing sensitive data (e.g. only hashed values would be required from other data entries to prove our preimage data entry).

<img src="https://bitcoinsv.academy/storage/photos/4318/BSVA-MerkleTrees_Ch1Less2_VA5%20updated.jpg" alt="" />

- Merkle trees are used in various distributed and peer-to-peer systems where data verification is critical. These systems have multiple versions of network data existing in multiple different locations. If all the data objects within these networked systems had to be verified using the objects themselves, this would create an extremely burdensome network throughput. Instead of sending an entire object, the hashing of these objects allows them to be transmitted efficiently as a secure fixed byte length representation.
- If network participants wish to know they are operating upon the same data set, they only need to send the Merkle root to verify whether the sets are the same. If the Merkle root is different, they can work downwards from the root through the branch structures to determine where the inconsistency lies.
- For example, if they can determine that their values for HABCDEFGH are both the same while they have different values for HIJKLMNOP then the inconsistency obviously lies within the right-hand branch structure and all objects within the left-hand branch are identical and need not be transmitted for examination. The Merkle tree can be traversed downwards to ascertain which data element has been modified
- If, however, there are multiple inconsistencies within the Merkle tree, the process of verification requires more steps to isolate the problematic nodes. In the diagram above it can be seen how if a new element Q replaces L, the Merkle proof fails. Examples of these distributed or peer-to-peer systems that employ Merkle trees as a data verification process are Git, BitTorrent and Bitcoin.
