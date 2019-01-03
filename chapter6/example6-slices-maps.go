// Карта (также известна как ассоциативный массив или словарь) 

package main

import "fmt"

func main() {

    // вместо имени элемента сохраним какую-нибудь дополнительную информацию о нем
    elements := map[string]map[string]string{
        "H": map[string]string{
        "name":"Hydrogen",
        "state":"gas",
    },
    "He": map[string]string{
        "name":"Helium",
        "state":"gas",
    },
    "Li": map[string]string{
        "name":"Lithium",
        "state":"solid",
    },
    "Be": map[string]string{
        "name":"Beryllium",
        "state":"solid",
    },
    "B":  map[string]string{
        "name":"Boron",
        "state":"solid",
    },
    "C":  map[string]string{
        "name":"Carbon",
        "state":"solid",
    },
    "N":  map[string]string{
        "name":"Nitrogen",
        "state":"gas",
    },
    "O":  map[string]string{
        "name":"Oxygen",
        "state":"gas",
    },
    "F":  map[string]string{
        "name":"Fluorine",
        "state":"gas",
    },
    "Ne":  map[string]string{
        "name":"Neon",
        "state":"gas",
    },
}

//вывод, если существует  
if el, ok := elements["Li"]; ok {    
    fmt.Println(el["name"], el["state"])
    }

}

/* карта с int похожа на массив, но есть различия:
- длина карты (которую мы можем найти так: len(x)) может измениться, когда мы добавим в нее новый элемент
В самом начале при создании длина 0, после x[1] = 10 она станет равна 1
- карта не является последовательностью. В нашем примере у нас есть элемент x[1], в случае массива должен быть и первый элемент x[0], но в картах это не так
*/
