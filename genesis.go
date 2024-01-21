package main

func Genesis() (*Block, error) {
	prevHash := []byte{}
	return NewBlock(prevHash, "Genesis Block!")
}
