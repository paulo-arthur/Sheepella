import os, json, random

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

datas = readJSON()

score = 0

print()
print("Welcome to your quiz")
print("Type the meaning of the words I will show")
print()

for data in datas:
    print()
    print("=-" * 10)
    print(data[0])
    anwser = input("Meaning: ").lower()

    if anwser in data[1]:
        print("Correct")
        score += 1
    else:
        print("No, you were wrong")
        print("Correct meaning(s):")

        for meaning in data[1]:
            print(meaning)
    print("=-" * 10)

print("Total Score: ", score*10/len(datas))
os.system("pause")