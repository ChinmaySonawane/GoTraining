package player
/*
Player struct
*/
import ("fmt"
		"time"
)

type Player struct{
	name string
	lastPlayed time.Time
}

//constuctor for Player
func NewPlayer(name string) *Player {

	player := new(Player)
	player.name=name
	player.lastPlayed=time.Now()
	return player

}
 
//tocheck player
func CheckPlayer(players []Player, playerName string) []Player {
	//var c int
	for _, v := range players {
		//c = k
		fmt.Println(v.name)
		if v.name == playerName {
			fmt.Println("player exist")
			return players
		}
		
	}

	p := NewPlayer(playerName)
	players = append(players, *p)
	if len(players) != 0{
		for _,v := range players{
			
			fmt.Println(v.name, v.lastPlayed)
		}
	}
	return players
}