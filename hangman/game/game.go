package game

import ("fmt"
		"time"
		"math/rand"
		"Training/hangman/player"
		"strings"
)

//GameStat struct for game info
type GameInfo struct
{		
		playerName string
		chance int
		word string
		playerGuess string
		win bool
}

//contructor for GameStat
func newGameInfo(name string, word string) GameInfo  {	
	// TODO: change this syntax as per Go
	/*GameStatObj := new(GameStat)
	GameStatObj.playerName = name
	GameStatObj.word = word
	GameStatObj.win = false
	GameStatObj.chance = 6
	GameStatObj.playerGuess = word*/
	return GameInfo{playerName : name, word : word, win : false, chance : 6, playerGuess :word}
}

//word guess 
func (obj *GameInfo)guessWord()  {
	
	fmt.Println("\nENTER Guess Word :")
	var guessWord string
	fmt.Scan(&guessWord)
	 
	if strings.ToLower(guessWord) == obj.word {
		obj.win = true
		//return true
	} else {
		obj.chance = 0
	}
}

//letter guess
func (obj *GameInfo)guessLetter()  {
	
	fmt.Println("\nENTER Guess Letter :")
	var guessLetter string
	fmt.Scan(&guessLetter)
	
	temp := strings.Replace(obj.playerGuess, strings.ToLower(guessLetter), "_", -1)
	
	if temp == obj.playerGuess {
		obj.chance -= 1
	}else {
		obj.playerGuess=temp
		for	_,v := range obj.playerGuess{
			
			if(string(v) != "_"){
				return
			}
		}
		fmt.Println("hello")
		obj.win = true
	}

}

//chance to start game 
func (obj *GameInfo) startGame(){

	for obj.chance != 0 && obj.win != true	{
		fmt.Println("player :", obj.playerName,"chances remaining:", obj.chance)
		//fmt.Println(obj.word, "\t", obj.playerGuess, "\t", obj.win)
		orignal := []rune(obj.word)
		guess := []rune(obj.playerGuess)
		op := []rune("_")
		for i := 0; i < len(orignal); i++ {
			if(guess[i] == op[0]){
				fmt.Printf("\t %v",string(orignal[i]) )
			}else{
				fmt.Printf("\t _")
			}
		}	
	
		fmt.Println("\nENTER YOUR CHIOCE\n\t1.Guess Word \n\t2.Guess Letter \nCHOICE:")
		var choice int
		fmt.Scan(&choice)
		
		switch choice {
			
			case 1:
				obj.guessWord()
			
			case 2:
				obj.guessLetter()
			
			default	:
				fmt.Println("WRONG CHOICE")
			
		}
	}
}

//printing last result
func (obj GameInfo) show()  {
	
	fmt.Println("Player Name:", obj.playerName)
	fmt.Println("Word:", obj.word)
	if obj.win{
		fmt.Println("YOU Win")
	}else{
		fmt.Println("YOU Lose")
	}
}

func Play()  {
	words:=[]string{"hello", "people","hangman","play","if"}
	players := make([]player.Player,0)
	for{
		
		fmt.Println("\nENTER YOUR CHIOCE\n\t1.PLAY \n\t2.EXIT \nCHOICE:")
		var choice int
		fmt.Scan(&choice)
		switch choice {

			case 1  :	fmt.Println("\nENTER YOUR NAME")
						var name string
						fmt.Scan(&name)
						players = player.CheckIfExists(players, name)
						if len(players) != 0{
							//fmt.Println("player exist")
							for _,v := range players{

								fmt.Println(v)
							}
						}
					
					p := time.Now().UnixNano()
					rand.Seed(p)
					random := rand.Intn(5)
						obj := newGameInfo(name,words[random])
						obj.startGame()
						obj.show()

			case 2  :	return

			default :	fmt.Println("WRONG CHOICE")
		}
	}
}