// Карта (также известна как ассоциативный массив или словарь) 
// это неупорядоченная коллекция пар вида ключ-значение
// var x map[string]int

package main

import "fmt"

func main() {
//    var x map[string]int
    x := make(map[string]int)   // необходима инициализация карты перед использованием
    x["key"] = 10
    fmt.Println(x)

    // карта с ключом типа int
    y := make(map[int]int)
    y[1] = 10
    y[3] = 10
    y[5] = 10
    //удаление элемента
    delete(y, 3)
    fmt.Println(y, y[3]) //y[3] вернет значение 0

    //карта с химическими элементами
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements["Li"])
    //обратимся к несуществующему элементу, вернется пустая строка 
    fmt.Println(elements["Un"])
    // результат существования элемента - в переменную ok
    name, ok := elements["Un"]
    fmt.Println(name, ok)

    //Сперва мы пробуем получить значение из карты, а затем, если это удалось, мы выполняем код внутри блока
    if name, ok := elements["Ne"]; ok {    
        fmt.Println(name)
    }

}

/* карта с int похожа на массив, но есть различия:
- длина карты (которую мы можем найти так: len(x)) может измениться, когда мы добавим в нее новый элемент
В самом начале при создании длина 0, после x[1] = 10 она станет равна 1
- карта не является последовательностью. В нашем примере у нас есть элемент x[1], в случае массива должен быть и первый элемент x[0], но в картах это не так
*/
