package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!")
}

/*
-javaagent:/Users/ryyt/dev/app/jetbrains/ja-netfilter.jar=jetbrains
-Dsun.java2d.metal.displaySync=false

--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED
--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED
*/
