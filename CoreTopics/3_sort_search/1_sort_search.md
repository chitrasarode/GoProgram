Let's implement following sorting searching algo
1. Linear search
2. Binary search
3. Bubble sort
4. Merge sort
5. Quick sort

Let's implement these algo where we can sort different types of data:
1. Int
2. Float
3. String

It should work for both ascending and descending. 


myBubbleSortIntAscending(arr []int){

//two nested for loop
for {
    for {
        //ascending order
        if num1 < num2 {
            //swap the ele
        }
    }


}

myBubbleSortIntDescending(arr []int) {

//two nested for loop
for {
    for {
        //descending order
        if num1 > num2 {
            //swap the ele
        }
    }
}

myBubbleSortInt(arr []int, campInt func(int, int)bool ) {

//two nested for loop
for {
    for {
        //user defined comparison order
        if campInt(num1, num2) {
            //swap the ele
        }
    }
}

}

func AscendInt(num1, num2 int) bool {
    if num1 < num2 {
        return true
    } else {
        return false
    }
}

func DescendInt(num1, num2 int) bool {
    if num1 > num2 {
        return true
    } else {
        return false
    }
}


main() {

    arr := []int { 10, 12, 101, 0, 4}

    myBubbleSortInt(arr, AscendInt)

    myBubbleSortInt(arr, DescendInt)
    
}

Test cases and also use time to check the speed of your algo.
Also use the rand function to generate large set of data, 100, 1000, 100000, 1000000. 
