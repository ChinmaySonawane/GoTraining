package game

import ("fmt"
		"time"
		"math/rand"
		"Training/hangman/game/player"
		"Training/hangman/game/turn"
)


func Play()  {
	words:=[]string{"hello","people","hangman","play","if"}
	for{
		players := make([]player.Player,0)
		fmt.Println("\nENTER YOUR CHIOCE\n\t1.PLAY \n\t2.EXIT \nCHOICE:")
		var choice int
		fmt.Scan(&choice)
		switch choice {

			case 1  :	fmt.Println("\nENTER YOUR NAME")
						var name string
						fmt.Scan(&name)
						players = player.CheckPlayer(players, name)
						if len(players) != 0{
							//fmt.Println("player exist")
							for _,v := range players{

								fmt.Println(v)
							}
						}
					
					p := time.Now().UnixNano()
					rand.Seed(p)
					random := rand.Intn(5)
						turn := turn.NewGameStat(name,words[random])
						turn.Chance()
						turn.Show()

			case 2  :	return

			default :	fmt.Println("WRONG CHOICE")
		}
	}
}