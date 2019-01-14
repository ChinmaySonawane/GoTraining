package turn

import ("fmt"
		"strings"
)

//GameStat struct for game info

type GameStat struct
{		
		playerName string
		chance int
		word string
		playerGuess string
		win bool;
}

//contructor for GameStat
func NewGameStat(name string, word string)*GameStat  {	

	GameStatObj := new(GameStat)
	GameStatObj.playerName = name
	GameStatObj.word = word
	GameStatObj.win = false
	GameStatObj.chance = 6
	GameStatObj.playerGuess = word
	return GameStatObj
}


//word guess 
func (obj *GameStat)GuessWord()  {
	
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
func (obj *GameStat)GuessLetter()  {
	
	fmt.Println("\nENTER Guess Letter :")
	var guessLetter string
	fmt.Scan(&guessLetter)
	
	temp := strings.Replace(obj.playerGuess,strings.ToLower(guessLetter),"_",-1)
	
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
func (obj *GameStat)Chance(){

	for obj.chance != 0 && obj.win != true	{
		fmt.Println("player :",obj.playerName)
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
			
			case 1		:obj.GuessWord()
			
			case 2		:obj.GuessLetter()
			
			default		:fmt.Println("WRONG CHOICE")
			
		}
	}
}

//printing last result
func (obj *GameStat)Show()  {
	
	fmt.Println("Player Name:", obj.playerName)
	fmt.Println("Word:", obj.word)
	if obj.win{
		fmt.Println("YOU Win")
	}else{
		fmt.Println("YOU Lose")
	}
}