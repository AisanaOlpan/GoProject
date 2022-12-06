package internal_test

import (
	"flag"
	"fmt"
)

func flagexample() {
	wordPtr := "рпа" //flag.String("word", "foo", "a string")

	//Тут мы декларируем флаги numb и fork, используя аналогичный подход, как и с флагом word.

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	//Также флаги можно определять, используя существующую переменную, объявленную ранее в программе. Обратите внимание, что мы передаем в функцию указатель.

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//Как только все флаги определены, вызываем flag.Parse(), чтобы выполнить непосредственно разбор командной строки.

	flag.Parse()

	//Тут мы выводим на экран считанные значения флагов, и оставшиеся аргументы командной строки. Обратите внимание. что мы должны разыменовывать указатель, *wordPtr, например, чтобы получить реальные значения.

	fmt.Println("word:", wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
