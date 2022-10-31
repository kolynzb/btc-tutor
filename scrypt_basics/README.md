sCrypt is a high-level programming language for writing Bitcoin smart contracts.

An sCrypt contract is the basic module of Bitcoin dApp. All variables and functions belong to a contract. It is the starting point for all your dApps.

An sCrypt contract is defined by the keyword contract, similar to class in object-oriented languages such as Java/C++. An contract called HelloWorld is as follows:

```srycpt
    contract HelloWorld {
	public function unlock(bytes message) {
		require(message == "hello world, sCrypt 😊");
	}
}
```
