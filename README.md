# This is Golang quiz console app
### Input
You need a ```.csv``` file with data stored like ```question,answer``` <br>
for example:
>1+1,2 <br>
>here 1+1 is a question and 2 is an answer
### Flags
You should run this app with 2 flags ```-csv``` and ```-limit``` <br>
```-csv``` stands for location of quiz csv file (dont forgeet to use quotes) <br>
```-limit``` stands for time limit for all quiz <br>
### Run example
Clone repository and use ```go run .\quiz_game.go -csv='quiz.csv' -limit=5```
