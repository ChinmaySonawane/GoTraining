package player

import ("fmt"
		"time"
)

type Player struct{
	name string
	lastPlayed time.Time
}

func NewPlayer(name string) *Player {

	player := new(Player)
	player.name=name
	player.lastPlayed=time.Now()
	return player

}
 
func CheckPlayer(players []Player, playerName string) []Player {
	//var c int
	for _, v := range players {
		//c = k
		if v.name == playerName {
			fmt.Println("player exist")
			return players
		}
		
	}

	p := NewPlayer(playerName)
	players = append(players, *p)
	/*if len(players) != 0{
		fmt.Println("player exist")
		for _,v := range players{

			fmt.Println(v.name, v.lastPlayed)
		}
	}
//	return c+1
*/
	return players
}