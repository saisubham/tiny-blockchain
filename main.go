package main

func main() {
	bc := MakeBlockchain()
	defer bc.db.Close()

	cli := CLI{bc: bc}
	cli.Run()
}
