- How to get WIF from Private Key
- first: get private key
  `5ee98fbc968fb9abd3bb98e0d7792080d62014aac3ea7245676716dcfca9e160`

- 2: add prefix to show if you are using the main net or the testnet. add 80 if mainnet

`805ee98fbc968fb9abd3bb98e0d7792080d62014aac3ea7245676716dcfca9e160`

- Append the type of compression 01
  `805ee98fbc968fb9abd3bb98e0d7792080d62014aac3ea7245676716dcfca9e16001`

- get hash-256 of what we have
  `a030a0b3c78c828fc16a610bc0ce9761561759edad91ef7fb739a0d17b7db07c`

- get First 4 bytes of our result and append it to result in line 10
  `805ee98fbc968fb9abd3bb98e0d7792080d62014aac3ea7245676716dcfca9e16001a030a0b3`

- get base58 of previous value, and that is the answer:
  `KzQD2FkVBJXE3m69eeWodG3oFQTuUA1kE2LNmAdCq77ksG1sCd5c`
