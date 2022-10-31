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

## Chapter 3: Basic data types and property

### Basic data types

- sCrypt is a strongly typed language. The basic data types include:

1. `bool`: boolean value, true or false

2. `int`: signed integer

3. `int[3]`: Integer array of size 3

4. `PubKey`: public key

5. `Sig`: signature

Among them, PubKey and Sig are subtypes of bytes. If you want to know more about basic types, you can [check language reference documentation](https://scryptdoc.readthedocs.io/en/latest/)

## Properties

- Each contract can have several properties, which can be accessed through the `this` keyword in the function of the contract, as in Object-Oriented languages:

```srypt
contract TicTakToe {
    int x;
    bool y;
    bytes z;

	public function unlock(int y) {
		require(this.x == y);
	}

}
```

- The Tic-Tac-Toe game contract supports two players. The contract needs to save the addresses of the two players. After the contract runs, the contract automatically sends the locked bitcoin to the winner unless the game ends in draw.
- Define two attributes alice and bob, the data types are both PubKey.

```scrpt
contract TicTakToe {
   PubKey alice ;
   PubKey bob ;
}
```

## Chapter 4: Static Properties and Const Variables

### Static Properties

- A property decorated with the keyword static is a static property, which must be initialized when the property is declared. In the function of the contract, it can be accessed through the contract name and the property name (with a dot in the middle). The contract prefix can be omitted if used in the same contract where it is defined.

```scrypt
contract Test {
    static int x = 12;

    public function unlock(int y) {
        Test.x = y;
        // using prefix
        int z = Test.x + y;
        // without prefix
        z = x;
    }
}
```

### Const Variables

- A variable declared using the const keyword cannot be changed once it is initialized, as shown below:

```scypto
contract Test {
    const int x;

    constructor(int x) {
        // initialize a const property
        this.x = x; // good
    }

    public function equal(const int y) {
        y = 1; // <-- error

        const int a = 36;
        a = 11; // <-- error

    }
}
```

- A property can be declared using both. This is typically used to define constants.

```scrypt
contract Test {
    static const int PERCENT = 100;
}
```


Chapter 5: Function and Require Statement
Function
Functions are declared using the function keyword. The main purpose of a function is to encapsulate the internal logic of the contract and code reuse. When defining a function, you need to use a colon : to indicate the type of return value.

The following Test contract contains a clone function:
```scrypt
contract Test {

    function clone(int x) : int {
        return x;
    }
}
```
