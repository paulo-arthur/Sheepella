import os, json

#json.dumps, json.loads

def pause():
    os.system("pause")

os.system("go run main.go scrap.go")
pause()
os.system("cls")

print()
print("Now read the news and add the words and meanings that you don't understand")
print()

def readJSON():
    try:
        with open("./DataBases/data.json", "r") as data_file:
            data = json.load(data_file)
        return data
    except FileNotFoundError:
        print("Nao localizei o arquivo data.json")

def writeJSON(data):
    try:
        with open("./DataBases/data.json", "w") as data_file:
            json.dump(data, data_file, indent=4)
    except:
        print("Nao consegui escrever em data.json")

def addNewWord():
    print()
    word = input("Type the word in English: ").lower()
    meanings = []

    print()
    print("Enter the meanings in portuguese, to finish press Enter")
    while True:
        meaning = input("Enter here: ").lower()
        meaning = meaning if meaning != "" else False

        if (not meaning): break

        meanings.append(meaning)
        print()
    
    print()
    example = input("Enter an example sentence where this word is present:\n")

    data = readJSON()
    data.append([word, meanings, example])
    writeJSON(data)

    #word, meanings, example

def showKnownWords():
    datas = readJSON()
    for data in datas:
        print("=-"*10)
        print(f"Word: {data[0]}")

        print("Meanings: ")
        for meaning in data[1]:
            print(meaning)

        print(f"Example: {data[2]}")
        print("=-"*10)
        print()
    print()
    pause()

while True:
    print("""Press N to add a new word\nPress E to open the database\nPress S to show words\nPress T to start the test\nPress Q to quit """)
    print()

    op = input(">").lower()

    if op == "n":
        addNewWord()
    elif op == "e":
        os.system("notepad Databases/data.json")
    elif op == "s":
        showKnownWords()
    elif op == "t":
        os.system("python quiz.py")
    elif op == "q":
        break
    else:
        print("Try again.")
    os.system("cls")